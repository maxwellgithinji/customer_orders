server:
	go run main.go

swaginstall:
	go get -u github.com/swaggo/swag/cmd/swag

swag:
	swag init

.PHONY: server  swaginstall swag