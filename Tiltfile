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
    db_port = 5432
    backend_port = 8080
    frontend_port = 5173
    print("""
    -----------------------------------------------------------------
    ✨ Localhost Environment
    -----------------------------------------------------------------
    """.strip())
    
    update_settings(suppress_unused_image_warnings=["localhost:5005/frontend-sveltekit"])
    update_settings(suppress_unused_image_warnings=["localhost:5005/backend-fiber"])

    # ------------ db (postgresql) ------------
    docker_compose("./backend/database/docker-compose.yml")
    
    # ------------ backend (goFiber) ------------
    local_resource('localhost-backend',
    serve_dir='./backend',
    serve_cmd='go run main.go',
    deps='./backend/src',
    readiness_probe=probe(
        period_secs=60,
        http_get=http_get_action(port=backend_port, path="/ancillary/version/")
        )
    )

    # ------------ frontend (sveltekit) ------------
    local_resource('localhost-frontend',
    serve_dir='./frontend',
    serve_cmd='pnpm run dev',
    deps='./frontend/pages',
    readiness_probe=probe(
        period_secs=60,
        http_get=http_get_action(port=frontend_port, path="/")
        )
    )
if selection == 'localhost':
    localhost()
elif selection == 'infrastructure':
    pass
