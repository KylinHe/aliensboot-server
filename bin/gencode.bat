cd ../
set PROJECT_PATH=%cd%
set SRC_PATH=%PROJECT_PATH%/src

cd %SRC_PATH%/protocol/
set GOGOPATH=%GOPATH%/src
protoc --proto_path=%GOPATH%;%GOGOPATH%;./; --gogofast_out=plugins=grpc:. *.proto


cd %PROJECT_PATH%
set modules=game gate passport defaultmodule

for %%i in (%modules%) do (
    aliensboot module gen %%i
)

cd bin