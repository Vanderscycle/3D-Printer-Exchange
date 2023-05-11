##
# 3dPrinterExchange
#
# @file
# @version 0.1

# BUG: swaggo/swag
# https://github.com/swaggo/swag/issues/1568
swag:
	(cd backend && swag init --parseDependency --parseInternal)
	grep -v ".*Delim" ./backend/docs/docs.go > tmpfile && mv tmpfile ./backend/docs/docs.go

start:
	make swag
	go fmt .
	(cd backend && go run main.go)
# end