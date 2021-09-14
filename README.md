# TKS-proto
TKS-proto contains protocol buffer definition and gRPC output files for TKS services.

## Getting Started
### Add new .proto
`.proto` files are located under `protos/` directory.
Please make and write new `.proto` file following the [Style Guide](https://developers.google.com/protocol-buffers/docs/style).

### Generate gRPC output files
You can easily generate output files for golang and python by running `make build` command.
```console
$ make build
protoc --proto_path=protos --go_out=tks_pb --go_opt=paths=source_relative \
--go-grpc_out=tks_pb --go-grpc_opt=paths=source_relative protos/*.proto
python3 -m grpc_tools.protoc -I./protos --python_out=./tks_pb_python --grpc_python_out=./tks_pb_python protos/*.proto
mockgen -source=tks_pb/cluster_lcm_grpc.pb.go > tks_pb/mock/mock_cluster_lcm_grpc.pb.go
mockgen -source=tks_pb/contract_grpc.pb.go > tks_pb/mock/mock_contract_grpc.pb.go
mockgen -source=tks_pb/info_grpc.pb.go > tks_pb/mock/mock_info_grpc.pb.go

$ ls tks_pb
cluster_lcm.pb.go       common.pb.go    contract_grpc.pb.go  info_grpc.pb.go
cluster_lcm_grpc.pb.go  contract.pb.go  info.pb.go           mock

$ ls tks_pb_python
cluster_lcm_pb2.py       common_pb2.py       contract_pb2.py       info_pb2.py
cluster_lcm_pb2_grpc.py  common_pb2_grpc.py  contract_pb2_grpc.py  info_pb2_grpc.py
```
