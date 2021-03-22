server:
	go run main.go

test:
	go test -v -cover ./...

testcover: 
	go test ./... -coverprofile=cover.out

testview:
	go tool cover -html=cover.out

swaginstall:
	go get -u github.com/swaggo/swag/cmd/swag

swag:
	swag init

.PHONY: server test  testcover testview swaginstall swag 