#!/bin/sh
#basepath=$(cd `dirname $0`; pwd)
source ./env.sh

#生成protobuf go协议代码
cd ${SRC_PATH}/protocol/
GOGOPATH=${GOPATH}/src; protoc --proto_path=${GOPATH}:${GOGOPATH}:./ --gogofast_out=plugins=grpc:. *.proto

cd ${PROJECT_PATH}
modules=$(ls ${SRC_PATH}/module)
for module in $modules
do
    aliensboot module gen $module
done

#生成服务代码
#modules=(game gate passport hall room scene)
#for i in "${!modules[@]}"; do
#	aliensboot module gen ${modules[$i]}
#done