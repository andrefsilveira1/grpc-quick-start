# GRPC


- Create proto file first
- Define contracts
- Create rpc methods
- Generate proto interfaces with protoc 
- Create main and start server 
- Deal with requests 

## protoc:

- Run: `protoc --go_out=. --go-grpc_out=. proto/register.proto`

## How to run:

- Execute: `go run cmd/grpc/main.go`
- Then, start Evans: `docker run --rm --network host -it ghcr.io/ktr0731/evans --host localhost --port 50051 --reflection` 


## TODO: 

- [x] - Refactor README.md and include all step by step
- [x] - Refactor registers, remove haversine structure from registers.go
- [x] - Add tests and errors validation
- [x] - Create an article at Dev.to or Medium