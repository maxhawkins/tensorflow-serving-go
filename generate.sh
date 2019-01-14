#!/bin/bash

OPTIONS=""
for file in $(find proto/tensorflow | grep 'proto$'); do
    package="github.com/maxhawkins/tensorflow-serving-go/$(dirname $file | sed 's#^proto/tensorflow/##')"
    proto="$(echo $file | sed 's#^proto/tensorflow/##')"
    OPTIONS+="M$proto=$package,"
done
for file in $(find proto/serving | grep 'proto$'); do
    package="github.com/maxhawkins/tensorflow-serving-go/$(dirname $file | sed 's#^proto/serving/##')"
    proto="$(echo $file | sed 's#^proto/serving/##')"
    OPTIONS+="M$proto=$package,"
done
OPTIONS+="paths=source_relative,"
OPTIONS+="plugins=grpc"
OPTIONS+=":."

protoc -I proto/tensorflow -I proto/serving --go_out="$OPTIONS" proto/serving/tensorflow_serving/apis/*.proto
protoc -I proto/tensorflow -I proto/serving --go_out="$OPTIONS" proto/serving/tensorflow_serving/config/*.proto
protoc -I proto/tensorflow -I proto/serving --go_out="$OPTIONS" proto/serving/tensorflow_serving/sources/storage_path/*.proto
protoc -I proto/tensorflow -I proto/serving --go_out="$OPTIONS" proto/serving/tensorflow_serving/util/*.proto
protoc -I proto/tensorflow -I proto/serving --go_out="$OPTIONS" proto/tensorflow/tensorflow/core/example/*.proto
protoc -I proto/tensorflow -I proto/serving --go_out="$OPTIONS" proto/tensorflow/tensorflow/core/framework/*.proto
protoc -I proto/tensorflow -I proto/serving --go_out="$OPTIONS" proto/tensorflow/tensorflow/core/lib/core/*.proto
protoc -I proto/tensorflow -I proto/serving --go_out="$OPTIONS" proto/tensorflow/tensorflow/core/protobuf/*.proto
 