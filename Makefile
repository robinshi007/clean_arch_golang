run:
	go run cmd/server/main.go
build:
	go build cmd/server/main.go
clean:
	go clean
test:
	go test clean_arch/interface/postgres
