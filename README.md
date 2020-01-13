# dev-installer
tooling for creating and managing OCP 4 clusters

```bash
bin/dev-installer cluster \
  -p aws \
  -r registry.svc.ci.openshift.org/ocp/release:4.3.0-0.ci-2020-01-10-150540 \
  -t ci \
  -s ~/.ssh/libra.pub \
  -n cluster-test 
```
