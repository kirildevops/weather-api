# Multi methods

## Method 1

Idea is to generate API server from existing swagger specification file using the [Go-Swagger](https://goswagger.io/go-swagger/generate/server/)

see the `go-swagger` [branch](https://github.com/kirildevops/weather-api/tree/go-swagger)

Status: **On Hold** due to SSL certificates being mandatory

## Method 2

Idea is to use Protobufs and gRPC and in the end add gRPC Gateway to conform with traditional HTTP/JSON standard REST API

see the `grpc_proto` [branch](https://github.com/kirildevops/weather-api/tree/grpc_proto)

Status: **Active** development
