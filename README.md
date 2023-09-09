# 3D Printer Exchange 

A proof of competence displaying my full stack and devops abilities.

## Localhost deployment
For an easy way to deploy the Fronted, backend, and other ancillaries this program relies on `Tilt` to start and manage the entire stack. 

Just time `tilt up` in the root of the folder to get started


## Localhost Infrastructure integration testing

There are a variety of tools available to simulate a kubernetes cluster. In this case we rely on `kind` because it allows us to simulate a multi-node cluster. When used with `ctlptl`

### Kind

To create the cluster run `ctlptl apply -f ./devops/localhost/kind.yaml`
For deletion `ctlptl delete cluster king-infrastructure` 


First create
This stack packages itself into its respective containers (e.g. backend/frontend) and are fed to k8s.

## Stack
- Backend: GoFiber
- DB: Postgres
- IaC: Kubernetes
- Frontend: Sveltekit
- CSS: TailwindCss
- Testing:
  * Frontend: Cypress
  * Golang: Golang Testing lib

## References:
- [Tilt](https://docs.tilt.dev/tiltfile_authoring)
- [ctlptl](https://github.com/tilt-dev/ctlptl)
- [king](https://kind.sigs.k8s.io/docs/user/quick-start/)
- [Gofiber](https://github.com/gofiber/fiber)
- [SvelteKit]()
