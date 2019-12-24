package user

import (
	"HelloWorld/io/db"
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/model"
	conn2 "HelloWorldServer/model/online"
	"HelloWorldServer/packet/trait"
	UserLogin "HelloWorldServer/packet/user/Login"
	"HelloWorldServer/packet/user/me"
	"crypto/rand"
	"fmt"
)

func init() {
	route.Register(UserLogin.ForEmail{}, Login)
	route.Register(me.GetInfo{}, GetUserInfo)
	route.Register(UserLogin.Resuming{}, actionResuming)
}

func Login(login UserLogin.ForEmail, conn connect.Connector) {
	fmt.Printf("用户 %s 尝试登陆\n", login.Email)

	userAccount := &model.UserAccount{}
	err := db.Find(userAccount).Where(map[interface{}]interface{}{"account": login.Email}).One()
	if err != nil {
		conn.Send(UserLogin.Fail{Fail: trait.Fail{Message: "账户或密码不正确", Code: 6005}})
		return
	}
	userPassword := &model.UserPassword{}
	err = db.Find(userPassword).Where(map[interface{}]interface{}{"user_id": userAccount.UserId, "password": login.Password}).One()
	if err != nil {
		conn.Send(UserLogin.Fail{Fail: trait.Fail{Message: "账户或密码不正确", Code: 6005}})
		return
	}

	user, err := model.GetUserById(userAccount.UserId)
	if err != nil {
		conn.Send(UserLogin.Fail{Fail: trait.Fail{Message: "账户或密码不正确", Code: 6005}})
		return
	}

	conn2.SignIn(conn, me.Info{Id: user.Id, Nickname: user.Nickname, Avatar: user.Avatar, Language: user.Language, Region: user.Region})
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	token := fmt.Sprintf("%x", b)
	conn.Send(UserLogin.Success{Token: token})
}

// 获取自身用户信息
func GetUserInfo(info me.GetInfo, conn connect.Connector) {
	_user, err := conn2.Auth(conn.GetId())
	if err != nil {
		return
	}
	conn.Send(_user)
}

// 恢复登陆
func actionResuming(resuming UserLogin.Resuming, conn connect.Connector) {
	fmt.Printf("用户恢复登陆%s", resuming.Token)
	_user, err := model.GetUserByToken(resuming.Token)
	if err != nil {
		fmt.Println(err)
		conn.Send(UserLogin.ResumingFail{Fail: trait.Fail{Message: "Token无效"}})
		return
	}
	conn2.SignIn(conn, me.Info{Id: _user.Id, Nickname: _user.Nickname, Language: _user.Language, Region: _user.Region})
	conn.Send(UserLogin.ResumingSuccess{})
}
