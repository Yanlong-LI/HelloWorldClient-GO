package user

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	conn2 "HelloWorldServer/model/Login"
	UserLogin "HelloWorldServer/packet/user/Login"
	"HelloWorldServer/packet/user/me"
	"crypto/rand"
	"fmt"
	"strconv"
)

func init() {
	route.Register(UserLogin.ForEmail{}, Login)
	route.Register(me.GetInfo{}, GetUserInfo)
	route.Register(UserLogin.Resuming{}, actionResuming)
}

func Login(login UserLogin.ForEmail, conn connect.Connector) {
	fmt.Printf("用户 %s 尝试登陆\n", login.Email)

	//conn.Send(UserLogin.Fail{Fail: trait.Fail{Message: "账户或密码错误", Code: 6005}})
	//return

	conn2.SignIn(conn.GetId(), me.Info{Id: strconv.Itoa(int(conn.GetId())), UserName: login.Email, Language: "zh-chs", Region: "China"})
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

func actionResuming(resuming UserLogin.Resuming, conn connect.Connector) {
	fmt.Printf("用户恢复登陆%s", resuming.Token)
	conn2.SignIn(conn.GetId(), me.Info{Id: strconv.Itoa(int(conn.GetId())), UserName: "未可知", Language: "zh-chs", Region: "China"})
	conn.Send(UserLogin.ResumingSuccess{})
}
