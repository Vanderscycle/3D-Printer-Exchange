# Welcome to Tilt!
#   To get you started as quickly as possible, we have created a
#   starter Tiltfile for you.
#
#   Uncomment, modify, and delete any commands as needed for your
#   project's configuration.
modes = ['localhost', 'infrastructure']
selection = modes[1]



# Variables
# sync_src_frontend = sync('./frontend', '/src')
# sync_src_backend = sync('./backend', '/src')
localhost_dbs = "./backend/database/docker-compose.yml"

def localhost():
    print("""
    -----------------------------------------------------------------
    ✨ Localhost Environment
    -----------------------------------------------------------------
    """.strip())
    docker_compose(localhost_dbs)
localhost()
