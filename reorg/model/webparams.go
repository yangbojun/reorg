package model

type Web struct {
	Db
	Error string
	IsCompact string
	IsRebuild string
	Parrel string
	ParrelInfo string
}

func NewWeb() *Web {
	return &Web{IsCompact: "checked", IsRebuild: "", Parrel: "2",
		ParrelInfo: "执行碎片整理并发数,请视数据库性能配置,建议不要超过5个",}
}