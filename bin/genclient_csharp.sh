${C_SHARP_CLIENT_PATH}
source ./env.sh

filelist=`ls ${SRC_PATH}/protocol/*.proto`
for file in $filelist
do
    echo $file
    protogen -i:$file -o:${C_SHARP_CLIENT_PATH}/$(basename $file .proto).cs
done


echo done!