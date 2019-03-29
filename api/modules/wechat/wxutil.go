package wechat

import (
	"gsms/GSMSMerchant/api/modules/app"
	"gsms/GSMSMerchant/api/modules/channel"
	"gsms/GSMSMerchant/api/modules/common"
	"fmt"
	"sync"

	"github.com/micro-plat/wechat/mp"
)

var (
	utilContext    *mp.Context
	utilOncelocker sync.Once
)

func GetUtil() *mp.Context {

	utilOncelocker.Do(func() {
		wechatConf := app.GetWechatConf(common.Container)
		channelHandle := channel.NewChannel(common.Container)
		channelInfo, _ := channelHandle.Get(wechatConf.ChannelID)

		appID := channelInfo.GetString("appid")
		appSecret := channelInfo.GetString("app_secret")
		tokenServer := channelInfo.GetString("token_server")

		tk := mp.NewDefaultAccessTokenByURL(appID,
			appSecret,
			fmt.Sprintf("%s/%s/wechat/token/get", tokenServer, appID))
		utilContext = mp.NewContext(tk)
	})
	return utilContext
}
