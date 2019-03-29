package businesstype

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

// IBusinessType IBusinessType 接口
type IBusinessType interface {
	QueryAllBusnessType() (list db.QueryRows, err error)
}

// BusinessType BusinessType 对象
type BusinessType struct {
	c        component.IContainer
	dbHandle *DBBusinessType
}

// NewBusinessType构建 BusinessType 对象
func NewBusinessType(c component.IContainer) *BusinessType {
	return &BusinessType{
		c:        c,
		dbHandle: NewDBBusinessType(c),
	}
}

func (d *BusinessType) QueryAllBusnessType() (list db.QueryRows, err error) {
	return d.dbHandle.QueryAllBusnessType()
}
