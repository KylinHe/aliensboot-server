package internal

import (
	"github.com/KylinHe/aliensboot-server/constant"
	"github.com/KylinHe/aliensboot-server/module/oplog/cache"
	"github.com/KylinHe/aliensboot-server/module/oplog/db"
	"github.com/KylinHe/aliensboot-core/database"
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-core/task"
	"time"
)

var dayTimeStr string = time.Now().Format("2006-01-02")

func init() {
	// 向当前模块注册客户端发送的消息处理函数 handleMessage
	skeleton.RegisterChanRPC(constant.LogCommand, handle)

	//每天凌晨4点导出前一天的数据
	cron, err := task.NewCronExpr("0 4 * * *")
	if err != nil {
		log.Error("init dump timer error : %v", err)
	}
	skeleton.CronFunc(cron, dayDump)
}

func dayDump() {
	dayTimeStr = time.Now().Add(-12 * time.Hour).Format("2006-01-02")
	if cache.SetNX("dump:dayLog:"+dayTimeStr, 1) {
		db.DumpDayData(dayTimeStr)
	}
}

func handle(args []interface{}) {
	dbLog := args[0]
	err := db.InsertData(dbLog.(database.IData))
	if err != nil {
		log.Debug("insert log error %v", err)
	}
}
