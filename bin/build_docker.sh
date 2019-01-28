#!/bin/sh
source ./env.sh
cd $PROJECT_PATH

#镜像地址
IMAGE_NAME=registry.cn-shenzhen.aliyuncs.com/aliensidea/aliensboot-server
#镜像版本号
Version=latest

docker build -t $IMAGE_NAME .

docker push $IMAGE_NAME:$Version