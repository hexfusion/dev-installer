package cluster

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"text/template"
	"time"

	jsonm "github.com/evanphx/json-patch"
	"github.com/ghodss/yaml"
	"github.com/hexfusion/dev-installer/pkg/cluster/config"
	"github.com/hexfusion/dev-installer/pkg/template_assets"
	"github.com/openshift/oc/pkg/cli/admin/release"
	imagemanifest "github.com/openshift/oc/pkg/cli/image/manifest"
	"github.com/openshift/oc/pkg/cli/registry/login"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/klog/v2"
	kcmdutil "k8s.io/kubectl/pkg/cmd/util"
)

// clusterOpts holds values to drive the cluster command.
type clusterOpts struct {
	errOut           io.Writer
	provider         string
	providerRegion   string
	pullSecret       string
	releaseImage     string
	releaseImageType string
	installerPath    string
	keepBootstrap    bool
	baseDir          string
	name             string
	pullSecretName   string
	sshKeyPath       string
	singleStackIpv6  bool
	replicasMaster   string
	replicasWorker   string
	libvirtURI       string
	ocpVersion       string
}

const (
	// refresh details https://cloud.redhat.com/beta/openshift/token
	cloudRedHatTokenUrl       = "https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token"
	cloudRedHatAccessTokenUrl = "https://api.openshift.com/api/accounts_mgmt/v1/access_token"
	openShiftInstallerUrl     = "https://github.com/openshift/installer.git"
)

// NewClusterCommand creates a new cluster
func NewClusterCommand(errOut io.Writer) *cobra.Command {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		klog.Fatal(err)
	}

	// defaults
	clusterOpts := clusterOpts{
		errOut:     errOut,
		baseDir:    fmt.Sprintf("%s/clusters", homeDir),
		libvirtURI: "qemu+tcp://192.168.122.1/system",
	}
	cmd := &cobra.Command{
		Use:   "cluster",
		Short: "create cluster",
		Run: func(cmd *cobra.Command, args []string) {
			must := func(fn func() error) {
				if err := fn(); err != nil {
					if cmd.HasParent() {
						klog.Fatal(err)
					}
					fmt.Fprint(clusterOpts.errOut, err.Error())
				}
			}
			must(clusterOpts.Validate)
			must(clusterOpts.Run)
		},
	}

	clusterOpts.AddFlags(cmd.Flags())

	return cmd
}

func (c *clusterOpts) AddFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&c.name, "name", "n", c.name, "cluster name")
	fs.StringVarP(&c.provider, "provider", "p", c.provider, "cluster provider")
	fs.StringVar(&c.providerRegion, "provider-region", c.providerRegion, "region to use for the given provider")
	fs.StringVarP(&c.pullSecret, "pull-secret", "a", c.pullSecret, "pull secret to use for cluster creation")
	fs.StringVarP(&c.releaseImage, "release", "r", c.releaseImage, "release image")
	fs.StringVarP(&c.releaseImageType, "release-type", "t", c.releaseImageType, "the type of release image used. Example CI, Nightly, Custom")
	fs.StringVar(&c.installerPath, "installer-path", c.installerPath, "path of the compiled installer to use")
	fs.StringVar(&c.baseDir, "base-dir", c.baseDir, "path of the base dir to store cluster data")
	fs.StringVarP(&c.sshKeyPath, "ssh-path", "s", c.sshKeyPath, "path to public ssh key for cluster")
	fs.BoolVarP(&c.keepBootstrap, "keep-bootstrap", "k", c.keepBootstrap, "keep boostrap node around for debug")
	fs.BoolVar(&c.singleStackIpv6, "single-stack-ipv6", c.singleStackIpv6, "single stack IPV6, default false (IPv4)")
	fs.StringVarP(&c.replicasMaster, "replicas-master", "m", c.replicasMaster, "number of master compute replicas")
	fs.StringVarP(&c.replicasWorker, "replicas-worker", "w", c.replicasWorker, "number of worker compute replicas")
	fs.StringVar(&c.ocpVersion, "version", c.ocpVersion, "ocp version used to generate releaseImage")
	fs.StringVar(&c.libvirtURI, "libvirt-uri", c.libvirtURI, "URI for libvirt instance")
}

// Validate verifies the inputs.
func (c *clusterOpts) Validate() error {
	if len(c.provider) == 0 {
		return errors.New("missing required flag: --provider -p")
	}
	if len(c.releaseImage) == 0 {
		return errors.New("missing required flag: --release -r")
	}
	if len(c.releaseImageType) == 0 {
		return errors.New("missing required flag: --release-type -t")
	}
	if c.releaseImageType == "release" && c.releaseImageType != "nightly" {
		return errors.New("invalid combination: --release flag with release value must use --release-type: nightly")
	}

	if len(c.sshKeyPath) == 0 {
		return errors.New("missing required flag: --ssh-path -s")
	}
	if c.releaseImage == "latest" && c.ocpVersion == "" {
		return errors.New("`latest` release requires: --version")
	}
	return nil
}

type Auth struct {
	Type     string
	FileName string
}

type Cluster struct {
	opts        *clusterOpts
	PullSecrets []string
	Dir         string
	TemplateData
	Config config.File
}

type RedHatCloud struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Created     string
}

type TemplateData struct {
	SSHKey                string
	PullSecret            string
	ClusterName           string
	WorkerReplicas        string
	MasterReplicas        string
	LogLevel              string
	ProviderRegion        string
	ClusterCidr           string
	ClusterCidrHostPrefix int
	MachineCidr           []string
	ServiceCidr           string
	NetworkType           string
	LibvirtURI            string
	ClusterDir            string
}

type releaseStream struct {
	Name        string `json:"name"`
	Phase       string `json:"phase"`
	PullSpec    string `json:"pullSpec"`
	DownloadUrl string `json:"downloadURL"`
}

type releasePayLoad struct {
	Nodes []struct {
		Version string `json:"version"`
		Payload string `json:"payload"`
	} `json:"nodes"`
}

func newCluster(opts *clusterOpts) (*Cluster, error) {
	t := time.Now()
	date := t.Format("2006-01-02")
	user, err := user.Current()
	if err != nil {
		return nil, err
	}

	if opts.name == "" {
		opts.name = t.Format("0102150405")
	}

	clusterName := fmt.Sprintf("%s-%s-%s", user.Username, opts.name, date)
	sshKey, err := ioutil.ReadFile(opts.sshKeyPath)
	if err != nil {
		return nil, err
	}

	dir := filepath.Join(opts.baseDir, opts.provider, date, opts.name, t.Format("15_04_05"))
	fmt.Printf("Building cluster in %s\n", dir)
	os.MkdirAll(dir, os.ModePerm)

	cluster := Cluster{
		opts: opts,
		Dir:  dir,
		//Config: confFile,
		TemplateData: TemplateData{
			ClusterName: clusterName,
			SSHKey:      string(sshKey),
			LogLevel:    "debug",
			ClusterDir:  dir,
		},
	}

	cluster.setComputeReplicas()
	cluster.setProviderRegion()
	cluster.setNetworkCidrs()

	// set release image to latest
	if opts.releaseImage == "latest" {
		var minorVersion string
		v := strings.Split(opts.ocpVersion, ".")
		switch {
		case len(v) == 2: // 4.8
			minorVersion = opts.ocpVersion
		case len(v) == 3: // 4.8.1
			minorVersion = fmt.Sprintf("%s.%s", v[0], v[1])
		}
		latest, err := getLatest(opts.releaseImageType, minorVersion)
		if err != nil {
			return nil, err
		}
		fmt.Printf("setting release image to %s latest %s: %s", minorVersion, opts.releaseImageType, latest)
		opts.releaseImage = latest
	}

	// set release image to release version
	if opts.releaseImage == "release" {
		if len(strings.Split(opts.ocpVersion, ".")) != 3 { // 4.8.1
			return nil, fmt.Errorf("release flag with release value must have full version ex: 4.8.1: %s", opts.ocpVersion)
		}
		release, err := getReleaseByVersion(opts.ocpVersion)
		if err != nil {
			return nil, err
		}
		fmt.Printf("setting release %q image to: %q", opts.ocpVersion, release)
		opts.releaseImage = release
	}

	if opts.pullSecret == "" {
		if err := cluster.setPullSecretCI(); err != nil {
			return nil, fmt.Errorf("failed to authenticate try oc login: %v", err)
		}
		if err := cluster.setPullSecretCloud(); err != nil {
			return nil, err
		}
		if err := cluster.setPullSecretDocker(); err != nil {
			return nil, err
		}
	}

	if err := cluster.setPullSecret(); err != nil {
		return nil, err
	}

	if os.Getenv("OPENSHIFT_INSTALL_KEEP_BOOTSTRAP_RESOURCES") != "" {
		opts.keepBootstrap = true
	}
	return &cluster, nil
}

// Run contains the logic of the render command.
func (c *clusterOpts) Run() error {
	cluster, err := newCluster(c)
	if err != nil {
		return err
	}

	//extract installer from release image.
	if err := cluster.extractInstaller(); err != nil {
		return err
	}

	// populate install-config.
	if err := cluster.writeInstallConfig(); err != nil {
		return err
	}

	// populate test docker-compose assets.
	if err := cluster.writeTestAssets(); err != nil {
		return err
	}

	//build cluster.
	if err := cluster.runInstaller(); err != nil {
		return err
	}

	return nil
}

func (c *Cluster) setNetworkCidrs() {
	var (
		clusterCidr           string
		clusterCidrHostPrefix int
		machineCidr           []string
		serviceCidr           string
		networkType           string
	)

	switch c.opts.provider {
	case "azure":
		if c.opts.singleStackIpv6 {
			clusterCidr = "fd01::/48"
			clusterCidrHostPrefix = 64
			machineCidr = []string{"10.0.0.0/16", "fc00::/48"}
			serviceCidr = "fd02::/112"
			networkType = "OVNKubernetes"
		} else {
			clusterCidr = "10.128.0.0/14"
			clusterCidrHostPrefix = 23
			machineCidr = []string{"10.0.0.0/16"}
			serviceCidr = "172.30.0.0/16"
			networkType = "OpenShiftSDN"
		}
	case "gcp":
	case "aws":
	}
	c.TemplateData.ClusterCidr = clusterCidr
	c.TemplateData.ClusterCidrHostPrefix = clusterCidrHostPrefix
	c.TemplateData.MachineCidr = machineCidr
	c.TemplateData.ServiceCidr = serviceCidr
	c.TemplateData.NetworkType = networkType
}

func (c *Cluster) setComputeReplicas() {
	// defaults
	masters := "3"
	workers := "3"

	switch c.opts.provider {
	case "libvirt":
		masters = "1"
		workers = "1"
	}

	if c.opts.replicasMaster != "" {
		masters = c.opts.replicasMaster
	}
	if c.opts.replicasWorker != "" {
		workers = c.opts.replicasWorker
	}
	c.TemplateData.WorkerReplicas = workers
	c.TemplateData.MasterReplicas = masters
}

func (c *Cluster) setPullSecretDocker() error {
	dockerToken := os.Getenv("DOCKER_TOKEN")
	if dockerToken == "" && c.opts.releaseImageType == "custom" {
		return fmt.Errorf("DOCKER_TOKEN is required for custom images")
	}
	if dockerToken == "" {
		klog.Infof("DOCKER_TOKEN not found skipping docker pullSecret...")
		return nil
	}
	// TODO maybe we read from file I am lazy :)
	pullSecretDocker := fmt.Sprintf("{\"auths\":{\"https://index.docker.io/v1/\":{\"auth\":\"%s\"}}}", dockerToken)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(pullSecretDocker), "", "  "); err != nil {
		return err
	}

	// write to disk
	destinationFile := fmt.Sprintf("%s/%s", c.Dir, ".PULL_SECRET_DOCKER")
	if err := ioutil.WriteFile(destinationFile, prettyJSON.Bytes(), 0644); err != nil {
		return err
	}
	c.PullSecrets = append(c.PullSecrets, pullSecretDocker)
	return nil
}

func (c *Cluster) setPullSecretCloud() error {
	// how this works https://access.redhat.com/solutions/4844461
	cloudRedHatToken := os.Getenv("CLOUD_RED_HAT_TOKEN")
	if cloudRedHatToken == "" {
		return fmt.Errorf("unable to get token from ENV CLOUD_RED_HAT_TOKEN, please set and try again")
	}

	var r RedHatCloud
	res, err := http.PostForm(cloudRedHatTokenUrl,
		url.Values{"grant_type": {"refresh_token"}, "client_id": {"cloud-services"}, "refresh_token": {cloudRedHatToken}})
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &r); err != nil {
		return err
	}

	bearer := fmt.Sprintf("Bearer %s", r.AccessToken)
	pullSecretBytes, err := getCloudToken(bearer)
	if err != nil {
		return err
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, pullSecretBytes, "", "  "); err != nil {
		return err
	}

	// write to disk
	destinationFile := fmt.Sprintf("%s/%s", c.Dir, ".PULL_SECRET_CLOUD")
	err = ioutil.WriteFile(destinationFile, prettyJSON.Bytes(), 0644)
	if err != nil {
		return err
	}

	c.PullSecrets = append(c.PullSecrets, string(pullSecretBytes))
	return nil
}

func (c *Cluster) getPullSecretBuild() ([]byte, error) {
	//TODO be more flexible with more than 2
	if len(c.PullSecrets) < 2 {
		return nil, fmt.Errorf("2 pullSecrets required for merge")
	}

	var src, pullSecretBuildBytes []byte

	for i, _ := range c.PullSecrets {
		if i == 0 {
			// skip we are doing N-1
			continue
		}
		var err error
		if i == 1 {
			// seed with 0
			src = []byte(c.PullSecrets[0])
		} else {
			// use combined
			src = pullSecretBuildBytes
		}
		pullSecretBuildBytes, err = jsonm.MergeMergePatches(src, []byte(c.PullSecrets[i]))
		if err != nil {
			return nil, err
		}
	}

	return pullSecretBuildBytes, nil
}

func (c *Cluster) setPullSecret() error {
	pullSecretBuildBytes, err := c.getPullSecretBuild()
	if err != nil {
		return fmt.Errorf("1 %v", err)
		// return err
	}
	// fmt.Fprintf("found pullSecret %s", pullSecretBuildBytes)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, pullSecretBuildBytes, "", "  "); err != nil {
		return err
	}

	// write to disk
	destinationFile := fmt.Sprintf("%s/%s", c.Dir, ".PULL_SECRET_BUILD")
	if err := ioutil.WriteFile(destinationFile, prettyJSON.Bytes(), 0644); err != nil {
		return err
	}

	// compress json
	pullSecret := new(bytes.Buffer)
	if err := json.Compact(pullSecret, pullSecretBuildBytes); err != nil {
		return err
	}
	c.TemplateData.PullSecret = pullSecret.String()
	return nil
}

func getCloudToken(bearerToken string) ([]byte, error) {
	req, err := http.NewRequest("POST", cloudRedHatAccessTokenUrl, nil)
	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	pullSecret := new(bytes.Buffer)
	if err := json.Compact(pullSecret, body); err != nil {
		return nil, err
	}
	return pullSecret.Bytes(), nil
}

func (c *Cluster) setProviderRegion() {
	var region string

	//defaults
	//TODO allow override in config
	switch c.opts.provider {
	case "aws":
		region = "us-west-1"
	case "gcp":
		region = "us-east1"
	case "azure":
		region = "eastus"
	}

	if c.opts.providerRegion != "" {
		region = c.opts.providerRegion
	}
	c.TemplateData.ProviderRegion = region
}
func (c *Cluster) setPullSecretCI() error {
	kubeConfigFlags := genericclioptions.NewConfigFlags(true)
	matchVersionKubeConfigFlags := kcmdutil.NewMatchVersionFlags(kubeConfigFlags)
	f := kcmdutil.NewFactory(matchVersionKubeConfigFlags)
	pullPath := fmt.Sprintf("%s/%s", c.Dir, ".PULL_SECRET_CI")
	o := &login.LoginOptions{
		ConfigFile: pullPath, // "-", prints stdout
		IOStreams: genericclioptions.IOStreams{
			In:     os.Stdin,
			Out:    os.Stdout,
			ErrOut: os.Stderr,
		},
	}

	if err := o.Complete(f, []string{""}); err != nil {
		return fmt.Errorf("setPullSecretCI() %s", err)
	}
	if err := o.Run(); err != nil {
		return fmt.Errorf("setPullSecretCI() %s", err)
	}

	raw, err := ioutil.ReadFile(pullPath)
	if err != nil {
		return fmt.Errorf("setPullSecretCI() %s", err)
	}

	pullSecret := new(bytes.Buffer)
	if err := json.Compact(pullSecret, raw); err != nil {
		return fmt.Errorf("setPullSecretCI() %s", err)
	}
	c.PullSecrets = append(c.PullSecrets, pullSecret.String())

	return nil
}

func (c *Cluster) extractInstaller() error {
	o := &release.ExtractOptions{
		Directory: fmt.Sprintf("%s/%s", c.Dir, "bin"),
		IOStreams: genericclioptions.IOStreams{
			In:     os.Stdin,
			Out:    os.Stdout,
			ErrOut: os.Stderr,
		},
		Command: "openshift-install",
		From:    c.opts.releaseImage,
		SecurityOptions: imagemanifest.SecurityOptions{
			RegistryConfig: fmt.Sprintf("%s/%s", c.Dir, ".PULL_SECRET_BUILD"),
		},
	}
	if err := o.Run(); err != nil {
		return err
	}
	return nil
}

func (c *Cluster) initCustomInstaller() error {
	gitPath := fmt.Sprintf("%s/src/github.com/openshift/installer", c.Dir)
	args := []string{"clone", openShiftInstallerUrl, gitPath}
	if _, err := exec.Command("git", args...).Output(); err != nil {
		return fmt.Errorf("initCustomInstaller() %s", err)
	}

	commit, err := extractInstallerCommmit(c.Dir)
	if err != nil {
		return fmt.Errorf("initCustomInstaller() %s", err)
	}

	fmt.Printf("Checking out installer commit %s", commit)

	cmd := fmt.Sprintf("cd %s ;git checkout %s", gitPath, commit)
	if _, err := exec.Command("bash", "-c", cmd).Output(); err != nil {
		fmt.Printf("installer commit %s not found using master", commit)
	}

	return nil
}

func (c *Cluster) patchCustomInstaller(patch []byte) error {
	f, err := os.Create(fmt.Sprintf("%s/src/github.com/openshift/installer/patch", c.Dir))
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(string(patch)); err != nil {
		return fmt.Errorf("patchCustomInstaller() %s", err)
	}

	//TODO tried git-go but really wasn't working well maybe revist?
	args := []string{"apply", "--reject", "--whitespace=fix", "patch"}

	cmd := exec.Command("git", args...)
	cmd.Dir = fmt.Sprintf("%s/src/github.com/openshift/installer", c.Dir)
	stdout, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("patchCustomInstaller() %s", err)
	}

	fmt.Printf("patching installer\n%s", stdout)

	return nil
}

func (c *Cluster) buildInstallCustomInstaller() error {
	args := []string{""}
	cmd := exec.Command(fmt.Sprintf("%s/src/github.com/openshift/installer/hack/build.sh", c.Dir), args...)
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("GOPATH=%s", c.Dir),
		"GO111MODULE=off",
		fmt.Sprintf("OUTPUT=%s/bin", c.Dir),
	)

	//TODO create func
	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("cmd.Start() failed with '%s'\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
		wg.Done()
	}()

	stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)

	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		return fmt.Errorf("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdout), string(stderr)
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

	return err
}

func extractInstallerCommmit(dir string) (string, error) {
	installerPath := fmt.Sprintf("%s/%s", dir, "bin/openshift-install")
	cmd := fmt.Sprintf("%s version 2> /dev/null | grep commit | awk '{ print $4 }'", installerPath)
	stdout, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", fmt.Errorf("extractInstallerCommmit() failed with %s\n", err)
	}
	return string(stdout), nil
}

func getConfigFile() (config.File, error) {
	var c config.File
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return c, err
	}
	configPath := fmt.Sprintf("%s/.config/dev-installer/config.yaml", homeDir)
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return c, err
	}

	err = yaml.Unmarshal(configFile, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (c *Cluster) writeInstallConfig() error {
	src := template_assets.MustAsset(fmt.Sprintf("bindata/templates/installer/%s/install-config.yaml", c.opts.provider))
	tpl, err := template.New("install-config").Parse(string(src))
	//tpl, err := template.ParseFiles(string(src))
	if err != nil {
		return fmt.Errorf("writeInstallConfig() failed with %s\n", err)
	}
	//
	out, err := os.Create(fmt.Sprintf("%s/%s", c.Dir, "install-config.yaml"))
	if err != nil {
		return fmt.Errorf("writeInstallConfig() failed with %s\n", err)
	}
	defer out.Close()

	err = tpl.Execute(out, c.TemplateData)
	if err != nil {
		return fmt.Errorf("writeInstallConfig() failed with %s\n", err)
	}
	return nil
}

func (c *Cluster) writeTestAssets() error {
	tpl, err := template.ParseFiles(fmt.Sprintf("./templates/tests/origin/%s", "docker-compose.yaml"))

	if err != nil {
		return fmt.Errorf("writeTestAssets() failed with %s\n", err)
	}

	out, err := os.Create(fmt.Sprintf("%s/%s", c.Dir, "docker-compose.yaml"))
	if err != nil {
		return fmt.Errorf("writeTestAssets() failed with %s\n", err)
	}
	defer out.Close()

	err = tpl.Execute(out, c.TemplateData)
	if err != nil {
		return fmt.Errorf("writeTestAssets() failed with %s\n", err)
	}
	return nil
}
func (c *Cluster) runInstaller() error {
	installerPath := fmt.Sprintf("%s/%s", c.Dir, "bin/openshift-install")
	if c.opts.installerPath != "" {
		if err := os.Remove(installerPath); err != nil {
			return fmt.Errorf("runInstaller() failed with %s\n", err)
		}
		err := os.Symlink(c.opts.installerPath, installerPath)
		if err != nil {
			return fmt.Errorf("runInstaller() failed with %s\n", err)
		}
	}

	args := []string{"create", "cluster", "--dir", c.Dir, "--log-level", c.LogLevel}

	cmd := exec.Command(installerPath, args...)

	cmd.Env = append(os.Environ(),
		fmt.Sprintf("OPENSHIFT_INSTALL_RELEASE_IMAGE_OVERRIDE=%s", c.opts.releaseImage),
		fmt.Sprintf("OPENSHIFT_INSTALL_AZURE_EMULATE_SINGLESTACK_IPV6=%v", c.opts.singleStackIpv6), //TODO this should be a provider level config
		fmt.Sprintf("OPENSHIFT_INSTALL_PRESERVE_BOOTSTRAP=%v", c.opts.keepBootstrap),
	)

	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	}

	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("cmd.Start() failed with '%s'\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
		wg.Done()
	}()

	stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		return fmt.Errorf("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdout), string(stderr)
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
	return nil
}

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

func getReleaseByVersion(version string) (string, error) {
	var r releasePayLoad
	resp, err := http.Get("https://amd64.ocp.releases.ci.openshift.org/graph")
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(body, &r); err != nil {
		return "", fmt.Errorf("failed to unmarshal: %v", err)
	}
	for _, release := range r.Nodes {
		if release.Version == version {
			return release.Payload, nil
		}
	}
	return "", fmt.Errorf("no release payload found for version: %s", version)
}

func getLatest(releaseImageType, release string) (string, error) {
	return httpGetRelease(fmt.Sprintf("https://amd64.ocp.releases.ci.openshift.org/api/v1/releasestream/%s.0-0.%s/latest", release, releaseImageType))
}

func httpGetRelease(release string) (string, error) {
	var r releaseStream
	resp, err := http.Get(release)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(body, &r); err != nil {
		return "", fmt.Errorf("failed to unmarshal: %v", err)
	}
	return r.PullSpec, nil
}
