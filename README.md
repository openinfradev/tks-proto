# TKS-proto
TKS-proto contains `.pb.go` files and protocol buffers for TKS services.  

## Getting Started
### Add new .proto
`.proto` files are located under `protos/` directory.
Please make and write new `.proto` file following the [Style Guide](https://developers.google.com/protocol-buffers/docs/style).

### Generate .pb.go files
You can easily generate `.pb.go` files using `make build` command.  
If there are no errors, `.pb.go` files would be generated in `pbgo/` directory.
```console
$ make build
go get -u github.com/golang/protobuf/protoc-gen-go
protoc --proto_path=protos --go_out=tks_pb --go_opt=paths=source_relative protos/*.proto

$ ls tks_pb
common.pb.go   contract.pb.go
```

### Generate python binding for grpc client
You can generate python binding using following command.  
```console
$ pip3 install grpcio grpcio-tools
$ python3 -m grpc_tools.protoc -I./protos --python_out=./tks_pb_python --grpc_python_out=./tks_pb_python protos/*.proto

$ ls tks_pb_python
cluster_lcm_pb2.py  cluster_lcm_pb2_grpc.py  ...  contract_pb2.py  contract_pb2_grpc.py  info_pb2.py  info_pb2_grpc.py
```
