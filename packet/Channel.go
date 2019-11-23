package packet

type Channel struct {
	Id     uint64 //频道ID
	Name   string //频道名称
	Avatar string //频道图片
	Holder uint64 //持有人id
}

type ChannelList struct {
	List []Channel
}

type GetChannelList struct {
}
