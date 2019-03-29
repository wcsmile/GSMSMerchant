package wechat

import (
	"gsms/GSMSMerchant/api/modules/const/sql"
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/wechat/mp/oauth2"
)

//IDbWeChatUser is
type IDbWeChatUser interface {
	Subscribe(userInfo *oauth2.UserInfo, extData types.XMap)
	UnSubscribe(openID string)
	SaveUser(userInfo *oauth2.UserInfo, extData types.XMap) (customerNo string, err error)
}

//DbWechatUser is
type DbWechatUser struct {
	c component.IContainer
}

//NewDbWeChat is
func NewDbWeChat(c component.IContainer) *DbWechatUser {
	return &DbWechatUser{
		c: c,
	}
}

//UnSubscribe 取消关注,更新状态
func (d *DbWechatUser) UnSubscribe(openID string) {
	db := d.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.UpdateUnSubscribe, map[string]interface{}{
		"openid":  openID,
		"is_fans": 1,
	})
	if err != nil {
		fmt.Printf("用户取消关注发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
}

//Subscribe 用户关注，保存用户数据
func (d *DbWechatUser) Subscribe(userInfo *oauth2.UserInfo, extData types.XMap) {
	if userInfo == nil {
		return
	}
	db := d.c.GetRegularDB()
	openID := userInfo.OpenId
	//检查用户是否存在
	data, q, a, err := db.Query(sql.QueryUserInfo, map[string]interface{}{
		"openid": openID,
	})

	if err != nil {
		fmt.Printf("查询用户信息错误(err:%v),sql:%s,输入参数:%v,data:%v", err, q, a, data)
	}

	if data.Get(0).GetString("customer_no") != "" {
		d.updateUser(userInfo, extData)
		return
	}

	d.addUser(userInfo, extData)
}

func (d *DbWechatUser) addUser(userInfo *oauth2.UserInfo, extData types.XMap) (customerNo string, err error) {
	if userInfo == nil {
		return "", nil
	}
	if val, ok := extData["is_fans"]; !ok || val == "" {
		extData["is_fans"] = "1"
	}
	if val, ok := extData["source"]; !ok || val == "" {
		extData["source"] = "0"
	}
	db := d.c.GetRegularDB()

	noData, _, _, err := db.Scalar(sql.GetCustomerNo, nil)

	_, q, a, err := db.Execute(sql.SaveCustomerInfo, map[string]interface{}{
		"customer_no": noData,
		"appid":       extData.GetString("appid"), // | varchar2(32)  |         | 否    | IS,UNQ    | appid                            |
		"openid":      userInfo.OpenId,            // | varchar2(32)  |         | 否    | IS,UNQ    | openid                           |
		"unionid":     userInfo.UnionId,           // | varchar2(32)  |         | 是    | IS,UNQ    | unionid                          |
		"nick_name":   userInfo.Nickname,          // | varchar2(32)  |         | 是    | IS        | 昵称                             	|
		"head_url":    userInfo.HeadImageURL,      // | varchar2(128) |         | 是    | IS        | 头像地址                          |
		"status":      0,                          // | number(1)     |         | 否    | IS        | 状态（0：启用1.禁用 2.锁定）      	|
		"is_fans":     extData["is_fans"],         // | number(1)     | 1       | 否    | IS        | 是否是粉丝（0：是 1：否）         	|
		"source":      extData["source"],          // | number(1)     | 0       | 否    | IS        | 用户来源（0.自由用户 1.推荐用户） 	|
		"referee":     extData["referee"],
		"scene":       extData["scene"],
	})
	if err != nil {
		err = fmt.Errorf("添加用户数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
		return "", err
	}
	customerNo = types.GetString(noData)
	return customerNo, nil
}

func (d *DbWechatUser) updateUser(userInfo *oauth2.UserInfo, extData types.XMap) error {
	if userInfo == nil {
		return nil
	}
	if val, ok := extData["is_fans"]; !ok || val == "" {
		extData["is_fans"] = "1"
	}
	db := d.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.UpdateUserInfo, map[string]interface{}{
		"openid":    userInfo.OpenId,
		"nick_name": userInfo.Nickname,     // | varchar2(32)  |         | 是    | IS        | 昵称                             	|
		"head_url":  userInfo.HeadImageURL, // | varchar2(128) |         | 是    | IS        | 头像地址
		"is_fans":   extData["is_fans"],
	})
	if err != nil {
		err = fmt.Errorf("更新用户数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
		return err
	}
	return nil

}

//SaveUser 网页授权获取的用户信息
func (d *DbWechatUser) SaveUser(userInfo *oauth2.UserInfo, extData types.XMap) (customerNo string, err error) {
	db := d.c.GetRegularDB()

	//检查用户是否存在
	rows, q, a, err := db.Query(sql.QueryUserInfo, map[string]interface{}{
		"openid": userInfo.OpenId,
	})

	if err != nil {
		err = fmt.Errorf("查询用户信息错误(err:%v),sql:%s,输入参数:%v,data:%v", err, q, a, rows)
		return "", err
	}

	if !rows.IsEmpty() {
		err = d.updateUser(userInfo, extData)
		if err != nil {
			return "", err
		}
		return rows.Get(0).GetString("customer_no"), nil
	}

	customerNo, err = d.addUser(userInfo, extData)
	if err != nil {
		return "", err
	}
	return customerNo, err
}
