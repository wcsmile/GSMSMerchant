package businesstype

import (
	"fmt"
	"gsms/GSMSMerchant/api/modules/const/sql"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

// DBBusinessType DBBusinessType 对象
type DBBusinessType struct {
	c component.IContainer
}

// NewDBBusinessType 构建 DBBusinessType 对象
func NewDBBusinessType(c component.IContainer) *DBBusinessType {
	return &DBBusinessType{
		c: c,
	}
}

//@todo
func (d *DBBusinessType) QueryAllBusnessType() (list db.QueryRows, err error) {
	dbConn := d.c.GetRegularDB("gsms_common_v2")
	list, q, a, err := dbConn.Query(sql.QueryAllSysBusinessType, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("查询系统业务类型失败:err:%+v sql:%s params:%+v", err, q, a)
	}
	return list, nil
}
