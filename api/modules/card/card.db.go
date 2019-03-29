package card

import (
	"fmt"
	"gsms/GSMSMerchant/api/modules/const/sql"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

// DBCard DBCard 对象
type DBCard struct {
	c component.IContainer
}

// NewDBCard 构建 DBCard 对象
func NewDBCard(c component.IContainer) *DBCard {
	return &DBCard{
		c: c,
	}
}

//@todo
func (d *DBCard) QueryCardInfoByCardNo(cardNOs string) (list db.QueryRows, err error) {
	dbConn := d.c.GetRegularDB("gsms_common_v2")
	list, q, a, err := dbConn.Query(sql.QueryCardListByCardNo, map[string]interface{}{
		"card_no": cardNOs,
	})
	if err != nil {
		return nil, fmt.Errorf("根据卡号获取卡信息列表失败:err:%+v sql:%s params:%+v", err, q, a)
	}
	return list, nil
}
