package util

import (
	"coupon/customer/recv-coupon/modules/app"
	"fmt"

	"github.com/micro-plat/wechat/mp"
)

var (
	WxUtil *mp.Context
)

func New(conf *app.WeChatConf) {

	tk := mp.NewDefaultAccessTokenByURL(conf.AppID, conf.Secret,
		fmt.Sprintf("%s/%s/wechat/token/get", conf.GetWeChatServer(), conf.AppID))
	WxUtil = mp.NewContext(tk)

}
