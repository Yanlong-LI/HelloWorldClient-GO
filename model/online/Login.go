package online

import (
	"errors"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorldServer/packetModel/user"
)

// 连接id->User
var conUser = make(map[uint64]user.Info)

// userId -> [连接ID]->连接器
var userCons = make(map[uint64]map[uint64]connect.Connector)

var signInTask = make(chan struct {
	Conn connect.Connector
	User user.Info
})
var signOutTask = make(chan uint64)

var sendMessage = make(chan struct {
	UserId uint64
	Model  interface{}
})

func init() {

	go func() {
		for {
			select {
			case s := <-signInTask:
				conUser[s.Conn.GetId()] = s.User
				if _, ok := userCons[s.User.Id]; !ok {
					userCons[s.User.Id] = make(map[uint64]connect.Connector)
				}
				userCons[s.User.Id][s.Conn.GetId()] = s.Conn
			case CId := <-signOutTask:
				_user := conUser[CId]
				delete(conUser, CId)
				delete(userCons[_user.Id], CId)

				// 如果用户多端下线，销毁数据
				if len(userCons) <= 0 {
					delete(userCons, _user.Id)
				}

			case sendMessage := <-sendMessage:

				if _userCons, ok := userCons[sendMessage.UserId]; ok {

					for _, conn := range _userCons {
						_ = conn.Send(sendMessage.Model)
					}

				}

			}
		}

	}()

}

// 登陆
func SignIn(_conn connect.Connector, _user user.Info) {
	signInTask <- struct {
		Conn connect.Connector
		User user.Info
	}{
		_conn, _user,
	}
}

// 退出登陆
func SignOut(CID uint64) {
	signOutTask <- CID
}

// 验证连接是否登陆用户
func Auth(CId uint64) (user.Info, error) {

	if _user, ok := conUser[CId]; ok {
		return _user, nil
	}
	return user.Info{}, errors.New("验证未通过")
}

/**
用户是否在线
*/
func UserOnlineByUserId(UserId uint64) bool {

	if c, ok := userCons[UserId]; ok {
		if len(c) >= 1 {
			return true
		}
	}
	return false

}

func UserSendMessage(UserId uint64, Model interface{}) {
	sendMessage <- struct {
		UserId uint64
		Model  interface{}
	}{UserId: UserId, Model: Model}
}
