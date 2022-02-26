# gRPC Tutorial

simple tutorial to run grpc server and grpc client
both server and client share same .proto file. If you change .proto one side only, you may not be able to run server or client.
.proto is not language specific. you may run python grpc client or server using same proto file.
I would appreciate adding grpc server in other languages(python, nodejs etc.)

## COMMON

if you change your proto file
you should run

```Shell
protoc -I . --go_out=plugins=grpc:. protos/chat.proto
```

## SERVER

run following command in `grpc-server` directory

```Shell
go run server.go
```

## Client

run following command in `grpc-client` directory

```Shell
go run client.go
```

## Reference

https://tutorialedge.net/golang/go-grpc-beginners-tutorial/