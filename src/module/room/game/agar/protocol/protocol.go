package protocol

import "github.com/KylinHe/aliensboot-server/module/room/game/agar/util"

type BeginSeeResp struct {
	Cmd 	string `json:"cmd"`
	Timestamp int32 `json:"timestamp"`
	Balls []*BallInfo `json:"balls"`
}

type BallInfo struct {
	UserID		int64				`json:"userID"`
	ID			int32				`json:"id"`
	Pos			util.Position		`json:"pos"`
	R			float64				`json:"r"`
	Velocitys	[]*Velocity			`json:"velocitys"`
	Color		int32				`json:"color"`
}

type Velocity struct {
	Duration  float64 	`json:"duration"`
	AccRemain float64 	`json:"accRemain"`
	V         *Vector2D	`json:"v"`
	TargetV   *Vector2D `json:"targetV"`
	A 		  *Vector2D	`json:"a"`
}

type Vector2D struct {
	X float64	`json:"x"`
	Y float64	`json:"y"`
}


type BallUpdateResp struct {
	Cmd string                   `json:"cmd"`
	Timestamp int32              `json:"timestamp"`
	Elapse int32                 `json:"elapse"`
	Balls []*util.UpdateBallInfo `json:"balls"`
}

type EndSeeResp struct {
	Cmd 	string `json:"cmd"`
	Timestamp int32 `json:"timestamp"`
	Balls []int32 `json:"balls"`
}

type ServerTickResp struct {
	Cmd			string  	`json:"cmd"`
	ServerTick 	int32  		`json:"serverTick"`
}

type EnterRoomResp struct {
	Cmd			string  	`json:"cmd"`
	Timestamp 	int32 		`json:"timestamp"`
	Stars 		[]uint32 	`json:"stars"`
	UserID		int64		`json:"userID"`
}

type StarReliveResp struct {
	Cmd 		string 		`json:"cmd"`
	Stars		[]int32 	`json:"stars"`
	Timestamp	int32		`json:"timestamp"`
}

type StarDeadResp struct {
	Cmd 		string		`json:"cmd"`
	Stars 		[]int32		`json:"stars"`
	Timestamp   int32  		`json:"timestamp"`
}