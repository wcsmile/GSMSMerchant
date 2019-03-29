package task

import (
	"gsms/GSMSMerchant/api/modules/util"

	"github.com/micro-plat/hydra/component"
)

//ITask 任务接口
type ITask interface {
	TaskBackup() error //任务后补处理
}

//Task 任务对象
type Task struct {
	c  component.IContainer
	db IDbTask
}

//NewTask 任务对象实体化
func NewTask(c component.IContainer) *Task {
	return &Task{
		c:  c,
		db: NewDbTask(c),
	}
}

//TaskBackup 任务后补处理
func (t *Task) TaskBackup() error {

	//查询任务信息
	rows, err := t.db.Get()
	if err != nil || rows.Len() == 0 {
		return err
	}

	//循环发送消息
	for _, row := range rows {
		queItem := map[string]string{}
		queItem[row.GetString("group_code")] = row.GetString("content")
		util.SendMessageQueue(queItem)
		if er := t.db.SetDealt(row.GetInt("id")); er != nil {
			err = er
		}
	}
	return err
}
