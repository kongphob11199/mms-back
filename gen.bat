@echo off
set PROTO_FILE=./mms/pkg/proto/*.proto

set OUT_DIR_SERVER=./mms/internal

protoc %PROTO_FILE% --go_out=%OUT_DIR_SERVER% --go-grpc_out=%OUT_DIR_SERVER%

echo Compilation complete!
pause