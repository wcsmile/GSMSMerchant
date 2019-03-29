package upchannel

import (
	"fmt"
	"gsms/GSMSMerchant/api/modules/const/sql"
	"gsms/GSMSMerchant/api/modules/const/usertype"
	"gsms/GSMSMerchant/api/modules/member"

	"github.com/micro-plat/hydra/component"
)

// DBUpChannel DBUpChannel 对象
type DBUpChannel struct {
	c component.IContainer
}

// NewDBUpChannel 构建 DBUpChannel 对象
func NewDBUpChannel(c component.IContainer) *DBUpChannel {
	return &DBUpChannel{
		c: c,
	}
}

//@todo
func (d *DBUpChannel) QueryAllUpChannelAccountAndIDByUser(user *member.LoginAdminState) (channelNos string, err error) {
	dbConn := d.c.GetRegularDB("gsms_common_v2")
	agent, channel := 0, 0
	if user.AccType == usertype.Agent { //代理
		agent = 1
	} else if user.AccType == usertype.Channel { //渠道
		channel = 1
	} else {
		err = fmt.Errorf("未知的账号类型")
		return
	}
	data, q, a, err := dbConn.Query(sql.QueryAllUpChannelAccountAndIDByUser, map[string]interface{}{
		"channel_no": user.RefAccID,
		"agent_id":   user.RefAccID,
		"agent":      agent,
		"channel":    channel,
	})
	if err != nil {
		return "", fmt.Errorf("查询账户上游账号信息失败:err:%+v sql:%s params:%+v", err, q, a)
	}

	if len(data) == 0 {
		err = fmt.Errorf("QueryAllUpChannelAccountAndIDByUser：当前渠道没有数据")
		return
	}

	channelNos = data.Get(0).GetString("channel_no")
	return
}
