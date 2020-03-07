package cluster

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"testing"
)

var (
	testClusterConfig = `
auths:
   - name: "quay.io"
     email: ""
     auth: "example"
ssh:
   - publicKeyPath: "` + os.UserHomeDir() + `/.ssh/libra.pub"
`
)

type testConfig struct {
	t                    *testing.T
	clusterNetworkConfig string
	want                 Cluster
}

func TestClusterCI(t *testing.T) {
	want := Cluster{}
	config := &testConfig{
		t:    t,
		want: want,
	}
	testCluster(config)
}

func testCluster(tc *testConfig) {
	var errOut io.Writer

	clusterDir, err := ioutil.TempDir("/tmp", "clusters-")
	if err != nil {
		tc.t.Fatal(err)
	}

	clusterOpts := &clusterOpts{
		name:			  "test-cluster",
		provider:         "aws",
		releaseImage:     "registry.svc.ci.openshift.org/ocp/release:4.1.0-0.ci-2020-01-06-225050",
		releaseImageType: "ci",
		baseDir: clusterDir,
		errOut:           errOut,
		sshKeyPath:       os.UserHomeDir() + "/.ssh/libra.pub",
	}

	cluster, err := newCluster(clusterOpts)
	if err != nil {
		tc.t.Fatal(err)
	}

	if err := cluster.extractInstaller(); err != nil {
		tc.t.Fatal(err)
	}

	if err := cluster.writeInstallConfig(); err != nil {
		tc.t.Fatal(err)
	}

	fmt.Printf("dump got %s", cluster.PullSecrets)
}
