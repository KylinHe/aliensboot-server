call env.bat

:生成protobuf go协议代码
cd %SRC_PATH%/protocol/
set GOGOPATH=%GOPATH%/src
protoc --proto_path=%GOPATH%;%GOGOPATH%;./; --gogofast_out=plugins=grpc:. *.proto

:生成服务代码
cd %PROJECT_PATH%
set modules=game gate passport defaultmodule

for %%i in (%modules%) do (
    aliensboot module gen %%i
)

cd bin