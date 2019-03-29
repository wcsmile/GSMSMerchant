package order

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

// IOrder 方法
type IOrder interface {
	Query(inputData *QueryOrderInput) (data db.QueryRows, count int, err error)
	QueryChannel() (data db.QueryRows, err error)
}

// Order 对象
type Order struct {
	c  component.IContainer
	db IDbOrder
}

// NewOrder 实例化
func NewOrder(c component.IContainer) *Order {
	return &Order{
		c:  c,
		db: NewDbOrder(c),
	}
}

// Query 查询订单信息
func (p *Order) Query(inputData *QueryOrderInput) (data db.QueryRows, count int, err error) {
	return p.db.Query(inputData)
}

// QueryChannel 查询订单来源
func (p *Order) QueryChannel() (data db.QueryRows, err error) {
	return p.db.QueryChannel()
}
