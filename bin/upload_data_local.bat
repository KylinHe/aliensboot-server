cd ../
set PROJECT_PATH=%cd%

set REGISTRY=127.0.0.1:2379
set DATA_PATH=%PROJECT_PATH%/data
set ENV=aliensboot-local

java -jar %ALIENSBOOT_HOME%/bin/datatool.jar -d json -i %DATA_PATH% -o %DATA_PATH%/table_out_json

aliensboot data upload -e %REGISTRY% -r /root/%ENV%/config -l  %DATA_PATH%/table_out_json

cd bin