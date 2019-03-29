package member

import (
	"gsms/GSMSMerchant/api/modules/member"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

//InfoHandler is
type InfoHandler struct {
	c component.IContainer
	m member.IUser
}

//NewInfoHandler is
func NewInfoHandler(container component.IContainer) (u *InfoHandler) {
	return &InfoHandler{
		c: container,
		m: member.NewUser(container),
	}
}

//Handle 获取菜单
func (u *InfoHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------获取系统信息----------")
	ctx.Log.Info("1. 获取参数")

	ctx.Log.Info("2. 执行操作")
	info, err := u.m.GetSysInfo()

	if err != nil {
		return err
	}

	ctx.Log.Info("3. 返回数据")
	return info
}
