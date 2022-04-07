```
protoc  pb/hello_grpc.proto --go_out=service
protoc  pb/hello_grpc.proto --go-grpc_out=require_unimplemented_servers=false:service
```
