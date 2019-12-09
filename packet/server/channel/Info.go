package channel

type Info struct {
	Name       string
	Icon       string
	CreateUser string
	OwnerUser  string // 实际掌控着
	CreateTime uint64
	Public     bool //是否公开
	Verify     bool // 是否经过验证
	Commerce   bool // 是否可商业
}
