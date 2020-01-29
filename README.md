# dev-installer
tooling for creating and managing OCP 4 clusters

### CI images require you to be logged in to ci via `oc login`

##### create cluster based on CI generated release image and patch installer to keep bootstrap node with --keep-bootstrap flag.

```bash
./bin/dev-installer cluster \
  -p gcp \
  -r registry.svc.ci.openshift.org/ci-op-z0qywyr4/release:latest \
  -t ci \
  -s ~/.ssh/libra.pub \
  --keep-bootstrap \
  -n fix-operator-1
```

```bash
bin/dev-installer cluster \
  -p aws \
  -r registry.svc.ci.openshift.org/ocp/release:4.3.0-0.ci-2020-01-10-150540 \
  -t ci \
  -s ~/.ssh/libra.pub \
  -n cluster-test 
```

#### Custom pull-secret

```bash
./bin/dev-installer cluster \
  -p aws \
  -r docker.io/hexfusion/origin-release:v4.4 \
  -t custom \
  -s ~/.ssh/libra.pub \
  --pull-secret=/home/remote/sbatsche/.PULL_SECRET_BUILD \
  -n test-cluster
  ```

#### libvirt requires a precompiled installer with libvirt support pass that with --installer-path

```bash
./bin/dev-installer cluster \
  -p libvirt \
  -r docker.io/hexfusion/origin-release:v4.4 \
  -t custom \
  -s ~/.ssh/libra.pub \
  --pull-secret=/home/remote/sbatsche/.PULL_SECRET_BUILD \
  --installer-path=/home/remote/sbatsche/projects/openshift/installer/bin/openshift-install \
  -n test-cluster
  ```
