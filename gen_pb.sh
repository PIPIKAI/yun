#!/bin/sh
protoc  --go_out=.  --go-grpc_out=. pb/yun.proto