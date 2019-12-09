package user

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/packet/user"
	UserLogin "HelloWorldServer/packet/user/Login"
	"crypto/rand"
	"fmt"
)

func init() {
	route.Register(UserLogin.ForEmail{}, Login)
	route.Register(user.GetInfo{}, GetUserInfo)
}

func Login(login UserLogin.ForEmail, conn connect.Connector) {
	fmt.Printf("用户 %s 尝试登陆\n", login.Email)
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	token := fmt.Sprintf("%x", b)
	conn.Send(UserLogin.Success{Token: token})
}

// 获取自身用户信息
func GetUserInfo(info user.GetInfo, conn connect.Connector) {
	conn.Send(user.User{Id: "123412312", UserName: "1231231", Language: "zh-hans", Region: "asia/china/jiangsu/nanjing"})
}
