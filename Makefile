generate:
	@templ generate ./...

build: generate
	@NODE_ENV=production npx tailwindcss build -o ./misc/output.css
	@go build -gcflags='all=-N -l' -o ./bin/app

run: build
	@./bin/app 

migrate:
	@migrate -database "postgres://postgres:postgres@localhost:5433/pronty?sslmode=disable" -path ./db/migrations up

create-migration:
	@migrate create -ext sql -dir ./db/migrations -seq $(migration_name)

migrate-rollback:
	@migrate -database "postgres://postgres:postgres@localhost:5433/pronty?sslmode=disable" -path ./db/migrations down

tailwind:
	@cd misc/; npx tailwindcss-cli@latest build -i input.css -o ../static/output.css
