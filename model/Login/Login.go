package Login

import (
	"HelloWorldServer/packet/user/me"
	"errors"
)

var Login = make(map[uint64]me.Info, 1)
var signInTask = make(chan struct {
	Id   uint64
	User me.Info
})
var signOutTask = make(chan uint64)

func init() {

	go func() {
		for {
			select {
			case s := <-signInTask:
				Login[s.Id] = s.User
			case CId := <-signOutTask:
				delete(Login, CId)
			}
		}

	}()

}

// 登陆
func SignIn(CId uint64, _user me.Info) {
	signInTask <- struct {
		Id   uint64
		User me.Info
	}{
		CId, _user,
	}
}

// 退出登陆
func SignOut(CID uint64) {
	signOutTask <- CID
}

// 验证连接是否登陆用户
func Auth(CId uint64) (me.Info, error) {

	if _user, ok := Login[CId]; ok {
		return _user, nil
	}
	return me.Info{}, errors.New("验证未通过")

}
