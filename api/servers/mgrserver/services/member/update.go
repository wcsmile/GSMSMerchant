package member

import (
	"fmt"
	"gsms/GSMSMerchant/api/modules/const/errorcode"
	"gsms/GSMSMerchant/api/modules/member"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

//UpdateHandler 用户登录对象
type UpdateHandler struct {
	c component.IContainer
	m member.IUser
}

//NewUpdateHandler 创建登录对象
func NewUpdateHandler(container component.IContainer) (u *UpdateHandler) {
	return &UpdateHandler{
		c: container,
		m: member.NewUser(container),
	}
}

//Handle 处理用户修改密码
func (u *UpdateHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------修改密码----------")
	//检查输入参数

	ctx.Log.Info("1. 获取参数")
	if err := ctx.Request.Check("password", "password_old", "passwords"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	password := ctx.Request.GetString("password")
	passwords := ctx.Request.GetString("passwords")
	passwordOld := ctx.Request.GetString("password_old")
	ctx.Log.Info("password", password, passwords, passwordOld)
	if !strings.EqualFold(password, passwords) {
		return context.NewError(errorcode.HTTPErrorBankCardPWDError, "两次密码不相同")
	}
	//签名
	mem := member.GetMember(ctx)
	fmt.Println("mem", mem)

	//处理密码修改
	ctx.Log.Info("2. 执行操作")
	err := u.m.ChangePwd(fmt.Sprintf("%d", mem.UserID), passwordOld, password)
	if err != nil {
		return err
	}
	ctx.Log.Info("3. 返回数据")
	return "success"
}
