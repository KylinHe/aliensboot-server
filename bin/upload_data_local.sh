#!/bin/sh
source ./env.sh

REGISTRY=127.0.0.1:2379    #配置服务器地址
DATA_PATH=${PROJECT_PATH}/data   #工程跟路径
ENV=aliensboot-local

#导出json格式数据
java -jar ${ALIENSBOOT_HOME}/bin/datatool.jar -d json -i ${DATA_PATH} -o ${DATA_PATH}/table_out_json

#上传数据
aliensboot data upload -e ${REGISTRY} -r /root/${ENV}/config -l  ${DATA_PATH}/table_out_json #-m ${ENV}-md5.sum
