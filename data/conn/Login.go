package conn

import (
	"HelloWorldServer/packet/user"
	"errors"
)

var Login map[uint64]user.User = make(map[uint64]user.User, 1)

func SignUp(CID uint64, _user user.User) {
	Login[CID] = _user
}

func Auth(CID uint64) (user.User, error) {

	if _user, ok := Login[CID]; ok {
		return _user, nil
	}
	return user.User{}, errors.New("验证未通过")

}
