package contacts

import (
	"github.com/yanlong-li/HelloWorld-GO/io/db"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/model"
	"github.com/yanlong-li/HelloWorldServer/model/online"
	"github.com/yanlong-li/HelloWorldServer/packetModel/user"
	"github.com/yanlong-li/HelloWorldServer/packetModel/user/contacts"
)

func init() {
	route.Register(contacts.SearchUser{}, searchUser)
}

// 搜索账户
func searchUser(searchUser contacts.SearchUser, connector connect.Connector) {

	selfUser, _ := online.Auth(connector.GetId())

	_userModel := &model.User{}

	err := db.Model(_userModel).Find().Where("=", "account", searchUser.Account).AndWhere("!=", "id", selfUser.Id).One()
	if !err.Status() || err.Empty() {
		_ = connector.Send(contacts.SearchUserFail{Fail: struct {
			Code    uint32
			Message string
		}{Code: uint32(1), Message: "未搜索到账户"}})
		return
	}
	_userInfo := user.Info{Id: _userModel.Id, Nickname: _userModel.Nickname, Avatar: _userModel.Avatar, Region: _userModel.Region, Language: _userModel.Language}
	_ = connector.Send(contacts.SearchUserSuccess{
		Info: _userInfo,
	})
}
