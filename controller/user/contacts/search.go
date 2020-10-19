package contacts

import (
	"github.com/yanlong-li/hi-go-server/common"
	"github.com/yanlong-li/hi-go-server/model"
	"github.com/yanlong-li/hi-go-server/packet_model/user"
	"github.com/yanlong-li/hi-go-server/packet_model/user/contacts"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(contacts.SearchUser{}, searchUser)
}

// 搜索账户
func searchUser(searchUser contacts.SearchUser, connector connect.Connector) {

	selfUser, _ := common.Auth(connector.GetId())

	_userAccountModel := &model.UserAccount{}

	err := db.Model(_userAccountModel).Find().Where("=", "account", searchUser.Account).AndWhere("!=", "user_id", selfUser.Id).One()
	if !err.Status() || err.Empty() {
		_ = connector.Send(contacts.SearchUserFail{Fail: struct {
			Code    uint32
			Message string
		}{Code: 1, Message: "未搜索到账户"}})
		return
	}
	_user, err := _userAccountModel.GetUser()
	if !err.Status() || err.Empty() {
		_ = connector.Send(contacts.SearchUserFail{Fail: struct {
			Code    uint32
			Message string
		}{Code: 1, Message: "搜索异常，用户数据丢失"}})
		return
	}
	_userInfo := user.Info{Id: _user.Id, Nickname: _user.Nickname, Avatar: _user.Avatar, Region: _user.Region, Language: _user.Language}
	_ = connector.Send(contacts.SearchUserSuccess{
		Info: _userInfo,
	})
}
