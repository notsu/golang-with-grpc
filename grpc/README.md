# gRPC

## Client

```
cd client

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/greet.proto
```

## Server

```
cd server

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/greet.proto
```

## Table of Contents

1. Unary
2. Streaming on the server
3. Streaming on the client
4. Bi-direction streaming
