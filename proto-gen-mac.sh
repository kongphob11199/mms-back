#!/bin/bash

# กำหนดไฟล์ proto และ output directory
PROTO_FILE=./mms/pkg/proto/*.proto
OUT_DIR_SERVER=./mms/internal

# Compile proto files
protoc $PROTO_FILE --go_out=$OUT_DIR_SERVER --go-grpc_out=$OUT_DIR_SERVER

echo "Compilation complete!"

# รอผู้ใช้กด Enter ก่อนจบ
read -p "Press [Enter] to continue..."
