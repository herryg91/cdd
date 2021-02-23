# protoc-gen-cdd
This is still in Alpha Version.

### What is protoc-gen-cdd
This is implementation of contract driven development. protoc-gen-cdd act as generator of the contracts. Right now it can generate cdd-grst frameworks. https://github.com/herryg91/cdd/tree/main/grst

### Requirement
```
- protoc-gen-go v1.25.0 (https://github.com/protocolbuffers/protobuf-go/tree/master/cmd/protoc-gen-go)
- protoc v3.13.0 (https://github.com/protocolbuffers/protobuf/tree/v3.13.0)
- protoc-gen-grpc-gateway v2.0.1 (https://github.com/grpc-ecosystem/grpc-gateway)
- protoc-gen-go-grpc 1.33.1 (https://github.com/grpc/grpc-go/tree/master/cmd/protoc-gen-go-grpc)

go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### How To Use
```
cd examples/province-api

protoc \
--proto_path=$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/cddapis/ \
--proto_path=$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/googleapis/ \
--proto_path=$GOPATH/src/ \
--proto_path=grpc/proto/ \
--proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--go_out ./grpc/pb/province \
--go-grpc_out ./grpc/pb/province \
--grpc-gateway_out ./grpc/pb/province --grpc-gateway_opt logtostderr=true --grpc-gateway_opt generate_unbound_methods=true \
--cdd_out=type=grst:./grpc/pb/province \
province.proto

protoc \
--proto_path=$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/cddapis/ \
--proto_path=$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/googleapis/ \
--proto_path=$GOPATH/src/ \
--proto_path=grpc/proto/ \
--proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--cdd_out=type=grst-validation,pbpath=$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/examples/province-api/grpc/pb/province:./grpc/pb/province \
province.proto

protoc \
--proto_path=$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/cddapis/ \
--proto_path=$GOPATH/src/github.com/herryg91/cdd/protoc-gen-cdd/ext/googleapis/ \
--proto_path=$GOPATH/src/ \
--proto_path=grpc/proto/ \
--proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--cdd_out=type=crud,pkgpath=github.com/herryg91/cdd/protoc-gen-cdd/examples/province-api:./ \
province.proto

notes:
1. 1st protoc call will generate *.pb.go, *_grpc.pb.go, *.pb.gw.go, *.pb.grst.go
2. 2nd protoc call will generate additional attribute in *.pb.go. Validation attribute and make json tag sam as jsonpb tag (json inside protobuf)
3. 3rd protoc call will generate repository and usecase CRUD for specific domain (if set in *.proto).
```
