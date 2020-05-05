package model

type Profile struct {
	Id          int
	Name        string
	Gender      string
	//Age         int
	//Height      int
	//Weight      int
	Age         string
	Height      string
	Weight      string
	Income      string // 收入
	Marriage    string // 婚姻状况
	Education   string // 教育
	Occupation  string // 职业
	NativePlace string // 籍贯
	WorkPlace   string // 工作地
	Xinzuo      string // 星族
	House       string // 房子
	Car         string // 车子
}
