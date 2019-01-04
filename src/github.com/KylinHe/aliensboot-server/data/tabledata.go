package data


// [兵种表]
type Army struct {
	
	Tid int32 `json:"tid"`   //编号
	
	Name string `json:"name"`   //名称
	
	ShowName string `json:"show_name"`   //显示名称
	
	EnumTid int32 `json:"enum_tid"`   //枚举编号
	
	ArmyTypeBase int32 `json:"army_type_base"`   //兵种初始形态
	
	Display int32 `json:"display"`   //显示排序
	
	TEST int32 `json:"TEST"`   //试数数组
	
}

// [兵种形态表]
type ArmyType struct {
	
	Tid int32 `json:"tid"`   //编号
	
	Name string `json:"name"`   //名称
	
	English string `json:"english"`   //英文
	
	Icon int32 `json:"icon"`   //图标
	
	LvStr string `json:"lvStr"`   //等级
	
	Desc string `json:"desc"`   //描述
	
	Army int32 `json:"army"`   //兵种
	
	GenerationFlag int32 `json:"generation_flag"`   //代数标记
	
	Reduction int32 `json:"reduction"`   //减伤系数
	
	Dodge float32 `json:"dodge"`   //闪避率
	
	Armys int32 `json:"armys"`   //兵种数组依赖
	
	Testjson int32 `json:"testjson"`   //测试json
	
}
