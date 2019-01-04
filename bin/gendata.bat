call env.bat

java -jar  %ALIENSBOT_HOME%/bin/datatool.jar -d go -i %PROJECT_PATH%/data -o %PROJECT_PATH%/server/constant/tableconstant.go -t %PROJECT_PATH%/templates/data/go_constant.template
java -jar  %ALIENSBOT_HOME%/bin/datatool.jar -d go -i %PROJECT_PATH%/data -o %PROJECT_PATH%/server/data/tabledata.go -t %PROJECT_PATH%/templates/data/go_model.template