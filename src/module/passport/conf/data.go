//generate by aliensboot
package conf

import (
	"github.com/KylinHe/aliensboot-server/data"
	"encoding/json"
	"github.com/KylinHe/aliensboot-core/cluster/center"
	"github.com/KylinHe/aliensboot-core/common/util"
	"github.com/KylinHe/aliensboot-core/log"
)

func Init() {

	center.ClusterCenter.SubscribeConfig("maintain", updateMaintainBase, center.OptionEmpty)

	// [ip白名单表]
	center.ClusterCenter.SubscribeConfig("whitelist_ip", UpdateWhitelistIpData, center.OptionEmpty)

	// [账号白名单表]
	center.ClusterCenter.SubscribeConfig("whitelist_account", UpdateWhitelistAccountData, center.OptionEmpty)

	// [uid白名单表]
	center.ClusterCenter.SubscribeConfig("whitelist_account", UpdateWhitelistUidData, center.OptionEmpty)

}

func Close() {
}

var (
	// [维护信息]
	Maintain = &data.MaintainBase{}

	// [ip白名单表]
	WhitelistIpData = make(map[string]bool)

	// [账号白名单表]
	WhitelistAccountData = make(map[string]bool)

	// [uid白名单表]
	WhitelistUidData = make(map[int64]bool)
)

func UpdateWhitelistIpData(content []byte) {
	var dataArray []*data.WhitelistIp
	err := json.Unmarshal(content, &dataArray)
	if err != nil {
		log.Errorf("update data %v, err %v", "whitelist_ip", err)
	}
	results := make(map[string]bool)
	for _, data := range dataArray {
		results[data.Ip] = data.Enable
	}
	WhitelistIpData = results
}

func UpdateWhitelistAccountData(content []byte) {
	var dataArray []*data.WhitelistAccount
	err := json.Unmarshal(content, &dataArray)
	if err != nil {
		log.Errorf("update data %v, err %v", "whitelist_account", err)
	}
	results := make(map[string]bool)
	for _, data := range dataArray {
		results[data.Account] = data.Enable
	}
	WhitelistAccountData = results
}

func UpdateWhitelistUidData(content []byte) {
	var dataArray []*data.WhitelistUid
	err := json.Unmarshal(content, &dataArray)
	if err != nil {
		log.Errorf("update data %v, err %v", "whitelist_uid", err)
	}
	results := make(map[int64]bool)
	for _, data := range dataArray {
		results[data.Uid] = data.Enable
	}
	WhitelistUidData = results
}

func updateMaintainBase(content []byte) {
	maintain := &data.MaintainBase{}
	_ = json.Unmarshal(content, maintain)
	//startTimestamp, _ := time.ParseInLocation("2006-01-02 15:04:05", maintain.StartTime, time.Local)
	maintain.StartTime1 = util.GetTime(maintain.StartTime)
	//endTimestamp, _ := time.ParseInLocation("2006-01-02 15:04:05", maintain.EndTime, time.Local)
	maintain.EndTime1 = util.GetTime(maintain.EndTime)
	Maintain = maintain
}
