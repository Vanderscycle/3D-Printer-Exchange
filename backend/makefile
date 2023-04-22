##
# Backend
#
# @file
# @version 0.1

fmt:
	swag init --parseDependency --parseInternal
	swag fmt
	go fmt github.com/...

start:
	(cd ./database && docker compose up -d)
	make fmt
	go run main.go

# end
