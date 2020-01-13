package cluster

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"io"
	"io/ioutil"
	"k8s.io/klog"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"text/template"
	"time"
	"os/exec"
	"runtime"
	"sync"

	"github.com/hexfusion/dev-installer/pkg/cluster/admin/release"
	imagemanifest "github.com/hexfusion/dev-installer/pkg/cluster/image/manifest"
	"github.com/hexfusion/dev-installer/pkg/cluster/registry"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	kcmdutil "k8s.io/kubectl/pkg/cmd/util"
)


// clusterOpts holds values to drive the cluster command.
type clusterOpts struct {
	errOut io.Writer
	provider string
	pullSecret string
	releaseImage string
	releaseImageType string
	installerPath string
	keepBootstrap string
	baseDir string
	name string
	pullSecretName string
	sshKeyPath string
}

const (
	cloudRedHatTokenUrl = "https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token"
	// https://cloud.redhat.com/beta/openshift/token
	cloudRedHatToken = ""
    quayRedhatToken = ""
	quayRedhatEmail = ""
    )

// NewClusterCommand creates a new cluster
func NewClusterCommand(errOut io.Writer) *cobra.Command {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		klog.Fatal(err)
	}

	clusterOpts := clusterOpts{
		errOut:   errOut,
		baseDir: fmt.Sprintf("%s/clusters",homeDir),
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
	fs.StringVar(&c.pullSecret, "pull-secret", c.pullSecret, "pull secret to use for cluster creation")
	fs.StringVarP(&c.releaseImage, "release", "r", c.releaseImage, "release image")
	fs.StringVarP(&c.releaseImageType, "release-type", "t", c.releaseImageType, "the type of release image used. Example CI, Nightly, Custom")
	fs.StringVar(&c.installerPath, "installer-path", c.installerPath, "path of the compiled installer to use")
	fs.StringVar(&c.baseDir, "base-dir", c.baseDir, "path of the base dir to store cluster data")
	fs.StringVarP(&c.sshKeyPath, "ssh-path", "s", c.sshKeyPath, "path to public ssh key for cluster")
	//fs.StringVar(&c.keepBootstrap, "keep-bootstrap", r.keepBootstrap, "keep boostrap node around for debug")
}

// Validate verifies the inputs.
func (c *clusterOpts) Validate() error {
	if len(c.provider) == 0 {
		return errors.New("missing required flag: --provider -p")
	}
	if len(c.releaseImage) == 0 {
		return errors.New("missing required flag: --release -r")
	}
	//TODO parse from image name
	if len(c.releaseImageType) == 0 {
		return errors.New("missing required flag: --release-type -rt")
	}
	if len(c.sshKeyPath) == 0 {
		return errors.New("missing required flag: --ssh-path -s")
	}
	if len(c.name) == 0 {
		return errors.New("missing required flag: --name -n")
	}
	return nil
}

type Auth struct {
	Type string
	FileName string
}

type Cluster struct {
	opts *clusterOpts
	PullSecrets []Auth
	Dir string
	TemplateData
}

type RedHatCloud struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
	Created string
}

type TemplateData struct {
	SSHKey string
	PullSecret string
	ClusterName string
	WorkerReplicas int
	MasterReplicas int
	LogLevel string
}

//type Auths struct {
//	CloudOpenshift AuthContent `json:"registry.connect.redhat.com"`
//	Quay AuthContent `json:"quay.io"`
//	RegistryConnectRedhat AuthContent `json:"registry.connect.redhat.com"`
//	RegistryRedhat AuthContent `json:"registry.redhat.io"`
//}
//
//type AuthContent struct {
//	Auth string `json:"email"`
//	Email string `json:"auth"`
//}

func newCluster(opts *clusterOpts) (*Cluster, error) {
	if opts.baseDir == "" {

	}

	t := time.Now()
	date := t.Format("2006-01-02")
	user, err := user.Current()
	if err != nil {
		return nil, err
	}

	clusterName := fmt.Sprintf("%s-%s-%s",  user.Username, opts.name, date)
	sshKey, err := ioutil.ReadFile(opts.sshKeyPath)
	if err != nil {
		return nil, err
	}

	dir := fmt.Sprintf("%s/%s/%s/%s-%s", opts.baseDir, opts.provider, date, opts.name, t.Format("15_04_05"))
	os.MkdirAll(dir, os.ModePerm)
	cluster := Cluster{
		opts: opts,
		Dir: dir,
		TemplateData: TemplateData{
			ClusterName: clusterName,
			SSHKey: string(sshKey),
			WorkerReplicas: 3,
			MasterReplicas: 3,
			LogLevel: "debug",
		},
	}

	if err := cluster.setPullSecret(); err != nil {
		return nil, err
	}

	if err := cluster.setPullSecretCI(); err != nil {
			return nil, err
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

	//build cluster.
	if err :=  cluster.runInstaller(); err != nil {
		return err
	}

	return nil
}

func (t *Cluster) setPullSecret() error {
	var r RedHatCloud
	res, err := http.PostForm(cloudRedHatTokenUrl,
		url.Values{"grant_type": {"refresh_token"}, "client_id": {"cloud-services"}, "refresh_token": {cloudRedHatToken}})
	if err != nil {
		return err
	}
	if err := json.NewDecoder(res.Body).Decode(&r);err != nil {
		return err
	}
//	t.PullSecrets = r.AccessToken
	return nil
}

func (c *Cluster) setPullSecretCI() error {
	kubeConfigFlags := genericclioptions.NewConfigFlags(true)
	matchVersionKubeConfigFlags := kcmdutil.NewMatchVersionFlags(kubeConfigFlags)
	f := kcmdutil.NewFactory(matchVersionKubeConfigFlags)
	pullPath := fmt.Sprintf("%s/%s", c.Dir, "CI_PULL_SECRET")
	o := &registry.LoginOptions{
		ConfigFile: pullPath, // "-", prints stdout
		IOStreams: genericclioptions.IOStreams{
			In:     os.Stdin,
			Out:    os.Stdout,
			ErrOut: os.Stderr,
		},
	}

	if err := o.Complete(f, []string{""}); err != nil {
		return err
	}
	if err := o.Run(); err != nil {
		return err
	}

	raw, err := ioutil.ReadFile(pullPath)
	if err != nil {
		return err
	}

	pullSecret := new(bytes.Buffer)
	if err := json.Compact(pullSecret, raw);err != nil {
		return err
	}
	c.TemplateData.PullSecret = pullSecret.String()

	return nil
}

func (c *Cluster) extractInstaller() error {
	o := &release.ExtractOptions{
		Directory: fmt.Sprintf("%s/%s", c.Dir,"bin"),
		IOStreams: genericclioptions.IOStreams{
			In:     os.Stdin,
			Out:    os.Stdout,
			ErrOut: os.Stderr,
		},
		Command: "openshift-install",
		From: c.opts.releaseImage,
		SecurityOptions: imagemanifest.SecurityOptions{
			RegistryConfig: fmt.Sprintf("%s/%s", c.Dir, "CI_PULL_SECRET"),
		},
	}
	if err := o.Run(); err != nil {
		return err
	}
	return nil
}

func (c *Cluster) writeInstallConfig() error {
	tpl, err := template.ParseFiles(fmt.Sprintf("./templates/installer/%s/%s", c.opts.provider, "install-config.yaml"))

	if err != nil {
		return err
	}

	out, err := os.Create(fmt.Sprintf("%s/%s", c.Dir, "install-config.yaml"))
	if err != nil {
		return err
	}
	defer out.Close()

	err = tpl.Execute(out, c.TemplateData)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cluster) runInstaller() error {
	installerPath := fmt.Sprintf("%s/%s", c.Dir, "bin/openshift-install")
	args := []string{"create", "cluster", "--dir", c.Dir, "--log-level", c.LogLevel}

	cmd := exec.Command(installerPath, args...)
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