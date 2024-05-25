# Define the Go binary and source paths
VERSION=v1
PROTO_SRC = proto
PROTO_DEST = services
GO=go
SRC=./cmd/server/main.go
OUTPUT_DIR=./bin
OUTPUT=$(OUTPUT_DIR)/server

# Set environment variables for the gRPC server
ENV=development
PORT=50051

# Default target: build and run the server
.PHONY: all
all: build run

# Install Go dependencies
.PHONY: deps
deps:
	$(GO) mod tidy
	$(GO) mod vendor

# Build the Go gRPC server
.PHONY: build
build: deps
	@echo "Building Go gRPC server..."
	@mkdir -p $(OUTPUT_DIR)
	$(GO) build -o $(OUTPUT) $(SRC)

# Run the Go gRPC server
.PHONY: run
run: build
	@echo "Running Go gRPC server..."
	@ENV=$(ENV) PORT=$(PORT) $(OUTPUT)

# Clean up binaries and vendor directory
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(OUTPUT_DIR) ./vendor

# Generate gRPC code from proto files
.PHONY: proto
proto:
	@echo "Generating gRPC code..."
	@mkdir -p $(PROTO_DEST)/$(VERSION)
	@for file in $(PROTO_SRC)/$(VERSION)/*.proto; do \
		service_name=$$(basename $$file .proto); \
		echo $$service_name; \
		mkdir -p $(PROTO_DEST)/$(VERSION)/$$service_name; \
		protoc \
			--go_out=$(PROTO_DEST)/$(VERSION)/$$service_name --go_opt=paths=source_relative \
			--go-grpc_out=$(PROTO_DEST)/$(VERSION)/$$service_name --go-grpc_opt=paths=source_relative \
			$$file; \
		mv $(PROTO_DEST)/$(VERSION)/$$service_name/proto/v1/*.go $(PROTO_DEST)/$(VERSION)/$$service_name/; \
		rm -rf $(PROTO_DEST)/$(VERSION)/$$service_name/proto; \
	done


# Check code formatting and run linters
.PHONY: lint
lint:
	@echo "Running linters..."
	@golangci-lint run

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test -v ./...

# Display help
.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  all       - Build and run the server"
	@echo "  deps      - Install Go dependencies"
	@echo "  build     - Build the Go gRPC server"
	@echo "  run       - Run the Go gRPC server"
	@echo "  clean     - Clean up binaries and vendor directory"
	@echo "  proto     - Generate gRPC code from proto files"
	@echo "  lint      - Check code formatting and run linters"
	@echo "  test      - Run tests"
	@echo "  help      - Display this help message"
