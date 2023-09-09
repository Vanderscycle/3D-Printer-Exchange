##
# 3dPrinterExchange
#
# @file
# @version 0.1

# BUG: swaggo/swag
# https://github.com/swaggo/swag/issues/1568
swag:
	(cd backend && swag init --parseDependency --parseInternal && swag fmt)
	# grep -v ".*Delim" ./backend/docs/docs.go > tmpfile && mv tmpfile ./backend/docs/docs.go

start-backend:
	make swag
	(cd backend && go fmt .)
	(cd backend && go run main.go)

start-frontend:
	(cd frontend && pnpm run format)
	(cd frontend && pnpm run dev)

#WARN: will drop db
dev-db: 
	(bash ./backend/scripts/start-dev-db.sh)
