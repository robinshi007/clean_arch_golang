run:
	go run cmd/server/main.go
run_rpc:
	go run cmd/rpc_server/main.go
run_gen:
	go run cmd/gen/main.go
build:
	go build cmd/server/main.go
clean:
	go clean
test:
	go test ./... -v
test_db:
	go test clean_arch/adapter/postgres -v
test_ucase:
	go test clean_arch/usecase -v
test_handler:
	go test clean_arch/endpoint/api/handler -v

protoc:
	protoc --proto_path=. --go_out=plugins=grpc:./ endpoint/rpc/v1/protocol/*.proto
grpcc:
	grpcc --proto endpoint/rpc/v1.0/protocol/*.proto --address 127.0.0.1:8081 -i
gql:
	cd endpoint/api/graphql && go run github.com/99designs/gqlgen -v && cd -

# make ARGS="test" db_create
db_create:
	migrate create -dir db/migrations -ext sql ${ARGS}
db_up:
	migrate -path db/migrations -database postgres://postgres:postgres@localhost:5432/clean_arch_dev?sslmode=disable up 1
db_down:
	migrate -path db/migrations -database postgres://postgres:postgres@localhost:5432/clean_arch_dev?sslmode=disable down 1
db_force:
	migrate -path db/migrations -database postgres://postgres:postgres@localhost:5432/clean_arch_dev?sslmode=disable force ${ARGS}
