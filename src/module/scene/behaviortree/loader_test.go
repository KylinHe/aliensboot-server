package loader

import (
	"fmt"
	"testing"

	b3 "github.com/magicsea/behavior3go"
	//. "github.com/magicsea/behavior3go/actions"
	//. "github.com/magicsea/behavior3go/composites"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
	//. "github.com/magicsea/behavior3go/decorators"
	"github.com/magicsea/behavior3go/loader"
)

///////////////////////加载事例///////////////////////////
//自定义action节点
type LogTest struct {
	Action
	info string
}

func (this *LogTest) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.info = setting.GetPropertyAsString("info")
}

func (this *LogTest) OnTick(tick *Tick) b3.Status {
	fmt.Println("logtest:", this.info)
	return b3.SUCCESS
}

func TestLoadTree(t *testing.T) {
	treeConfig, ok := LoadTreeCfg("mytree.json")
	if ok {
		//自定义节点注册
		maps := b3.NewRegisterStructMaps()
		maps.Register("Log", new(LogTest))

		//载入
		tree := loader.CreateBevTreeFromConfig(treeConfig, maps)
		tree.Print()

		//输入板
		//board := NewBlackboard()
		////循环每一帧
		//for i := 0; i < 5; i++ {
		//	tree.Tick(i, board)
		//}
	} else {
		t.Error("LoadTreeCfg err")
	}
}
