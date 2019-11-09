run:
	go run cmd/server/main.go
run_rpc:
	go run cmd/rpc_server/main.go
build:
	go build cmd/server/main.go
clean:
	go clean
test:
	go test ./... -v
protoc:
	protoc --proto_path=. --go_out=plugins=grpc:./ endpoint/rpc/v1/protocol/*.proto
grpcc:
	grpcc --proto endpoint/rpc/v1.0/protocol/*.proto --address 127.0.0.1:8081 -i

db_create:
	migrate create -dir db/migrations -ext sql ${ARGS}
db_up:
	migrate -path db/migrations -database postgres://postgres:postgres@localhost:5432/clean_arch_dev?sslmode=disable up 1
db_down:
	migrate -path db/migrations -database postgres://postgres:postgres@localhost:5432/clean_arch_dev?sslmode=disable down 1
db_force:
	migrate -path db/migrations -database postgres://postgres:postgres@localhost:5432/clean_arch_dev?sslmode=disable force ${ARGS}
