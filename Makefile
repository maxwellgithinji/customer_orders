# Import .env
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)

psql: cmd-exists-psql
	psql "${DATABASE_URL}"

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

migrateinstall:
	curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
	echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
	apt-get update
	apt-get install -y migrate

migrationscreate:
	# migrate create -ext sql -dir ./databases/migrations -seq init_schema

migrateup: 
	migrate -path ./databases/migrations/ -database "${DATABASE_URL}" -verbose up

migratedown: 
	migrate -path ./databases/migrations/ -database "${DATABASE_URL}" -verbose down


.PHONY: server test  testcover testview swaginstall swag migrateinstall psql