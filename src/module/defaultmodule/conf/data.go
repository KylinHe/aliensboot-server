//generate by aliensboot
package conf

import (
	"github.com/KylinHe/aliensboot-server/data"
	"github.com/KylinHe/aliensboot-server/protocol"
	"encoding/json"
	"github.com/KylinHe/aliensboot-core/cluster/center"
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-core/module/base"
	"strconv"
)

func Init(skeleton *base.Skeleton) {
    
    // [兵种表]
    center.ClusterCenter.SubscribeConfig("army", center.NewDataProxy("army", skeleton, UpdateArmyData).OnDataChange)
    
}

func Close() {
}

var (
    
    // [兵种表]
    ArmyData map[int32]*data.Army
    
)

    
func UpdateArmyData(content []byte, init bool) {
	var dataArray []*data.Army
	err := json.Unmarshal(content, &dataArray)
	if err != nil {
		log.Errorf("update data %v, err %v","army", err)
	}
	results := make(map[int32]*data.Army)
	for _, data := range dataArray {
		results[data.Tid] = data
	}
	ArmyData = results
}

func GetArmyData(id int32) *data.Army {
	if ArmyData == nil {
		exception.GameException(protocol.Code_ConfigException)
	    exception.GameException(&protocol.CodeMessage{
            Code: protocol.Code_ConfigException,
            Param: []string{"Army",strconv.Itoa(int(id))},
        })
	}
	return ArmyData[id]
}
    

