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


## Dev.to:

[Dev.to article](https://dev.to/andrefsilveira1/grpc-quick-start-coding-with-streams-and-bidirectional-streaming-4dkd)

