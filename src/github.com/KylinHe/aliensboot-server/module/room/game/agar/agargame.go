/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2019/02/15
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package agar

import (
	"encoding/json"
	util2 "github.com/KylinHe/aliensboot-core/common/util"
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-core/mmo"
	"github.com/KylinHe/aliensboot-core/mmo/core"
	"github.com/KylinHe/aliensboot-server/constant"
	"github.com/KylinHe/aliensboot-server/data"
	"github.com/KylinHe/aliensboot-server/module/room/game"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/collision"
	protocol2 "github.com/KylinHe/aliensboot-server/module/room/game/agar/protocol"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/util"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/vision"
	"github.com/KylinHe/aliensboot-server/module/scene/entity"
	"github.com/KylinHe/aliensboot-server/protocol"
	"time"
)


func init () {
	mmo.RegisterEntity(&GameEntity{})
	mmo.RegisterEntity(&entity.Player{})
}

type GameFactory struct {

}

//var battles map[int32]*Game
var BattleIDCounter int32 = 1

func (this *GameFactory) NewGame(handler game.Handler) game.Game {
	return NewGame(handler)
	//&Game{
	//	CommonGame: &game.CommonGame{Handler: handler},
	//}
}

type GameEntity struct {
	mmo.Entity

}

const (
	TypeGameEntity mmo.EntityType = "GameEntity"
)

func (space *GameEntity) DescribeEntityType(desc *core.EntityDesc) {

}

func NewGame(handler game.Handler) *Game {

	o := &Game{}
	//mmo.RegisterEntityHandler(&EntityHandler{timerManager:timerManager})

	//utils.BuildEntity(o.Entity, true)
	o.CommonGame = &game.CommonGame{Handler:handler}
	o.id = BattleIDCounter
	o.users = make(map[int32]*BattleUser)
	o.tickCount = 0
	o.gameOverTick = data.AtarGame.GameTime * 1000
	o.lastSysTick = util.TimeNowSysTick()
	o.mapBorder = &util.MapBorder{
		BottomLeft:util.Position{X:0,Y:0},
		TopRight:util.Position{X:float64(data.AtarGame.MapWidth),Y:float64(data.AtarGame.MapWidth)},
	}
	o.ballIDCounter = 1
	o.colMgr = collision.NewCollision()
	o.visionMgr = vision.NewVision()
	o.starMgr = NewStarMgr(o)
	o.dummyUser = NewBattleUser(nil, 0, o)
	//o.dummyUser.Battle = o
	//o.dummyUser.visionObj = &vision.VisionObject{
	//	Vision:o.visionMgr,
	//	Proxy:o.dummyUser,
	//}
	//todo 暂时把ai去掉
	//o.aiMgr = NewAiMgr(o, 20)
	//todo 暂时去掉刺球
	//o.thornMgr = NewThornMgr(o)
	o.updateCount = 0
	o.lastSyncBallUpdate = 0
	BattleIDCounter += 1
	o.tickInterval = 50
	//timer open

	o.Start()
	return o
}

func (game *Game)Start() {
	timerMgr := game.Handler.GetTimerMgr()
	game.timer = timerMgr.AddTimer(50*time.Millisecond, game.GameUpdate)
}

type Game struct {
	*game.CommonGame

	id 			int32
	battleId 	string
	users	 	map[int32]*BattleUser
	tickCount 	int32
	tickInterval int32
	gameOverTick int32
	lastSysTick  int32
	updateCount int32
	//nowSysTick	int32
	lastSyncBallUpdate int32

	sysTimerID core.EntityTimerID

	mapBorder    *util.MapBorder

	colMgr collision.ICollision
	visionMgr vision.IVision
	starMgr *StarMgr
	aiMgr *AiMgr
	thornMgr *ThornMgr
	ballIDCounter int32 //球的id计数器

	dummyUser *BattleUser

	ballUpdate []*util.UpdateBallInfo

	timer *util2.Timer //定时器
}

type Player struct {
	playerID int64
	battleUser *BattleUser
}

func (this *Game) GetTickCount() int32 {
	return this.tickCount
}

func (game *Game) SetUser(uid int32, user *BattleUser) {
	game.users[uid] = user
}

//接收玩家数据，同步给其他玩家
func (game *Game) AcceptPlayerData(playerID int64, data string, roles int32)  {
	//游戏结束不处理游戏内的数据转发
	//log.Debugf("accpet msg %v - %v", playerID, data)
	//push := &protocol.Response{Room: &protocol.Response_GameDataPush {
	//		GameDataPush: &protocol.GameDataPush{
	//			Data: data,
	//		},
	//	},
	//}
	//game.BroadcastOtherPlayer(-1, constant.RoleAll, push)
	//push := &protocol.Response{Room: &protocol.Response_GameDataRet{
	//	GameDataRet: &protocol.GameDataRet{
	//		Data: data,
	//	},
	//},
	//}
	req := &Request{}
	err := json.Unmarshal([]byte(data), req)
	if err != nil {
		log.Error(err.Error())
	}
	if req.Cmd == "EnterRoom" {
		EnterRoom(game, &Player{playerID: playerID + 1000})
	}
	if req.Cmd == "Move" {
		userID2BattleUser[playerID + 1000].Move(req)
	}
	if req.Cmd == "Spit" {
		userID2BattleUser[playerID + 1000].Spit()
	}
	if req.Cmd == "Split" {
		userID2BattleUser[playerID + 1000].Split()
	}
}


type Request struct {
	Cmd		string	`json:"cmd"`
	Dir     float64 `json:"dir"`
}

type Response struct {
	Cmd			string  	`json:"cmd"`
	ServerTick 	int32  		`json:"serverTick"`
}

/**
 * 向直播间输出游戏状态数据(V2, 增量或全量)  主播/嘉宾
 * @param stateData          游戏状态数据
 * @param type          数据类型 0 - 增量数据 1 - 全量数据
 * @param ts            时间戳 单位毫秒
 * @param forceUpdate   强制所有人更新这次全量数据
 */

//func (game *Game) AcceptPlayerMessage(playerID int64, request interface{}, response interface{}) {
//	log.Debugf("accept %v - %v - %v", playerID, request, response)
//
//
//}

func (self *Game) GetBallID() int32 {
	id := self.ballIDCounter
	self.ballIDCounter = self.ballIDCounter + 1
	return id
}

func (self *Game) Send2Client(resp interface{}, userID int64) {
	respData, _ := json.Marshal(resp)
	msg := &protocol.Response{
		Room: &protocol.Response_GameDataPush{
			GameDataPush: &protocol.GameDataPush{
				Data: string(respData),
			},
		},
	}
	self.SendToPlayer(userID, constant.RolePlayer, msg)
}

func (self *Game) Broadcast(resp interface{}) {
	respData, _ := json.Marshal(resp)
	msg := &protocol.Response{
		Room: &protocol.Response_GameDataPush{
			GameDataPush: &protocol.GameDataPush{
				Data: string(respData),
			},
		},
	}
	self.BroadcastOtherPlayer(-1, constant.RolePlayer, msg)
}

func (self *Game) GameOver() {
	log.Debug("Game Over!!!")
	if self.timer.IsActive() {
		//self.timer.Cancel()
		self.timer.Cancel()
	}
	msg := &protocol.Response{}
	self.BroadcastOtherPlayer(-1, constant.RolePlayer, msg)
	for _, v := range self.users {
		if v.Player != nil {
			v.Battle = nil
			v.Player = nil
		}
		v.Balls = nil
		userID2BattleUser[v.userID] = nil
	}
	self.dummyUser.Balls = nil
	self.dummyUser.Balls = nil

}

func (self *Game) GameUpdate() {

	self.updateCount = self.updateCount + 1
	nowSysTick := util.TimeNowSysTick()
	elapse := nowSysTick - self.lastSysTick
	self.lastSysTick = nowSysTick
	self.tickCount = self.tickCount + elapse
	var needSyncBallUpdate bool
	var syncElapse int32
	if self.updateCount % 1 == 0 {
		needSyncBallUpdate = true
		syncElapse = self.tickCount - self.lastSyncBallUpdate
		self.lastSyncBallUpdate = self.tickCount
	}
	log.Debug("GameTickUpdate tickCount:", self.tickCount, "gameOverTick:", self.gameOverTick)
	if self.tickCount >= self.gameOverTick {
		//游戏结束
		self.GameOver()
		//battles[self.id] = nil
		//delete(battles, self.id)
		self.Stop()
		return
	} else {
		self.dummyUser.Update(float64(elapse))
		//todo 暂时把ai去掉
		//self.aiMgr.Update()

		self.ballUpdate = []*util.UpdateBallInfo{}
		for _, v := range self.users {
			v.Update(float64(elapse))
			if needSyncBallUpdate {
				v.RefreshBallsUpdateInfo()
			}
		}
		if needSyncBallUpdate {
			for _, v := range self.users {
				v.NotifyBalls2Client(float64(syncElapse))
			}
		}
	}
	self.starMgr.Update()
}

func (self *Game) Enter(battleUser *BattleUser) {
	if battleUser.Battle != nil {
		log.Debug("first enter")
		self.users[int32(battleUser.userID)] = battleUser
		battleUser.Battle = self
	} else {
		log.Debug("reenter", len(battleUser.Balls))
	}
	elapse := util.TimeNowSysTick() - self.lastSysTick
	resp1 := &protocol2.ServerTickResp{Cmd:"ServerTick", ServerTick:self.tickCount + elapse}
	self.Send2Client(resp1, battleUser.userID - 1000)
	//respData1, _ := json.Marshal(resp1)
	//push1 := &protocol.Response{
	//	Room: &protocol.Response_GameDataPush{
	//		GameDataPush: &protocol.GameDataPush{
	//			Data: string(respData1),
	//		},
	//	},
	//}
	resp2 := &protocol2.EnterRoomResp{Cmd:"EnterRoom", Stars:self.starMgr.GetStarBits(), UserID: battleUser.userID}
	self.Send2Client(resp2, battleUser.userID - 1000)

	//self.BroadcastOtherPlayer(battleUser.userID, constant.RoleAll, push1)
	//self.BroadcastOtherPlayer(battleUser.userID, constant.RoleAll, push2)

	battleUser.ViewPort = &util.MapBorder{}
	battleUser.visionUser.ViewObjs = make(map[*vision.VisionObject]*vision.ViewObj)

	battleUser.visionUser.BeginSee = nil
	//battleUser.ViewPort := make(map[])
	if battleUser.BallCount == 0 {
		//创建玩家的球
		battleUser.Relive()
	}
	log.Debug("user enter OK")
}

//func getFreeRoom() {
//	var room
//	for _, v := range battles {
//		room = v
//		break
//	}
//
//	if room == nil {
//		room = NewGame()
//	}
//
//}
var userID2BattleUser = make(map[int64]*BattleUser)


func EnterRoom(room *Game, player *Player) {
	userID := player.playerID
	//battleUser := userID2BattleUser[userID]
	battleUser := userID2BattleUser[userID]
	//var room *Game
	if battleUser != nil {
		battleUser.Player = player
		room = battleUser.Battle
	} else {
		log.Debug("new battleuser, userID",userID)
		battleUser = NewBattleUser(player, userID, room)
		//battleUser.Battle = room
		userID2BattleUser[userID] = battleUser
	}
	room.Enter(battleUser)
}

//func NewVisionObject(_collision collision.ICollision,_vision vision.IVision, _visionObj vision.IVisionObject, pos util.Position, r float64) *vision.VisionObject{
//	return &vision.VisionObject{
//		Collision:_collision,
//		Vision:_vision,
//		Proxy:_visionObj,
//		Pos:pos,
//		R:r,
//
//	}
//}