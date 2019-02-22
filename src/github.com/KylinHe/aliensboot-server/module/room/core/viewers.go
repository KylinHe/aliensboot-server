/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2019/1/18
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

type Viewers map[int64]*Player //观众


func (this Viewers) AddViewer(player *Player) {
	if player != nil {
		this[player.GetId()] = player
	}
}

func (this Viewers) GetViewer(id int64) *Player {
	return this[id]
}

func (this Viewers) RemoveViewer(id int64) {
	delete(this, id)
}

func (this Viewers) ForeachViewer(callback func(player *Player)) {
	for _, player := range this {
		callback(player)
	}
}






