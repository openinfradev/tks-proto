.PHONY: compile
PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go
PROTOC := $(shell which protoc)
MOCKGEN := $(shell which mockgen)

# Make it always try to install the packages,
# since pip checks if the package is installed
PY_GRPC_TOOLS := must-install

# If protoc isn't on the path, set it to a target that's never up-to-date, so
# the install command always runs.
ifeq ($(PROTOC),)
    PROTOC = must-rebuild
endif

# Same as PROTOC
ifeq ($(MOCKGEN),)
    MOCKGEN = must-rebuild
endif

# Figure out which machine we're running on.
UNAME := $(shell uname)

$(PROTOC):
# Run the right installation command for the operating system.
ifeq ($(UNAME), Darwin)
	brew install protobuf
endif
ifeq ($(UNAME), Linux)
	sudo apt-get install protobuf-compiler
endif
# You can add instructions for other operating systems here, or use different
# branching logic as appropriate.

# If $GOPATH/bin/protoc-gen-go does not exist, we'll run this command to install
# it.
$(PROTOC_GEN_GO):
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

$(MOCKGEN):
	go get -u github.com/golang/mock/mockgen

$(PY_GRPC_TOOLS):
	pip3 install grpcio grpcio-tools

generate: protos | $(PROTOC_GEN_GO) $(PROTOC) $(PY_GRPC_TOOLS)
	protoc --proto_path=protos --go_out=tks_pb --go_opt=paths=source_relative \
	--go-grpc_out=tks_pb --go-grpc_opt=paths=source_relative protos/*.proto
	python3 -m grpc_tools.protoc -I./protos --python_out=./tks_pb_python --grpc_python_out=./tks_pb_python protos/*.proto

generate-mock: $(MOCKGEN)
	mockgen -source=tks_pb/cluster_lcm_grpc.pb.go > tks_pb/mock/mock_cluster_lcm_grpc.pb.go
	mockgen -source=tks_pb/contract_grpc.pb.go > tks_pb/mock/mock_contract_grpc.pb.go
	mockgen -source=tks_pb/info_grpc.pb.go > tks_pb/mock/mock_info_grpc.pb.go

# This is a "phony" target - an alias for the above command, so "make build"
# still works.
build: generate generate-mock
