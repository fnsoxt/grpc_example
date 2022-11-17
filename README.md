```
protoc pb/hello_grpc.proto --go_out=service
protoc pb/hello_grpc.proto --grpc-gateway_out=service -I ./pb
protoc pb/hello_grpc.proto --go-grpc_out=require_unimplemented_servers=false:service
protoc pb/hello_grpc.proto --php_out=php/client --grpc_out=php/client --plugin=protoc-gen-grpc=/usr/local/bin/grpc_php_plugin

curl -X POST -k http://localhost:8090/v1/example/echo -d '{"name": " hello"}'
```
```shell
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
```
###启动
```shell
#启动项目
go run server/server.go
#启动预览
grpcui -plaintext  127.0.0.1:80880
```
##
- https://github.com/protocolbuffers/protobuf/tree/main/php
- https://github.com/grpc/grpc/tree/master/src/php
- google.golang.org/grpc/examples/helloworld
- https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/
