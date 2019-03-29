package product

import (
	"fmt"
	"gsms/GSMSMerchant/api/modules/const/sql"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

// DBProduct DBProduct 对象
type DBProduct struct {
	c component.IContainer
}

// NewDBProduct 构建 DBProduct 对象
func NewDBProduct(c component.IContainer) *DBProduct {
	return &DBProduct{
		c: c,
	}
}

//@todo
func (d *DBProduct) QueryPurchaseList(input map[string]interface{}) (list db.QueryRows, count int, err error) {
	dbConn := d.c.GetRegularDB("gsms_delivery_v2")
	if input["card_no"] == "" {
		data, q, a, err := dbConn.Scalar(sql.QueryProductPurchaseListCount, input)
		if err != nil {
			return nil, 0, fmt.Errorf("查询商品采购数量失败:err:%+v sql:%s params:%+v", err, q, a)
		}
		if types.GetInt(data, -1) < 1 {
			return nil, 0, nil
		}

		list, q, a, err = dbConn.Query(sql.QueryProductPurchaseList, input)
		if err != nil {
			return nil, 0, fmt.Errorf("查询商品采购列表失败:err:%+v sql:%s params:%+v", err, q, a)
		}
		return list, types.GetInt(data, -1), nil
	} else {
		data, q, a, err := dbConn.Scalar(sql.QueryProductPurchaseListCountWithCardNo, input)
		if err != nil {
			return nil, 0, fmt.Errorf("b查询商品采购数量失败:err:%+v sql:%s params:%+v", err, q, a)
		}
		if types.GetInt(data, -1) < 1 {
			return nil, 0, nil
		}

		list, q, a, err = dbConn.Query(sql.QueryProductPurchaseListWithCardNo, input)
		if err != nil {
			return nil, 0, fmt.Errorf("b查询商品采购列表失败:err:%+v sql:%s params:%+v", err, q, a)
		}
		return list, types.GetInt(data, -1), nil
	}
}

func (d *DBProduct) QueryPurcharseCardNoList(input map[string]interface{}) (list db.QueryRows, count int, err error) {
	dbConn := d.c.GetRegularDB("gsms_delivery_v2")
	data, q, a, err := dbConn.Scalar(sql.QueryDeliveryCardNoCountByOrderNo, input)
	if err != nil {
		return nil, 0, fmt.Errorf("查询发货记录卡号数量失败:err:%+v sql:%s params:%+v", err, q, a)
	}
	if types.GetInt(data, -1) < 1 {
		return nil, 0, nil
	}
	list, q, a, err = dbConn.Query(sql.QueryDeliveryCardNoListByOrderNo, input)
	if err != nil {
		return nil, 0, fmt.Errorf("查询发货记录卡号失败:err:%+v sql:%s params:%+v", err, q, a)
	}
	return list, types.GetInt(data, -1), nil
}

func (d *DBProduct) QueryPurcharseExportCardNoList(orderNo string) (list db.QueryRows, err error) {
	dbConn := d.c.GetRegularDB("gsms_delivery_v2")
	list, q, a, err := dbConn.Query(sql.QueryAllDeliveryCardNoListByOrderNo, map[string]interface{}{
		"order_no": orderNo,
	})
	if err != nil {
		return nil, fmt.Errorf("查询导出发货记录卡号失败:err:%+v sql:%s params:%+v", err, q, a)
	}
	return
}

func (d *DBProduct) QueryDownProductListByChannelNos(channelNos string) (list db.QueryRows, err error) {
	dbConn := d.c.GetRegularDB("gsms_common_v2")
	list, q, a, err := dbConn.Query(sql.QueryDownChannelProductsByChannelNo, map[string]interface{}{
		"down_channel_no": channelNos,
	})
	if err != nil {
		return nil, fmt.Errorf("查询下游渠道产品失败:err:%+v sql:%s params:%+v", err, q, a)
	}
	return
}

func (d *DBProduct) QuerydCardList(input map[string]interface{}) (list db.QueryRows, count int, err error) {
	dbConn := d.c.GetRegularDB("gsms_delivery_v2")
	data, q, a, err := dbConn.Scalar(sql.QueryProductCardListCount, input)
	if err != nil {
		return nil, 0, fmt.Errorf("查询芯片卡商品数量失败:err:%+v sql:%s params:%+v", err, q, a)
	}
	if types.GetInt(data, -1) < 1 {
		return nil, 0, nil
	}

	list, q, a, err = dbConn.Query(sql.QueryPoductCardList, input)
	if err != nil {
		return nil, 0, fmt.Errorf("查询芯片卡商品列表失败:err:%+v sql:%s params:%+v", err, q, a)
	}
	return list, types.GetInt(data, -1), nil
}

func (d *DBProduct) QuerydRechargeList(input map[string]interface{}) (list db.QueryRows, count int, err error) {
	dbConn := d.c.GetRegularDB("gsms_common_v2")
	data, q, a, err := dbConn.Scalar(sql.QueryProductRechargeListCount, input)
	if err != nil {
		return nil, 0, fmt.Errorf("查询充值商品数量失败:err:%+v sql:%s params:%+v", err, q, a)
	}
	if types.GetInt(data, -1) < 1 {
		return nil, 0, nil
	}

	list, q, a, err = dbConn.Query(sql.QueryProductRechargeList, input)
	if err != nil {
		return nil, 0, fmt.Errorf("查询充值商品列表失败:err:%+v sql:%s params:%+v", err, q, a)
	}
	return list, types.GetInt(data, -1), nil
}
