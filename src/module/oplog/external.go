package oplog

import "github.com/KylinHe/aliensboot-server/module/oplog/internal"

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)
