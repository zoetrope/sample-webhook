load('ext://restart_process', 'docker_build_with_restart')

DOCKERFILE = '''FROM golang:alpine
WORKDIR /
COPY ./bin/sample-webhook /
CMD ["/sample-webhook"]
'''

# Generate manifests and go files
local_resource('make manifests', "make manifests", deps=["hooks"], ignore=['*/*/zz_generated.deepcopy.go'])
local_resource('make generate', "make generate", deps=["hooks"], ignore=['*/*/zz_generated.deepcopy.go'])

# Deploy manager
watch_file('./config/')
k8s_yaml(kustomize('./config/dev'))

local_resource(
    'Watch & Compile', "make build", deps=['hooks', 'main.go'],
    ignore=['*/*/zz_generated.deepcopy.go'])

docker_build_with_restart(
    'sample-webhook:dev', '.',
    dockerfile_contents=DOCKERFILE,
    entrypoint=['/sample-webhook'],
    only=['./bin/sample-webhook'],
    live_update=[
        sync('./bin/sample-webhook', '/sample-webhook'),
    ]
)
