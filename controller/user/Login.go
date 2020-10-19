package user

import (
	"crypto/rand"
	"fmt"
	"github.com/yanlong-li/hi-go-server/common"
	"github.com/yanlong-li/hi-go-server/model"
	"github.com/yanlong-li/hi-go-server/packet_model/trait"
	"github.com/yanlong-li/hi-go-server/packet_model/user"
	UserLogin "github.com/yanlong-li/hi-go-server/packet_model/user/Login"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
	"strings"
	"time"
)

func init() {
	route.Register(UserLogin.ForEmail{}, Login)
	route.Register(user.GetInfo{}, GetUserInfo)
	route.Register(UserLogin.Resuming{}, actionResuming)
}

func Login(login UserLogin.ForEmail, conn connect.Connector) {
	fmt.Printf("用户 %s 尝试登陆\n", login.Email)
	var ormErr db.OrmError

	login.Password = strings.Trim(login.Password, " 	\n\r")
	login.Email = strings.Trim(login.Email, " 	\n\r")
	if len(login.Password) == 0 || len(login.Email) == 0 {
		_ = conn.Send(UserLogin.Fail{Fail: trait.Fail{Message: "账户或密码不能为空", Code: 6005}})
		return
	}

	userAccount := &model.UserAccount{}
	ormErr = db.Model(userAccount).Find().Where(map[interface{}]interface{}{"account": login.Email}).One()
	if !ormErr.Status() {
		_ = conn.Send(UserLogin.Fail{Fail: trait.Fail{Message: "账户或密码不正确", Code: 6005}})
		return
	}
	userPassword := &model.UserPassword{}
	ormErr = db.Model(userPassword).Find().Where(map[interface{}]interface{}{"user_id": userAccount.UserId, "password": login.Password, "type": 0}).One()
	if ormErr.Empty() {
		_ = conn.Send(UserLogin.Fail{Fail: trait.Fail{Message: "账户或密码不正确", Code: 6005}})
		return
	}

	_user, err := model.GetUserById(userAccount.UserId)
	if !err.Status() {
		_ = conn.Send(UserLogin.Fail{Fail: trait.Fail{Message: "账户或密码不正确", Code: 6005}})
		return
	}

	common.SignIn(conn, user.Info{Id: _user.Id, Nickname: _user.Nickname, Avatar: _user.Avatar, Language: _user.Language, Region: _user.Region})
	b := make([]byte, 64)
	_, _ = rand.Read(b)
	token := fmt.Sprintf("%x", b)
	userToken := &model.UserToken{UserId: _user.Id, Token: token, ExpireTime: uint64(time.Now().AddDate(0, 1, 0).Unix()), CreateTime: uint64(time.Now().Unix())}
	ormErr = db.Model(userToken).Insert().Insert()
	if !ormErr.Status() {
		logger.Fatal("写入token失败", 0)
	}
	_ = conn.Send(UserLogin.Success{Token: token})
}

// 获取自身用户信息
func GetUserInfo(info user.GetInfo, conn connect.Connector) {
	_user, err := common.Auth(conn.GetId())
	if err != nil {
		return
	}
	_ = conn.Send(_user)
}

// 恢复登陆
func actionResuming(resuming UserLogin.Resuming, conn connect.Connector) {
	fmt.Printf("用户恢复登陆%s", resuming.Token)
	_user, err := model.GetUserByToken(resuming.Token)
	if err.Empty() {
		fmt.Println(err)
		_ = conn.Send(UserLogin.ResumingFail{Fail: trait.Fail{Message: "Token无效"}})
		return
	}
	common.SignIn(conn, user.Info{Id: _user.Id, Nickname: _user.Nickname, Language: _user.Language, Region: _user.Region})
	_ = conn.Send(UserLogin.ResumingSuccess{})
}
