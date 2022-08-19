```
protoc pb/hello_grpc.proto --go_out=service
protoc pb/hello_grpc.proto --grpc-gateway_out=service -I ./pb
protoc pb/hello_grpc.proto --go-grpc_out=require_unimplemented_servers=false:service
protoc pb/hello_grpc.proto --php_out=php/client --grpc_out=php/client --plugin=protoc-gen-grpc=/usr/local/bin/grpc_php_plugin

curl -X POST -k http://localhost:8090/v1/example/echo -d '{"name": " hello"}'
```

##
- https://github.com/protocolbuffers/protobuf/tree/main/php
- https://github.com/grpc/grpc/tree/master/src/php
- google.golang.org/grpc/examples/helloworld
- https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/
