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
protoc --proto_path=protos --go_out=pbgo --go_opt=paths=source_relative protos/*.proto

$ ls pbgo
common.pb.go   contract.pb.go
```
