#!/bin/sh
source ./env.sh
java -jar  ${ALIENSBOOT_HOME}/bin/datatool.jar -d go -i ${PROJECT_PATH}/data -o ${SRC_PATH}/constant/tableconstant.go -t ${PROJECT_PATH}/templates/data/go_constant.template
java -jar  ${ALIENSBOOT_HOME}/bin/datatool.jar -d go -i ${PROJECT_PATH}/data -o ${SRC_PATH}/data/tabledata.go -t ${PROJECT_PATH}/templates/data/go_model.template