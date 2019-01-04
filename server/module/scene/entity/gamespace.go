/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/12/4
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package entity

import "github.com/KylinHe/aliensboot-core/mmo"

const (
	TypeGameSpace mmo.EntityType = "GameSpace"
)

type GameSpace struct {
	mmo.Space //
}

func (space *GameSpace) OnSpaceCreated() {
	// notify the SpaceService that it's ok
	//space.EnableAOI(100)
	//
	//goworld.CallService("SpaceService", "NotifySpaceLoaded", space.Kind, space.ID)
	//space.AddTimer(time.Second*5, "DumpEntityStatus")
	//space.AddTimer(time.Second*5, "SummonMonsters")
	//M := 10
	//for i := 0; i < M; i++ {
	//	space.CreateEntity("Monster", entity.Vector3{})
	//}
}
