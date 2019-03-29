package upchannel

import (
	"gsms/GSMSMerchant/api/modules/member"

	"github.com/micro-plat/hydra/component"
)

// IUpChannel IUpChannel 接口
type IUpChannel interface {
	QueryAllUpChannelAccountAndIDByUser(user *member.LoginAdminState) (channelNos string, err error)
}

// UpChannel UpChannel 对象
type UpChannel struct {
	c        component.IContainer
	dbHandle *DBUpChannel
}

// NewUpChannel构建 UpChannel 对象
func NewUpChannel(c component.IContainer) *UpChannel {
	return &UpChannel{
		c:        c,
		dbHandle: NewDBUpChannel(c),
	}
}

func (d *UpChannel) QueryAllUpChannelAccountAndIDByUser(user *member.LoginAdminState) (channelNos string, err error) {
	return d.dbHandle.QueryAllUpChannelAccountAndIDByUser(user)
}
