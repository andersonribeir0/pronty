generate:
	@templ generate ./...

build: generate
	@go build -o ./bin/app

run: build
	@./bin/app 

migrate:
	@migrate -database "postgres://postgres:postgres@localhost:5433/pronty?sslmode=disable" -path ./db/migrations up

create-migration:
	@migrate create -ext sql -dir ./db/migrations -seq $(migration_name)

migrate-rollback:
	@migrate -database "postgres://postgres:postgres@localhost:5433/pronty?sslmode=disable" -path ./db/migrations down

