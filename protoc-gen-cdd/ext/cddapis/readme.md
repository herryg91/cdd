```
#protobuf protoc-gen-go V2
#in protoc-gen-cdd/ext/cddapis folder
protoc \
    --proto_path=$HOME/Workspace/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/cddapis/ \
    --go_out $HOME/Workspace/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/cddapis --go_opt paths=source_relative \
    --go-grpc_out $HOME/Workspace/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/cddapis --go-grpc_opt paths=source_relative \
    cdd/api/cddext.proto  
```
