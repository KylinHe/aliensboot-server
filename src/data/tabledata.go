package data


// [兵种表]
type Army struct {
	
	Tid int32 `json:"tid"`   //编号
	
	ScreenSizeFactor float32 `json:"screenSizeFactor"`   //名称
	
	InitScore string `json:"initScore"`   //初始化分数
	
	XScore int32 `json:"xScore"`   //枚举编号
	
	ArmyTypeBase int32 `json:"army_type_base"`   //兵种初始形态
	
	Display int32 `json:"display"`   //显示排序
	
	TEST []string `json:"TEST"`   //试数数组
	
}
