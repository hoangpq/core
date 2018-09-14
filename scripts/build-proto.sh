#!/bin/bash -x

DIR=grpcclient
PWD=$(pwd)
# API_DOCS="--doc_out=$PWD/docs/api/ --doc_opt=$PWD/docs/api.template"
# DATA_DOCS="--doc_out=$PWD/docs/api/ --doc_opt=$PWD/docs/data.template"
GRPC_PLUGIN="--go_out=plugins=grpc:./"

protoc --go_out=plugins=grpc:$DIR --proto_path $DIR/ $DIR/service.proto
protoc --go_out=plugins=grpc:$DIR --proto_path $DIR/ $DIR/core_api.proto
protoc --go_out=plugins=grpc:$DIR --proto_path $DIR/ $DIR/service_api.proto
