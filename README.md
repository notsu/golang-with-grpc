# Golang with gRPC

## Prerequisite

- Install Golang on your machine
- Install dependencies
```bash
go get github.com/labstack/echo/v4
go get google.golang.org/grpc
go get google.golang.org/protobuf
go get golang.org/x/oauth2
```
- Protocol Buffer compiler https://grpc.io/docs/protoc-installation/
- Go plugins
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## Agenda

1. `rest` - Initiate a REST API with Golang's frameworks to simulate the web server in general
2. `grpc` - Simulate the communication between backend-to-backend using gRPC
3. `grpc-with-error-handling` - Setup custom error to handling
4. `grpc-with-authentication` - Setup authentication

## Future examples

Please click `watch` on the right top of this page to get updates

- `grpc-with-rest-and-swagger` - Generate REST API and Swagger from Protocol Buffer
- `grpc-with-kubernetes` - How to deal with load balancing
