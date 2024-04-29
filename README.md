# blueprint-proto
This repository is separate module from the service.
## Requirement
- protobuf `brew install protobuf`
- grpc gateway 
  ```
  go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
  
### Link reference

https://grpc.io/docs/protoc-installation

https://github.com/grpc-ecosystem/grpc-gateway
## Getting started
Please read Makefile for command

- `make pbgen`command for generate pb.go and pb.gw.go and swagger.json