/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/11/14
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package match

import (
	"github.com/KylinHe/aliensboot-core/common/util"
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-server/constant"
	"github.com/KylinHe/aliensboot-server/dispatch/rpc"
	"github.com/KylinHe/aliensboot-server/module/hall/conf"
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/eapache/queue"
)

var Manager = &manager{queues: make(map[string]*queue.Queue)}

func init() {
	//TODO 初始化房间
	Manager.queues[constant.TestAppID] = queue.New()
}

type manager struct {
	queues map[string]*queue.Queue //appid - queue
}

func (this *manager) Add(appID string, authID int64) {
	queue := this.EnsureQueue(appID)
	queue.Add(authID)

}

func (this *manager) TryMatch() {
	for appID, queue := range this.queues {
		configData := conf.GameData[appID]
		count := int(configData.MaxSeat)

		if count > 0 && queue.Length() >= count {
			results := make([]*protocol.Player, count)
			for i := 0; i < count; i++ {
				playerID := queue.Remove().(int64)
				results[i] = &protocol.Player{Playerid: playerID, Nickname: "蛇皮" + util.Int64ToString(playerID)}
			}

			rpc.Room.RoomCreate("", &protocol.RoomCreate{
				AppID:   appID,
				//Players: results,
			})
		}
	}

}

func (this *manager) EnsureQueue(appID string) *queue.Queue {
	queue := this.queues[appID]
	if queue == nil {
		exception.GameException(protocol.Code_appIDNotFound)
	}
	return queue
}
