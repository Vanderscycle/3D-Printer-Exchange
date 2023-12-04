# Custom UI
load('ext://uibutton', 'cmd_button', 'location', 'text_input')
cmd_button(name='frontend-dev',
          text='Frontend Dev',
          icon_name='bolt',
          location=location.NAV,
          argv=['tilt', 'patch', 'tiltfile', '(Tiltfile)', '--patch', '{"spec": {"args": ["--frontend=dev"]}}'])
cmd_button(name='frontend-prod',
          text='Frontend Prod',
          icon_name='local_shipping',
          location=location.NAV,
          argv=['tilt', 'patch', 'tiltfile', '(Tiltfile)', '--patch', '{"spec": {"args": ["--frontend=prod"]}}'])
local_resource('choose-a-mode', ['echo', """Must specify 'tilt up -- --frontend=dev' or 'tilt up -- --frontend=prod'
 Click the bolt for --frontend=dev.
 Click the truck for --frontend=prod.
 """])

# Variables
sync_src_frontend= sync('./frontend', '/src')
sync_src_backend= sync('./backend', '/')
db_port = 5432
backend_port = 8080
frontend_port = 5173

# Build Docker image
#   Tilt will automatically associate image builds with the resource(s)
#   that reference them (e.g. via Kubernetes or Docker Compose YAML).
#
#   More info: https://docs.tilt.dev/api.html#api.docker_build
#

docker_build('localhost:5000/frontend-sveltekit', context='./frontend', dockerfile='./frontend/Dockerfile', live_update=[sync_src_frontend] )
docker_build('localhost:5000/backend-fiber',context='./backend',dockerfile='./backend/Dockerfile', live_update=[sync_src_backend])

# Extensions are open-source, pre-packaged functions that extend Tilt
#
#   More info: https://github.com/tilt-dev/tilt-extensions
#
load('ext://helm_remote', 'helm_remote')

# Custom UI
load('ext://uibutton', 'cmd_button', 'location', 'text_input')
cmd_button(name='nav-hello-world',
           argv=['echo', 'Hello nav!'],
           text='Hello World',
           location=location.NAV,
           icon_name='waving_hand')

def infrastructure():
  print("""
  -----------------------------------------------------------------
  ✨ Kubernetes/Infrastructure Environment
  -----------------------------------------------------------------
  """.strip())
  # Apply Kubernetes manifests
  #   Tilt will build & push any necessary images, re-deploying your
  #   resources as they change.
  #
  #   More info: https://docs.tilt.dev/api.html#api.k8s_yaml
  k8s_fullstack="./manifests/overlays/non-prod"
  k8s_yaml([kustomize(k8s_fullstack, flags=['--enable-helm'])])

  # Customize a Kubernetes resource
  #   By default, Kubernetes resource names are automatically assigned
  #   based on objects in the YAML manifests, e.g. Deployment name.
  #
  #   Tilt strives for sane defaults, so calling k8s_resource is
  #   optional, and you only need to pass the arguments you want to
  #   override.
  #
  #   More info: https://docs.tilt.dev/api.html#api.k8s_resource
  #
  k8s_resource('sveltekit-frontend',labels="frontend",port_forwards='3000:3000')
  k8s_resource('gofiber-backend',labels="backend",port_forwards='8080:8080')
  k8s_resource('db-postgresql',labels="db")


def localhost():
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
    serve_env={},
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
    links=['http://localhost:{}'.format(frontend_port)],
    serve_env={},
    readiness_probe=probe(
        period_secs=60,
        http_get=http_get_action(port=frontend_port, path="/")
        )
    )


# Define the available modes and an initial selection
modes = ['localhost', 'infrastructure']
selection = modes[0]

if selection == 'localhost':
    localhost()
elif selection == 'infrastructure':
    infrastructure()
