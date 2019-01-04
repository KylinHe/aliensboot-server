package conf

import (
	"github.com/KylinHe/aliensboot-server/data"
)

func Init() {
	//center.ClusterCenter.SubscribeConfig("testdata", UpdateArmyData)
}

func Close() {

}

var (
	armyData map[int32]*data.Army
)

//func UpdateTestData(content []byte) {
//	var dataArray []*data.Army
//	json.Unmarshal(content, &dataArray)
//	results := make(map[int32]*data.Army)
//	for _, data := range dataArray {
//		results[data.Tid] = data
//	}
//	armyData = results
//}
//
//func GetArmyData(id int32) *data.Army {
//	if armyData == nil {
//		exception.GameException(protocol.Code_ConfigException)
//	}
//	return armyData[id]
//}
