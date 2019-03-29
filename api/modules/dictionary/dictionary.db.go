package dictionary

import (
	"fmt"
	"gsms/GSMSMerchant/api/modules/const/sql"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

// DBDictionary DBDictionary 对象
type DBDictionary struct {
	c component.IContainer
}

// NewDBDictionary 构建 DBDictionary 对象
func NewDBDictionary(c component.IContainer) *DBDictionary {
	return &DBDictionary{
		c: c,
	}
}

// QueryAll 加油卡申请
func (d *DBDictionary) QueryAll() (list []db.QueryRow, err error) {
	dbConn := d.c.GetRegularDB("industry_marketing")

	list, _, _, err = dbConn.Query(sql.QueryAllDictionary, nil)
	if err != nil {
		return nil, fmt.Errorf("查询枚举字典所有数据失败:err:%+v", err)
	}
	return list, nil
}
