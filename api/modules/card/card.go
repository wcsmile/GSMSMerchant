package card

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

// ICard ICard 接口
type ICard interface {
	QueryCardInfoByCardNo(cardNOs string) (list db.QueryRows, err error)
}

// Card Card 对象
type Card struct {
	c        component.IContainer
	dbHandle *DBCard
}

// NewCard构建 Card 对象
func NewCard(c component.IContainer) *Card {
	return &Card{
		c:        c,
		dbHandle: NewDBCard(c),
	}
}

func (d *Card) QueryCardInfoByCardNo(cardNOs string) (list db.QueryRows, err error) {
	return d.dbHandle.QueryCardInfoByCardNo(cardNOs)
}
