#!/bin/bash

PROJECT_ROOT="$(pwd)"

SERVICES=("payment-service" "booking-service" "property-service" "user-service")

TARGET_DIR="$PROJECT_ROOT/shared/proto"

for service in "${SERVICES[@]}"; do
    PROTO_DIR="$PROJECT_ROOT/$service/proto"

    if [ -d "$PROTO_DIR" ]; then
        for proto_file in "$PROTO_DIR"/*.proto; do
            protoc --proto_path="$PROTO_DIR" --go_out="$TARGET_DIR" --go_opt=paths=source_relative --go-grpc_out="$TARGET_DIR" --go-grpc_opt=paths=source_relative "$proto_file"
        done
    else
        echo "Proto directory for $service not found."
    fi
done

echo "Proto compilation complete."