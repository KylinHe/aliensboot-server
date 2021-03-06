${C_SHARP_CLIENT_PATH}
source ./env.sh

filelist=`ls ${SRC_PATH}/protocol/*.proto`
for file in $filelist
do
    echo $file
    protoc --csharp_out=${C_SHARP_CLIENT_PATH}/ --proto_path ${SRC_PATH}/protocol/ $(basename $file .proto).proto
done


echo done!