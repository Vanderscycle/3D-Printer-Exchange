##
# Backend
#
# @file
# @version 0.1

fmt-docs:
	swag init --parseDependency --parseInternal
	swag fmt
	go fmt github.com/...

start-backend:
	(cd ./database && docker compose up -d)
	make fmt-docs
	go run main.go


# end
