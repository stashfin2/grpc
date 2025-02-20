#!/bin/bash

# Define the base Go module path
BASE_PACKAGE="github.com/stashfin2/grpc/stashfin.com"
MAIN_PACKAGE="stashfin.com"

# Define the proto directory
PROTO_DIR="protos"

# Generate Go options dynamically for all .proto files
GO_OPTS=""
for proto_file in $(find "$PROTO_DIR" -name "*.proto"); do
    proto_file_rel="${proto_file#"$PROTO_DIR/"}"  # Remove leading "protos/"

    # Extract directory path (excluding filename)
    proto_dir=$(dirname "$proto_file_rel")

    # Construct the Go package path dynamically
    if [[ "$proto_dir" == "." ]]; then
        go_package_path="$BASE_PACKAGE"
    else
        go_package_path="$BASE_PACKAGE/$proto_dir"
    fi

    # Append the go option for this proto file
    GO_OPTS+=" --go_opt=M$proto_file_rel=$go_package_path"
    GO_OPTS+=" --go-grpc_opt=M$proto_file_rel=$go_package_path"
done

# Run the protoc command
protoc \
  --proto_path="$PROTO_DIR" \
  --go_out=. --go-grpc_out=. \
  $GO_OPTS \
  $(find "$PROTO_DIR" -name "*.proto")

# Move generated files to the correct location
mkdir -p "$MAIN_PACKAGE"
mv grpc/github.com/stashfin2/grpc/stashfin.com/* "$MAIN_PACKAGE/" 2>/dev/null || true
rm -rf grpc/github.com

echo "✅ Protobuf files successfully generated!"
