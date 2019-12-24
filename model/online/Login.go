package online

import (
	"HelloWorld/io/network/connect"
	"HelloWorldServer/packet/user/me"
	"errors"
)

var conUser = make(map[uint64]me.Info, 1)
var userCons = make(map[uint64]map[uint64]connect.Connector)

var signInTask = make(chan struct {
	Conn connect.Connector
	User me.Info
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
			case sendMessage := <-sendMessage:

				if _userCons, ok := userCons[sendMessage.UserId]; ok {

					for _, conn := range _userCons {
						conn.Send(sendMessage.Model)
					}

				}

			}
		}

	}()

}

// 登陆
func SignIn(_conn connect.Connector, _user me.Info) {
	signInTask <- struct {
		Conn connect.Connector
		User me.Info
	}{
		_conn, _user,
	}
}

// 退出登陆
func SignOut(CID uint64) {
	signOutTask <- CID
}

// 验证连接是否登陆用户
func Auth(CId uint64) (me.Info, error) {

	if _user, ok := conUser[CId]; ok {
		return _user, nil
	}
	return me.Info{}, errors.New("验证未通过")
}

func UserSendMessage(UserId uint64, Model interface{}) {
	sendMessage <- struct {
		UserId uint64
		Model  interface{}
	}{UserId: UserId, Model: Model}
}
