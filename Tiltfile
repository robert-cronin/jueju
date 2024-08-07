load('ext://restart_process', 'docker_build_with_restart')
load('ext://secret', 'secret_create_generic')
load('ext://namespace', 'namespace_create', 'namespace_inject')

# ================== Safety ==================
# don't allow any context except "minikube"
allow_k8s_contexts('minikube')
if k8s_context() != 'minikube':
  fail("failing early, needs context called 'minikube'")

docker_prune_settings(num_builds=2)

namespace_create('jueju')
k8s_resource(
    objects=['jueju:namespace'],
    new_name='jueju-namespace',
    labels=['chart'],
)

# Backend
docker_build(
    'ghcr.io/robert-cronin/jueju:backend-latest',
    './backend',
    entrypoint='scripts/start_tilt.sh build/main',
    dockerfile='./backend/Dockerfile',
    live_update=[
        sync('./backend/build', '/app/build'),
    ],
    target='dev'
)
local_resource(
    'build-backend',
    'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/main .',
    dir='./backend',
)
k8s_yaml('./deploy/backend.yaml')
k8s_resource('jueju-backend', port_forwards='3000:3000')

local_resource(
    'backend-lint',
    'golangci-lint run ./...',
    dir='./backend',
    deps=['backend'],
    labels=['dev']
)

# Frontend
docker_build(
    'ghcr.io/robert-cronin/jueju:frontend-latest',
    './frontend',
    dockerfile='./frontend/Dockerfile',
    live_update=[
        sync('./frontend/', '/app/'),
    ],
    target='dev'
)
k8s_yaml('./deploy/frontend.yaml')
k8s_resource('jueju-frontend', port_forwards='5173:5173')

# ============== Dev resources ==============
local_resource(
    'frontend-lint',
    'yarn lint',
    dir='./frontend',
    deps=['frontend'],
    labels=['dev']
)

# ============== Helm ==============
yaml = helm(
  './chart',
  name='jueju',
  namespace='kubemedic',
  values=['./chart/values.yaml'],
)
k8s_yaml(yaml)

# ============== Secrets ==============
secret_create_generic(
    'jueju',
    {
        'POSTGRES_PASSWORD': 'letmein',
    },
    namespace='jueju',
)


# ============== Utils ==============
k8s_yaml('./tilt/postgres.yaml')
k8s_yaml('./tilt/redis.yaml')
