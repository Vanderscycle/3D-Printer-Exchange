# Welcome to Tilt!
#   To get you started as quickly as possible, we have created a
#   starter Tiltfile for you.
#
#   Uncomment, modify, and delete any commands as needed for your
#   project's configuration.
modes = ['localhost', 'infrastructure']
selection = modes[0]


# Variables
sync_src_frontend= sync('./frontend', '/src')
sync_src_backend= sync('./backend', '/src')

# Build Docker image
#   Tilt will automatically associate image builds with the resource(s)
#   that reference them (e.g. via Kubernetes or Docker Compose YAML).
#
#   More info: https://docs.tilt.dev/api.html#api.docker_build
#

docker_build('localhost:5005/frontend-sveltekit', context='./frontend', dockerfile='./frontend/Dockerfile', live_update=[sync_src_frontend] )
docker_build('localhost:5005/backend-fiber',context='./backend',dockerfile='./backend/Dockerfile', live_update=[sync_src_backend])

# Extensions are open-source, pre-packaged functions that extend Tilt
#
#   More info: https://github.com/tilt-dev/tilt-extensions
#
load('ext://helm_remote', 'helm_remote')


def localhost():
    print("""
    -----------------------------------------------------------------
    ✨ Localhost Environment
    -----------------------------------------------------------------
    """.strip())
    
    update_settings(suppress_unused_image_warnings=["localhost:5005/frontend-sveltekit"])
    update_settings(suppress_unused_image_warnings=["localhost:5005/backend-fiber"])

    # Local ressources
    local_resource('localhost-postgres',
    serve_dir='./backend/scripts/',
    serve_cmd='bash start-dev-db.sh',
    # cmd='make dev-db'

    )
    local_resource('localhost-backend',
    resource_deps=['localhost-postgres'],
    serve_dir='./backend',
    serve_cmd='go run main.go',
    deps='./backend/src',
    readiness_probe=probe(
        period_secs=60,
        http_get=http_get_action(port=3001, path="/version")
        )
    )
    local_resource('localhost-frontend',
    resource_deps=['localhost-postgres'],
    serve_dir='./frontend',
    serve_cmd='cd pnpm run dev',
    deps='./frontend/pages',
    readiness_probe=probe(
        period_secs=60,
        http_get=http_get_action(port=5173, path="/")
        )
    )
localhost()
