package wechat

import (
	"gsms/GSMSMerchant/api/modules/app"
	"gsms/GSMSMerchant/api/modules/channel"
	"gsms/GSMSMerchant/api/modules/member"
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/wechat/mp/oauth2"
)

// IWeChat is
type IWeChatUser interface {
	//用户关注公众号
	Subscribe(openID string, extData types.XMap)
	//用户取消关注公众号
	UnSubscribe(openID string)
	//获取用户信息
	GetUserInfo(code, sceneID, referee string) (mem *member.LoginState, err error)
}

//WeChatUser is
type WeChatUser struct {
	c             component.IContainer
	db            IDbWeChatUser
	channelHandle channel.IChannel
}

//NewWechat is
func NewWechatUser(c component.IContainer) *WeChatUser {
	return &WeChatUser{
		c:             c,
		db:            NewDbWeChat(c),
		channelHandle: channel.NewChannel(c),
	}
}

//Subscribe 用户关注公众号
func (w *WeChatUser) Subscribe(openID string, extData types.XMap) {
	token, err := GetUtil().Token()
	fmt.Println("Subscribe.token:", token)
	if err != nil {
		fmt.Println("Subscribe.err:", err)
	}
	userInfo, err := oauth2.GetUserInfoDefault(token, openID, "", nil)
	if err != nil {
		fmt.Println("Subscribe.GetUserInfoDefault:", err)
		return
	}
	w.db.Subscribe(userInfo, extData)
}

//UnSubscribe 用户取消关注公众号
func (w *WeChatUser) UnSubscribe(openID string) {
	w.db.UnSubscribe(openID)
}

//GetUserInfo 获取用户信息
func (w *WeChatUser) GetUserInfo(code, sceneID, referee string) (mem *member.LoginState, err error) {
	//获取网页token
	conf := app.GetWechatConf(w.c)
	channelInfo, err := w.channelHandle.Get(conf.ChannelID)
	if err != nil {
		err = fmt.Errorf("WeChatUser.GetUserInfo.channelHandle(%v);%s", err, conf.ChannelID)
		return
	}
	appID := channelInfo.GetString("appid")
	token, err := oauth2.AuthCode(appID, channelInfo.GetString("app_secret"), code, nil)
	if err != nil {
		return nil, fmt.Errorf("获取网页token失败：err(%v)", err)
	}
	//拉取用户信息
	userinfo, err := oauth2.GetUserInfo(token.AccessToken, token.OpenID, "", nil)
	if err != nil {
		_ = fmt.Errorf("拉取用户信息失败：err(%v)", err)
		userinfo = &oauth2.UserInfo{OpenId: token.OpenID}
	}
	// fmt.Printf("拉取的用户信息：snsapi_userinfo(%v),\nimage:%v\n", userinfo, userinfo.HeadImageURL)

	//更新用户数据
	customerNo, err := w.db.SaveUser(userinfo, types.XMap{
		"scene":   sceneID,
		"referee": referee,
		"appid":   appID,
	})
	if err != nil {
		return nil, err
	}
	var m member.LoginState
	m.CustomerNo = customerNo
	m.NickName = userinfo.Nickname
	m.OpenID = userinfo.OpenId
	m.AppID = appID

	return &m, nil
}
