package member

import (
	"encoding/json"

	"github.com/micro-plat/hydra/context"
)

const maxErrorCnt = 5

//MemberInfo 用户信息
type MemberInfo struct {
	CustomerNo   string `json:"customer_no" m2s:"customer_no"`
	OpenID       string `json:"openid" m2s:"openid"`
	NickName     string `json:"nick_name" m2s:"nick_name"`
	IsFans       string `json:"is_fans" m2s:"is_fans"`
	ValidateType string `json:"validate_type"`
	ActivityTag  string `json:"activity_tag"`
}

//LoginState 用户登录状态
type LoginState MemberInfo

//MarshalJSON 修改marshal行为，去掉敏感字段
func (m LoginState) MarshalJSON() ([]byte, error) {
	type mem MemberInfo
	current := mem(m)

	return json.Marshal((*mem)(&current))
}

//Save 保存member信息
func Save(ctx *context.Context, m *LoginState) error {

	ctx.Meta.Set("login-state", m)
	return nil
}

//Get 获取member信息
func Get(ctx *context.Context) *LoginState {
	//var v interface{}
	v, _ := ctx.Meta.Get("login-state")
	//fmt.Println("\n\n获取member信息:", v)
	if v == nil {
		return nil
	}
	return v.(*LoginState)
}
