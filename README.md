```
protoc pb/hello_grpc.proto --go_out=service
protoc pb/hello_grpc.proto --go-grpc_out=require_unimplemented_servers=false:service
protoc pb/hello_grpc.proto --php_out=php/client --grpc_out=php/client --plugin=protoc-gen-grpc=/usr/local/bin/grpc_php_plugin
```

##
- https://github.com/protocolbuffers/protobuf/tree/main/php
- https://github.com/grpc/grpc/tree/master/src/php
- google.golang.org/grpc/examples/helloworld
