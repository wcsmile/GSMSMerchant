package dictionary

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

// IDictionary IDictionary 接口
type IDictionary interface {
	QueryAll() (list []db.QueryRow, err error)
}

// Dictionary Dictionary 对象
type Dictionary struct {
	c        component.IContainer
	dbHandle *DBDictionary
}

// NewDictionary 构建 Dictionary 对象
func NewDictionary(c component.IContainer) *Dictionary {
	return &Dictionary{
		c:        c,
		dbHandle: NewDBDictionary(c),
	}
}

// QueryAll 获取所有字典数据
func (d *Dictionary) QueryAll() (list []db.QueryRow, err error) {
	return d.dbHandle.QueryAll()
}
