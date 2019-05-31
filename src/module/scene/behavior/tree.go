/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/12/13
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package behavior

import (
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-server/module/scene/behavior/actions"
	"github.com/KylinHe/aliensboot-server/module/scene/behavior/composites"
	"github.com/KylinHe/aliensboot-server/module/scene/behavior/conditions"
	b3 "github.com/magicsea/behavior3go"
	b3config "github.com/magicsea/behavior3go/config"
	b3core "github.com/magicsea/behavior3go/core"
	b3loader "github.com/magicsea/behavior3go/loader"
)

func Init() {

}

//创建一个行为树
var trees map[string]*b3core.BehaviorTree

func CreateTree(path string) *b3core.BehaviorTree {
	b, ok := trees[path]
	if ok {
		return b
	}
	log.Info("create tree:%v", path)
	config, ok := b3config.LoadTreeCfg(path)
	if !ok {
		log.Fatal("LoadTreeCfg fail:" + path)
	}
	extMaps := createExtStructMaps()
	tree := b3loader.CreateBevTreeFromConfig(config, extMaps)
	tree.Print()
	trees[path] = tree
	return tree
}

//自定义的节点
func createExtStructMaps() *b3.RegisterStructMaps {

	st := b3.NewRegisterStructMaps()

	//actions
	st.Register("RandWait", &actions.RandWait{})
	st.Register("RandMove", &actions.RandMove{})
	st.Register("NormalAttack", &actions.NormalAttack{})
	st.Register("FindTarget", &actions.FindTarget{})

	//conditions
	st.Register("HaveTarget", &conditions.HaveTarget{})
	st.Register("HpLess", &conditions.HpLess{})

	//composite
	st.Register("Random", &composites.RandomComposite{})
	st.Register("Parallel", &composites.ParallelComposite{})
	return st
}
