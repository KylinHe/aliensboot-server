cd ../
set PROJECT_PATH=%cd%
set SRC_PATH=%PROJECT_PATH%/src
set JS_CLIENT_PATH=%ALIENSBOOT_HOME%/aliensboot-client-cocos/assets/Script/aliensboot/protocol/protocol.js

python %ALIENSBOOT_HOME%/bin/protoToJs.py %SRC_PATH%/protocol/ %JS_CLIENT_PATH%

cd bin