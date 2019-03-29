package app

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/component"
)

//Conf 应用程序配置
type WeChatConf struct {
	ChannelID string `json:"channel_id" valid:"ascii,required"`
}

//Valid 验证配置参数是否合法
func (c WeChatConf) Valid() error {
	if b, err := govalidator.ValidateStruct(&c); !b {
		return fmt.Errorf("app 配置文件有误:%v", err)
	}
	return nil
}

//GetAppID 获取微信AppID
func (c *WeChatConf) GetChannelID() string {
	return c.ChannelID
}

//SaveConf 保存当前应用程序配置
func SaveWechatConf(c component.IContainer, m *WeChatConf) {
	c.Set("__wechatConf__", m)
}

//GetConf 获取当前应用程序配置
func GetWechatConf(c component.IContainer) *WeChatConf {
	appval := c.Get("__wechatConf__")
	return appval.(*WeChatConf)
}
