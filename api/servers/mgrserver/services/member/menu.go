package member

import (
	"fmt"
	"gsms/GSMSMerchant/api/modules/member"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

//MenuHandler is
type MenuHandler struct {
	c component.IContainer
	m member.IUser
}

//NewMenuHandler is
func NewMenuHandler(container component.IContainer) (u *MenuHandler) {
	return &MenuHandler{
		c: container,
		m: member.NewUser(container),
	}
}

//Handle 获取菜单
func (u *MenuHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------获取菜单----------")
	ctx.Log.Info("1. 获取参数")
	mem := member.GetMember(ctx)

	ctx.Log.Info("2. 执行操作")
	Menus, err := u.m.GetMenu(fmt.Sprintf("%d", mem.UserID))
	if err != nil {
		return err
	}
	ctx.Log.Info("3. 返回数据")
	return Menus
}
