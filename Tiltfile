load('ext://restart_process', 'docker_build_with_restart')
load('ext://secret', 'secret_create_generic')

# Test kube context for kind
allow_k8s_contexts('kind-kind')

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
