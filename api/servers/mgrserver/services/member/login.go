package member

import (
	"fmt"
	"gsms/GSMSMerchant/api/modules/member"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
)

//LoginHandler 用户登录对象
type LoginHandler struct {
	c component.IContainer
	m member.IUser
}

//NewLoginHandler 创建登录对象
func NewLoginHandler(container component.IContainer) (u *LoginHandler) {
	return &LoginHandler{
		c: container,
		m: member.NewUser(container),
	}
}

//Handle 处理用户登录，登录成功后转跳到指定的系统
func (u *LoginHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("----------用户登录-----------")

	//检查输入参数
	if err := ctx.Request.Check("username", "password"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	userAccount := ctx.Request.GetString("username")
	passWord := ctx.Request.GetString("password")

	clientIP, err := ctx.Request.Http.GetClientIP()
	ctx.Log.Info("----------获取用户登录IP-----------", clientIP)
	if err != nil {
		return
	}

	//处理用户登录
	data, err := u.m.Login(userAccount, passWord, clientIP)
	if err != nil {
		return err
	}

	//签名
	state := member.LoginAdminState{
		UserID:      types.GetInt(data[0]["user_id"]),
		UserAccount: types.GetString(data[0]["user_account"]),
		AccType:     types.GetString(data[0]["acc_type"]),
		RefAccID:    types.GetString(data[0]["ref_acc_id"]),
		UserName:    types.GetString(data[0]["user_name"]),
	}
	fmt.Printf("用户信息%+v\n", state)

	//设置jwt数据
	ctx.Response.SetJWT(&state)
	//返回用户名和渠道信息
	return map[string]interface{}{
		"last_login_ip":   types.GetString(data[0]["last_login_ip"]),
		"last_login_time": types.GetString(data[0]["last_login_time"]),
		"login_account":   state.UserAccount,
		"name":            state.UserName,
		"acc_type":        state.AccType,
		"ref_acc_id":      state.RefAccID,
	}
}
