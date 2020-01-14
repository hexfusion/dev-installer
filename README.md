# dev-installer
tooling for creating and managing OCP 4 clusters


#### CI images require you to be logged in to ci via `oc login`

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
