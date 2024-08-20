k8s_yaml(kustomize('deploy/kubernetes'))

docker_build(
  'datastore-migrations',
  '.',
  dockerfile='deploy/docker/Dockerfile.migrations',
)

k8s_resource(
  workload='db-migrations',
  labels='datastore',
  resource_deps=['postgres']
)

k8s_resource(
  workload='postgres',
  port_forwards='5432:5432',
  labels='datastore'
)
