package capitalrecord

import (
	"gsms/GSMSMerchant/api/modules/member"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

// ICapitalrecord ICapitalrecord 接口
type ICapitalrecord interface {
	QueryRecordList(userInfo *member.LoginAdminState, input map[string]interface{}) (list db.QueryRows, count int, err error)
}

// Capitalrecord Capitalrecord 对象
type Capitalrecord struct {
	c        component.IContainer
	dbHandle *DBCapitalrecord
}

// NewCapitalrecord构建 Capitalrecord 对象
func NewCapitalrecord(c component.IContainer) *Capitalrecord {
	return &Capitalrecord{
		c:        c,
		dbHandle: NewDBCapitalrecord(c),
	}
}

func (d *Capitalrecord) QueryRecordList(userInfo *member.LoginAdminState, input map[string]interface{}) (list db.QueryRows, count int, err error) {
	return d.dbHandle.QueryRecordList(userInfo, input)
}
