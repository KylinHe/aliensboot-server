package conf

import (
	"encoding/json"
	"github.com/KylinHe/aliensboot-core/cluster/center"
	"github.com/KylinHe/aliensboot-server/data"
)

func Init() {
	center.ClusterCenter.SubscribeConfig("game", UpdateGameData)

	GameData = make(map[string]data.Game)
	GameData["0"] = data.Game{MaxSeat: 2}
}

func Close() {

}

var (
	GameData map[string]data.Game
)

func UpdateGameData(content []byte) {
	var dataArray []data.Game
	json.Unmarshal(content, &dataArray)
	results := make(map[string]data.Game)
	for _, data := range dataArray {
		results[data.AppId] = data
	}
	GameData = results
}
