module github.com/hexfusion/dev-installer

go 1.13

require (
	github.com/MakeNowJust/heredoc v0.0.0-20170808103936-bb23615498cd
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/blang/semver v3.5.0+incompatible
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v0.7.3-0.20190817195342-4760db040282
	github.com/docker/go-units v0.4.0
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7
	github.com/evanphx/json-patch v4.2.0+incompatible
	github.com/ghodss/yaml v1.0.0
	github.com/gotestyourself/gotestyourself v2.2.0+incompatible // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1
	github.com/openshift/api v0.0.0-20191217141120-791af96035a5
	github.com/openshift/client-go v0.0.0-20191216194936-57f413491e9e
	github.com/openshift/library-go v0.0.0-20200108105826-2cb27e8b3c7b
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	golang.org/x/crypto v0.0.0-20190820162420-60c769a6c586
	golang.org/x/sys v0.0.0-20190826190057-c7b8b68b1456
	k8s.io/api v0.17.0
	k8s.io/apimachinery v0.17.0
	k8s.io/apiserver v0.17.0
	k8s.io/cli-runtime v0.17.0
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/component-base v0.17.0
	k8s.io/klog v1.0.0
	k8s.io/kubectl v0.17.0
	k8s.io/kubernetes v0.0.0-00010101000000-000000000000
	sigs.k8s.io/yaml v1.1.0
)

replace (
	github.com/docker/docker => github.com/docker/docker v0.0.0-20180612054059-a9fbbdc8dd87
	github.com/openshift/client-go => github.com/openshift/client-go v0.0.0-20191216194936-57f413491e9e
	k8s.io/api => github.com/kubernetes/api v0.0.0-20191121175643-4ed536977f46
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20191204090830-8d4ebf9010bd
	k8s.io/apimachinery => github.com/openshift/kubernetes-apimachinery v0.0.0-20191211181342-5a804e65bdc1
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20191204085032-7fb3a25c3bc4
	k8s.io/cli-runtime => github.com/openshift/kubernetes-cli-runtime v0.0.0-20191211181810-5b89652d688e
	k8s.io/client-go => github.com/openshift/kubernetes-client-go v0.0.0-20191211181558-5dcabadb2b45
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20191121182543-b8af8c87a0d2
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20191121182434-3459e7278621
	k8s.io/code-generator => k8s.io/code-generator v0.17.1-beta.0
	k8s.io/component-base => k8s.io/component-base v0.0.0-20191128032904-4bcd454928ff
	k8s.io/cri-api => k8s.io/cri-api v0.17.1-beta.0
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20191121182650-d032a1f882e1
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20191121180901-7ce2d4f093e4
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20191121182328-3e5a379d6404
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20191121182004-c1057c1a0821
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20191121182219-f40ec664a26f
	k8s.io/kubectl => github.com/openshift/kubernetes-kubectl v0.0.0-20200109100530-0dbab4a25283
	k8s.io/kubelet => k8s.io/kubelet v0.0.0-20191121182112-95f295975fc9
	k8s.io/kubernetes => github.com/openshift/kubernetes v1.17.0-alpha.0.0.20191216151305-079984b0a154
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20191215115203-1896ee2ad49b
	k8s.io/metrics => k8s.io/metrics v0.0.0-20191121181631-c7d4ee0ffc0e
	k8s.io/node-api => k8s.io/node-api v0.0.0-20191121182916-ad54f283563d
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20191121181040-36c9528858d2
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.0.0-20191121181855-541d7bb23c26
	k8s.io/sample-controller => k8s.io/sample-controller v0.0.0-20191121181305-e6c211291103
)
