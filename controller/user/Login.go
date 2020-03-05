package user

import (
	"crypto/rand"
	"fmt"
	"github.com/yanlong-li/HelloWorld-GO/io/db"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/model"
	conn2 "github.com/yanlong-li/HelloWorldServer/model/online"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
	UserLogin "github.com/yanlong-li/HelloWorldServer/packetModel/user/Login"
	"github.com/yanlong-li/HelloWorldServer/packetModel/user/me"
	"strings"
	"time"
)

func init() {
	route.Register(UserLogin.ForEmail{}, Login)
	route.Register(me.GetInfo{}, GetUserInfo)
	route.Register(UserLogin.Resuming{}, actionResuming)
}

func Login(login UserLogin.ForEmail, conn connect.Connector) {
	fmt.Printf("用户 %s 尝试登陆\n", login.Email)

	login.Password = strings.Trim(login.Password, " 	\n\r")
	login.Email = strings.Trim(login.Email, " 	\n\r")
	if len(login.Password) == 0 || len(login.Email) == 0 {
		conn.Send(UserLogin.Fail{Fail: trait.Fail{Message: "账户或密码不能为空", Code: 6005}})
		return
	}

	userAccount := &model.UserAccount{}
	err := db.Model(userAccount).Where(map[interface{}]interface{}{"account": login.Email}).One()
	if err != nil {
		conn.Send(UserLogin.Fail{Fail: trait.Fail{Message: "账户或密码不正确", Code: 6005}})
		return
	}
	userPassword := &model.UserPassword{}
	err = db.Model(userPassword).Where(map[interface{}]interface{}{"user_id": userAccount.UserId, "password": login.Password}).One()
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
	userToken := &model.UserToken{UserId: user.Id, Token: token, ExpireTime: uint64(time.Now().AddDate(0, 1, 0).Unix()), CreateTime: uint64(time.Now().Unix())}
	err = db.Model(userToken).Insert()
	if err != nil {
		fmt.Println(err)
	}
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
