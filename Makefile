run:
	go run cmd/server/main.go
build:
	go build cmd/server/main.go
clean:
	go clean
test:
	go test clean_arch/interface/postgres

db_create:
	migrate create -dir db/migrations -ext sql ${ARGS}
db_up:
	migrate -path db/migrations -database postgres://postgres:postgres@localhost:5432/clean_arch_dev?sslmode=disable up 1
db_down:
	migrate -path db/migrations -database postgres://postgres:postgres@localhost:5432/clean_arch_dev?sslmode=disable down 1
db_force:
	migrate -path db/migrations -database postgres://postgres:postgres@localhost:5432/clean_arch_dev?sslmode=disable force ${ARGS}
