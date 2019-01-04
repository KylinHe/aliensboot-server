/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/12/6
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package utils

import (
	"github.com/KylinHe/aliensboot-core/mmo"
	"github.com/KylinHe/aliensboot-core/mmo/unit"
	"github.com/KylinHe/aliensboot-server/protocol"
	"encoding/json"
)

func BuildVector(vector unit.Vector) *protocol.Vector {
	return &protocol.Vector{
		X:float32(vector.X),
		Y:float32(vector.Y),
		Z:float32(vector.Z),
	}
}

func BuildEntity(entity mmo.Entity, self bool) *protocol.Entity {
	var data []byte = nil
	if self {
		data, _ = json.Marshal(entity.GetClientData())
	} else {
		data, _ = json.Marshal(entity.GetAllClientData())
	}
	return &protocol.Entity{
		Id:string(entity.GetID()),
		TypeName:string(entity.GetType()),
		Position:BuildVector(entity.Position),
		Yaw:float32(entity.Yaw),
		Attr:data,
	}
}
