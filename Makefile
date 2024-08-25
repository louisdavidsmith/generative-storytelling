# Define variables
PROTO_DIR = proto
GO_OUT_DIR = narrative/prompt/pkg
GENERATION_GO_OUT_DIR = narrative/generation/pkg
PROTOC = protoc
PROTOC_GEN_GO = $(PROTOC) --go_out=$(GO_OUT_DIR)
PROTOC_GEN_GO_GRPC = $(PROTOC) --go-grpc_out=$(GO_OUT_DIR)
PROTOC_GEN_GO_GENERATION = $(PROTOC) --go_out=$(GENERATION_GO_OUT_DIR)
PROTOC_GEN_GO_GRPC_GENERATION = $(PROTOC) --go-grpc_out=$(GENERATION_GO_OUT_DIR)

# Target for generating Go code from common.proto
common:
	$(PROTOC_GEN_GO) $(PROTO_DIR)/common.proto
	$(PROTOC_GEN_GO_GRPC) $(PROTO_DIR)/common.proto

# Target for generating Go code from prompt.proto
prompt:
	$(PROTOC_GEN_GO) $(PROTO_DIR)/prompt.proto
	$(PROTOC_GEN_GO_GRPC) $(PROTO_DIR)/prompt.proto

# Target for generating Go code from generation.proto
generation:
	$(PROTOC_GEN_GO_GENERATION) $(PROTO_DIR)/generation.proto
	$(PROTOC_GEN_GO_GRPC_GENERATION) $(PROTO_DIR)/generation.proto

# Default target
all: common prompt generation
