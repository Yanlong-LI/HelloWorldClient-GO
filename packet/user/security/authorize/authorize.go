package authorize

import "github.com/yanlong-li/HelloWorldServer/packet/trait"

// 获取授权
// scope 目前范围有 二 参考微信
// 一、 snsapi_base 仅获取 openid 不弹授权
// 二、 snsapi_userinfo  给用户提示授权页 获取所有公开资料
type GetAuthorize struct {
	ServerId uint64
	Scope    string
}

// 获取授权失败
/**
可能原因有 1、服务器未注册 2、服务器被封 3、服务器已注销 X、用户权限问题
*/
type GetAuthorizeFail struct {
	trait.Fail
}

// 获取授权成功 base
type GetAuthorizeSuccess struct {
	//授权令牌
	AuthorizeToken string
}

type GetAccessToken struct {
	ServerId         uint64
	Secret           string
	AuthorizeToken   string
	IdentificationId string
}
