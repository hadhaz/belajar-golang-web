build:
	go build -o bin/server cmd/server.go

run:
	bin/server

migrate:
	migrate -database "mysql://learner:secret@tcp(localhost:3306)/db_belajar_golang" -path db/migrations up

migrate-down:
	migrate -database "mysql://learner:secret@tcp(localhost:3306)/db_belajar_golang" -path db/migrations down

create-migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir db/migrations $$name
