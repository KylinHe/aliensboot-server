package data

import "math"


//目前适配正方形MapWidth BlockWidth只要配置宽度即可

var AtarGame = &AtarGameBase {
	ScreenSizeFactor:30.0,
	InitScore:10,
	MaxScore:100000,
	Sp0:14,
	StarScore:1,
	MaxUserBallCount:16,
	MaxThornBallCount:10,
	SpitV0Factor:3.2,
	SplitV0Factor:25.0,
	CentripetalSpeedCoef:0.3,
	ThornCount:10,
	MinThornScore:80,
	MaxThornScore:100,
	SplitDuration:800,
	GameTime:300,

	MapWidth:4000,
	BlockWidth: 4000/50,

	VisibleWidth:1024,
	VisibleHeight:768,

	ThornColor: []int32{1,0,1,1},
	ThornColorID:22,

	Colors:[]Color{
		{0.85,0.40,0.79,1},
		{0.80,0.92,0.20,1},
		{0.34,0.77,0.28,1},
		{0.56,0.48,0.63,1},
		{0.37,0.52,0.96,1},
		{0.92,0.64,0.72,1},
		{0.15,0.61,0.02,1},
		{0.25,0.14,0.81,1},
		{0.16,0.41,0.13,1},
		{0.11,1.00,0.22,1},
		{0.52,0.84,0.62,1},
		{0.30,0.64,0.53,1},
		{0.50,0.98,0.30,1},
		{0.78,0.53,0.77,1},
		{0.41,0.90,0.29,1},
		{0.36,0.81,0.92,1},
		{0.07,0.95,0.53,1},
		{0.09,0.20,0.67,1},
		{0.90,0.35,0.07,1},
		{0.03,0.46,0.07,1},
		{0.24,0.98,0.91,1},
	},

}



// [球球大作战基础配置表]
type AtarGameBase struct {
	ScreenSizeFactor float64
	InitScore int32  //初始化分数
	MaxScore  int32 //最大分数
	Sp0 	int32
	StarScore int32
	MaxUserBallCount int32
	MaxThornBallCount int32
	SpitV0Factor float32
	SplitV0Factor	float32
	CentripetalSpeedCoef float32
	ThornCount int32
	MinThornScore int32
	MaxThornScore int32
	SplitDuration	int32
	GameTime int32
	MapWidth float32
	ThornColorID int32
	ThornColor []int32
	Colors []Color

	MaxObjs int32 //最大的子对象
	MaxLevel int32 //最大的等级

	BlockWidth int32 //块的宽度
	VisibleWidth  int32  //视野宽度
	VisibleHeight int32  //视野高度


}

type Color struct {
	R float32
	G float32
	B float32
	U int32
}

func (this *AtarGameBase) Score2R(score int32) float64 {
	return math.Sqrt(float64(score) * 0.165 + 0.6) * 50.0 * 0.01 * this.ScreenSizeFactor
}


func (this *AtarGameBase) SpeedByR(r float64, speedLev float64) float64 {
	if speedLev == 0.0 {
		speedLev = 1.0
	}
	r = r / this.ScreenSizeFactor
	return 1.6 * math.Min(5.0, 9.0 / (r + 1.0) + 0.7) * this.ScreenSizeFactor * speedLev

}

func (this *AtarGameBase) EatFactor(score float64) float64 {
	if score <= 20.0 {
		return 1.3
	}
	if score >= 10000.0 {
		return 1.05
	} else {
		return (-0.000025)*score + 1.3
	}
}


