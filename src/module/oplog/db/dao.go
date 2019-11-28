/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/4/12
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package db

import (
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/KylinHe/aliensboot-core/database"
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-core/log"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var TableMapping = make(map[string]database.IData)

const (
	AnalysisTypeDau = "log_dau" //活跃用户
	AnalysisTypeDv  = "log_dv"  //访问次数
	AnalysisTypeDr  = "log_dr"  //注册用户
	AnalysisTypeDrt = "log_drt" //累计注册用户
	AnalysisTypeDav = "log_dav" //人均访问次数
	AnalysisTypeDat = "log_dat" //人均停留时长

	AnalysisTypeDrr1 = "log_drr1" //注册次留
	AnalysisTypeDrr3 = "log_drr3" //注册3留
	AnalysisTypeDrr7 = "log_drr7" //注册7留
	AnalysisTypeDar1 = "log_dar1" //活跃次留
	AnalysisTypeDar3 = "log_dar3" //活跃3留
	AnalysisTypeDar7 = "log_dar7" //活跃7留
)

func RegisterDayTable(key string, data database.IData) {
	TableMapping[key] = data
	EnsureTable(key, data)
}

const dayUtil = 24 * 60 * 60

func DumpDayData(dayStr string) {
	dayTime, _ := time.ParseInLocation("2006-01-02", dayStr, time.Local)
	startTime := dayTime.Unix()
	endTime := dayTime.Unix() + dayUtil

	log.Debugf("dump day log %v start....", dayStr)
	groupByRole := GetGroupSumData(&protocol.LogLogin{}, startTime, endTime, "$roleid")
	dau := int32(len(groupByRole))
	dv := int32(0)
	dav := int32(0)
	dayActiveRole := make(map[interface{}]bool, dau)
	for _, role := range groupByRole {
		dv += int32(role["count"].(int))
		dayActiveRole[role["_id"]] = true
	}

	if dau != 0 {
		dav = dv / dau
	}
	dr := GetCount(&protocol.LogRegister{}, startTime, endTime)
	drt := GetCount(&protocol.LogRegister{}, 0, endTime)
	dat := GetAvgData(&protocol.LogLogout{}, startTime, endTime, "$activetime")

	_ = _db.ForceUpdateOne(&protocol.LogDayActiveUser{Id: startTime, Value: dau})
	_ = _db.ForceUpdateOne(&protocol.LogDayVisit{Id: startTime, Value: dv})
	_ = _db.ForceUpdateOne(&protocol.LogDayRegister{Id: startTime, Value: dr})
	_ = _db.ForceUpdateOne(&protocol.LogDayRegisterTotal{Id: startTime, Value: drt})
	_ = _db.ForceUpdateOne(&protocol.LogDayAvgVisit{Id: startTime, Value: dav})
	_ = _db.ForceUpdateOne(&protocol.LogDayAvgTime{Id: startTime, Value: dat})

	// 导出活跃数据
	_ = _db.ForceUpdateOne(&protocol.LogDayRegisterRetention1{Id: startTime, Value: GetRetention(dayActiveRole, &protocol.LogRegister{}, startTime, endTime, 1)})
	_ = _db.ForceUpdateOne(&protocol.LogDayRegisterRetention3{Id: startTime, Value: GetRetention(dayActiveRole, &protocol.LogRegister{}, startTime, endTime, 3)})
	_ = _db.ForceUpdateOne(&protocol.LogDayRegisterRetention7{Id: startTime, Value: GetRetention(dayActiveRole, &protocol.LogRegister{}, startTime, endTime, 7)})
	_ = _db.ForceUpdateOne(&protocol.LogDayActiveRetention1{Id: startTime, Value: GetRetention(dayActiveRole, &protocol.LogLogin{}, startTime, endTime, 1)})
	_ = _db.ForceUpdateOne(&protocol.LogDayActiveRetention3{Id: startTime, Value: GetRetention(dayActiveRole, &protocol.LogLogin{}, startTime, endTime, 3)})
	_ = _db.ForceUpdateOne(&protocol.LogDayActiveRetention7{Id: startTime, Value: GetRetention(dayActiveRole, &protocol.LogLogin{}, startTime, endTime, 7)})

	log.Debugf("dump day log %v end....", dayStr)
}

func InsertData(data database.IData) error {
	return _db.Insert(data)
}

//GetRetention
func GetRetention(dayActiveRole map[interface{}]bool, data database.IData, currStartTime int64, currEndTime int64, day int64) float32 {
	if len(dayActiveRole) == 0 {
		return 0
	}
	compareStartTime := currStartTime + day*dayUtil
	compareEndTime := currEndTime + day*dayUtil
	compareData := GetGroupSumData(data, compareStartTime, compareEndTime, "$roleid")
	if compareData == nil || len(compareData) == 0 {
		return 0
	}
	//compareData1 := buildMapping(compareData)
	existCount := 0

	for _, d := range compareData {
		if dayActiveRole[d["_id"]] {
			existCount++
		}
	}
	return float32(existCount) / float32(len(compareData))
}

func GetDayData(data database.IData, startTime int64, endTime int64) []map[string]interface{} {
	var result []map[string]interface{}
	_ = _db.QueryAllConditions(data, bson.M{"_id": bson.M{"$gte": startTime, "$lte": endTime}}, &result)
	return result
}

func GetData(data database.IData, startTime int64, endTime int64, timeSize int64) []map[string]interface{} {
	var result []map[string]interface{}
	group := bson.M{
		"_id": bson.M{
			"$subtract": []bson.M{
				bson.M{"$subtract": []interface{}{"$time", 0}},
				bson.M{"$mod": []interface{}{
					bson.M{"$subtract": []interface{}{"$time", 0}},
					timeSize},
				},
			}},
		"count": bson.M{"$sum": 1},
	}
	aggregate := []bson.M{
		{"$match": bson.M{"time": bson.M{"$gte": startTime, "$lt": endTime}}},
		{"$group": group},
	}
	err := _db.PipeAllConditions(data, aggregate, &result)
	if err != nil {
		exception.GameException(protocol.Code_DBExcetpion)
	}
	return result
}

func GetGroupSumData(data database.IData, startTime int64, endTime int64, groupField interface{}) []map[string]interface{} {
	var result []map[string]interface{}
	group := bson.M{
		"_id":   groupField,
		"count": bson.M{"$sum": 1},
	}
	aggregate := []bson.M{
		{"$match": bson.M{"time": bson.M{"$gte": startTime, "$lt": endTime}}},
		{"$group": group},
	}
	err := _db.PipeAllConditions(data, aggregate, &result)
	if err != nil {
		exception.GameException(protocol.Code_DBExcetpion)
	}
	return result
}

func GetAvgData(data database.IData, startTime int64, endTime int64, avgField interface{}) int32 {
	var result []map[string]interface{}
	group := bson.M{
		"_id":   nil,
		"count": bson.M{"$avg": avgField},
	}
	aggregate := []bson.M{
		{"$match": bson.M{"time": bson.M{"$gte": startTime, "$lt": endTime}}},
		{"$group": group},
	}
	err := _db.PipeAllConditions(data, aggregate, &result)
	if err != nil {
		exception.GameException(protocol.Code_DBExcetpion)
	}
	if result == nil || len(result) == 0 {
		return 0
	}
	return int32(result[0]["count"].(float64))
}

func GetCount(data database.IData, startTime int64, endTime int64) int32 {
	result, _ := _db.QueryConditionsCount(data, bson.M{"time": bson.M{"$gte": startTime, "$lt": endTime}})
	return int32(result)
}

func GetLiveData(startTime int64, endTime int64, timeSize int64) []map[string]interface{} {
	return GetData(&protocol.LogLogin{}, startTime, endTime, timeSize)
}

func GetDauData(startTime int64, endTime int64) interface{} {
	return GetDayData(&protocol.LogDayActiveUser{}, startTime, endTime)
}

func GetDvData(startTime int64, endTime int64) interface{} {
	return GetDayData(&protocol.LogDayVisit{}, startTime, endTime)
}

func GetDrData(startTime int64, endTime int64) interface{} {
	return GetDayData(&protocol.LogDayRegister{}, startTime, endTime)
}

func GetDrtData(startTime int64, endTime int64) interface{} {
	return GetDayData(&protocol.LogDayRegisterTotal{}, startTime, endTime)
}

func GetDavData(startTime int64, endTime int64) interface{} {
	return GetDayData(&protocol.LogDayAvgVisit{}, startTime, endTime)
}

func GetDatData(startTime int64, endTime int64) interface{} {
	return GetDayData(&protocol.LogDayAvgTime{}, startTime, endTime)
}

// 查询钻石消费日志
func QueryDiamondLog(begin, end int64, skip int, limit int) []*protocol.LogDiamond {
	var result []*protocol.LogDiamond
	conditions := bson.M{"$and": []bson.M{{"time": bson.M{"$gt": begin}}, {"time": bson.M{"$lt": end}}}}
	err := _db.QueryAllConditionsSkipLimit(&protocol.LogDiamond{}, conditions, &result, skip, limit)
	if err != nil {
		exception.GameException(protocol.Code_DBExcetpion)
	}
	return result
}

// 查询登陆日志
func QueryLoginLog(begin, end int64, skip int, limit int) []*protocol.LogLogin {
	var result []*protocol.LogLogin
	conditions := bson.M{"$and": []bson.M{{"time": bson.M{"$gt": begin}}, {"time": bson.M{"$lt": end}}}}
	err := _db.QueryAllConditionsSkipLimit(&protocol.LogLogin{}, conditions, &result, skip, limit)
	if err != nil {
		exception.GameException(protocol.Code_DBExcetpion)
	}
	return result
}
