k8s_yaml(kustomize('deploy/kubernetes'))

# datastore

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

k8s_resource(
  workload="minio",
  port_forwards='9001:9001',
  labels='datastore'
)

# ingest lore

docker_build(
'datastore-ingestion',
'.',
dockerfile='deploy/docker/Dockerfile.datastore.ingestion',
)

k8s_resource(
  workload="lore-initial-job",
  labels="datastore",
  resource_deps=["postgres", "db-migrations", "text-embeddings-deployment"]
)


# narrative

docker_build(
'narrative-prompt',
'.',
dockerfile='deploy/docker/Dockerfile.narrative.prompt'
)

k8s_resource(
  workload='narrative-prompt',
  labels='narrative',
  port_forwards=["50051:50051"]
)

# embeddings

k8s_resource(
  workload='text-embeddings-deployment',
  labels='text-embeddings',
  port_forwards='3000:3000',
  resource_deps=[]
)
