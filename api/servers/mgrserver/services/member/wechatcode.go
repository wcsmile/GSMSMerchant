package member

// import (
// 	"coupon/customer/recv-coupon/modules/member"
// 	"coupon/customer/recv-coupon/modules/wechat"
// 	"encoding/json"
// 	"strings"

// 	"github.com/micro-plat/hydra/component"
// 	"github.com/micro-plat/hydra/context"
// )

// //WechatCodeHandler
// type WechatCodeHandler struct {
// 	c component.IContainer
// 	m member.IUser
// }

// //NewWechatCodeHandler
// func NewWechatCodeHandler(container component.IContainer) (u *WechatCodeHandler) {
// 	return &WechatCodeHandler{
// 		c: container,
// 		m: member.NewUser(container),
// 	}
// }

// //Handle
// func (u *WechatCodeHandler) Handle(ctx *context.Context) (r interface{}) {

// 	//检查输入参数
// 	if err := ctx.Request.Check("username", "ident"); err != nil {
// 		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
// 	}

// 	userName := ctx.Request.GetString("username")

// 	//处理用户登录
// 	info, err := u.m.GetUserInfo(userName, ctx.Request.GetString("ident"))
// 	if err != nil {
// 		ctx.Response.SetStatus(context.ERR_FORBIDDEN)
// 		return err
// 	}

// 	extParams := info.ExtParams

// 	senduser := GetSendUser(extParams)
// 	if senduser == "" {
// 		senduser = userName
// 	}

// 	data, err := wechat.SendValidCode(senduser)
// 	if err != nil {
// 		return
// 	}
// 	ctx.Response.SetPlain()
// 	return data //返回图形验证码
// }

// func GetSendUser(extParams string) (senduser string) {
// 	extObj := map[string]string{}
// 	if strings.EqualFold(extParams, "") {
// 		return
// 	}
// 	if err := json.Unmarshal([]byte(extParams), &extObj); err != nil {
// 		return
// 	}
// 	if extSenduser, ok := extObj["senduser"]; ok {
// 		senduser = extSenduser
// 	}
// 	return
// }
