package member

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

//Conf 应用程序配置
type Conf struct {
	Secret  string `json:"secret" valid:"ascii,required"`
	SSOHost string `json:"sso_host" valid:"ascii,required"`
	Ident   string `json:"ident"`
}

//LoginAdminState 用户信息
type LoginAdminState struct {
	UserID         int    `json:"user_id" m2s:"user_id"`
	UserName       string `json:"user_name" m2s:"user_name"`
	UserAccount    string `json:"user_account" m2s:"user_account"`
	RoleName       string `json:"role_name" m2s:"role_name"`
	RoleID         int    `json:"role_id"`
	Status         int    `json:"status" m2s:"status"`
	AccType        string `json:"acc_type" m2s:"acc_type"`
	RefAccID       string `json:"ref_acc_id" m2s:"ref_acc_id"`
	IndexURL       string `json:"index_url"`
	Code           string `json:"code"`
	ProfilePercent int    `json:"profile_percent"`
	LoginTimeout   int    `json:"login_timeout" m2s:"login_timeout"`
	Timestamp      int64  `json:"timestamp"`
	Secret         string `json:"secret"`
}

//SaveMember 保存member信息
func SaveMember(ctx *context.Context, m *LoginAdminState) error {
	ctx.Meta.Set("gsms-admin-login-state", m)
	return nil
}

//GetMember 获取member信息
func GetMember(ctx *context.Context) *LoginAdminState {
	v, _ := ctx.Meta.Get("gsms-admin-login-state")
	if v == nil {
		return nil
	}
	return v.(*LoginAdminState)
}

//Valid 验证配置参数是否合法
func (c Conf) Valid() error {
	if b, err := govalidator.ValidateStruct(&c); !b {
		return fmt.Errorf("app 配置文件有误:%v", err)
	}
	return nil
}

//GetAppSecret 获取微信AppSecret
func (c *Conf) GetAppSecret() string {
	return c.Secret
}

//GetSystemName 获取系统名称
func (c *Conf) GetIdent() string {
	return c.Ident
}

//GetLoginURL .
func (c *Conf) GetLoginURL() string {
	return c.SSOHost + "/subsys/login"
}

//GetGraphCodeURL .
func (c *Conf) GetUserInfoURL() string {
	return c.SSOHost + "/subsys/user/info"
}

//GetMenuURL .
func (c *Conf) GetMenuURL() string {
	return c.SSOHost + "/subsys/menu"
}

//GetLoginURL 获取sso服务器修改密码地址
func (c *Conf) GetChangePwdURL() string {
	return c.SSOHost + "/subsys/pwd"
}

//GetSysInfoURL 或取系统信息链接
func (c *Conf) GetSysInfoURL() string {
	return c.SSOHost + "/subsys/info"
}

//SaveConf 保存当前应用程序配置
func SaveConf(c component.IContainer, m *Conf) {
	c.Set("__AppConf__", m)
}

//GetConf 获取当前应用程序配置
func GetConf(c component.IContainer) *Conf {
	return c.Get("__AppConf__").(*Conf)
}
