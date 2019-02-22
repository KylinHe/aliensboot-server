package conf

import (
	"github.com/KylinHe/aliensboot-server/module/room/config"
)

func Init() {
	roomData = make(map[string]*config.RoomConfig)
	roomData["0"] = &config.RoomConfig{AppID: "0", MaxSeat: 2, Anchor:false, Viewer:false}
	roomData["1"] = &config.RoomConfig{AppID: "1", MaxSeat: 4, Anchor:true, Viewer:true}
	roomData["2"] = &config.RoomConfig{AppID: "2", MaxSeat: 10, Anchor:true, Viewer:true}
	//center.ClusterCenter.SubscribeConfig("testdata", UpdateArmyData)
}

func Close() {

}

var (
	roomData map[string]*config.RoomConfig
)

func GetRoomConfig(appID string) *config.RoomConfig {
	return roomData[appID]
}

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
