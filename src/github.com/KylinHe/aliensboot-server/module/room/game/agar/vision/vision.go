/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2019/2/16
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package vision



type VObject struct {

	proxy IVObject //代理实现对象

	viewObjs map[*VObject]struct{} //视野内对象

	enterSee bool

	ref int

}

func (self *VObject) IsRealUser() bool {
	return self.proxy.IsRealUser()
}

type IVObject interface {

	IsRealUser() bool //是否真实玩家

	PackOnBeginSee() struct{} //
}

