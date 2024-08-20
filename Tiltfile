k8s_yaml(kustomize('deploy/kubernetes'))

k8s_resource(
  workload='postgres',
  port_forwards='5432:5432',
  labels='datastore'
)
