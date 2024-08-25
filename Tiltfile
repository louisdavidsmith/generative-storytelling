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
  port_forwards=['9000:9000','9001:9001'],
  labels='datastore'
)

# ingest lore

docker_build(
'datastore-ingestion',
'.',
dockerfile='deploy/docker/Dockerfile.datastore.ingestion',
)

k8s_resource(
  workload="datastore-ingestion",
  labels="datastore",
  resource_deps=["postgres", "db-migrations", "text-embeddings-deployment"]
)

# rabbit mq

k8s_resource(
  workload="rabbitmq",
  labels="rabbitmq",
  port_forwards="5672:5672"
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

docker_build(
'narrative-generation',
'.',
dockerfile='deploy/docker/Dockerfile.narrative.generation'
)

k8s_resource(
  workload='narrative-generation',
  labels='narrative',
  port_forwards=["50002:50002"]
)

# embeddings

k8s_resource(
  workload='text-embeddings-deployment',
  labels='text-embeddings',
  port_forwards='3000:3000',
  resource_deps=[]
)
