package product

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

// IProduct IProduct 接口
type IProduct interface {
	QueryPurchaseList(input map[string]interface{}) (list db.QueryRows, count int, err error)
	QueryPurcharseCardNoList(input map[string]interface{}) (list db.QueryRows, count int, err error)
	QueryPurcharseExportCardNoList(orderNo string) (list db.QueryRows, err error)

	QuerydCardList(input map[string]interface{}) (list db.QueryRows, count int, err error)
	QuerydRechargeList(input map[string]interface{}) (list db.QueryRows, count int, err error)
	QueryDownProductListByChannelNos(channelNos string) (list db.QueryRows, err error)
}

// Product Product 对象
type Product struct {
	c        component.IContainer
	dbHandle *DBProduct
}

// NewProduct构建 Product 对象
func NewProduct(c component.IContainer) *Product {
	return &Product{
		c:        c,
		dbHandle: NewDBProduct(c),
	}
}

func (d *Product) QueryPurchaseList(input map[string]interface{}) (list db.QueryRows, count int, err error) {
	return d.dbHandle.QueryPurchaseList(input)
}

func (d *Product) QueryPurcharseExportCardNoList(orderNo string) (list db.QueryRows, err error) {
	return d.dbHandle.QueryPurcharseExportCardNoList(orderNo)
}

func (d *Product) QueryPurcharseCardNoList(input map[string]interface{}) (list db.QueryRows, count int, err error) {
	return d.dbHandle.QueryPurcharseCardNoList(input)
}

func (d *Product) QueryDownProductListByChannelNos(channelNos string) (list db.QueryRows, err error) {
	return d.dbHandle.QueryDownProductListByChannelNos(channelNos)
}

func (d *Product) QuerydCardList(input map[string]interface{}) (list db.QueryRows, count int, err error) {
	return d.dbHandle.QuerydCardList(input)
}

func (d *Product) QuerydRechargeList(input map[string]interface{}) (list db.QueryRows, count int, err error) {
	return d.dbHandle.QuerydRechargeList(input)
}
