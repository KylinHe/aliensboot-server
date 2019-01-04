call env.bat

:生成protobuf go协议代码
cd %PROJECT_PATH%/server/protocol/
set GOGOPATH=%GOPATH%/src
protoc --proto_path=%GOPATH%;%GOGOPATH%;./; --gogofast_out=plugins=grpc:. *.proto

:生成服务代码
cd %PROJECT_PATH%
set modules=game gate passport hall room scene

for %%i in (%modules%) do (
    aliensbot.exe module gen %%i
)

cd bin