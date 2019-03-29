package task

import (
	"fmt"

	"gsms/GSMSMerchant/api/modules/const/sql"
	"gsms/GSMSMerchant/api/modules/const/task"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//IDbTask 任务数据库操作接口
type IDbTask interface {
	Get() (db.QueryRows, error) //获取后补任务
	SetDealt(TID int) error     //任务设为已处理
}

//DbTask 任务数据库操作对象
type DbTask struct {
	c component.IContainer
}

//NewDbTask 任务数据库操作对象实例化
func NewDbTask(c component.IContainer) *DbTask {
	return &DbTask{c: c}
}

//Get 获取后补任务
func (t *DbTask) Get() (db.QueryRows, error) {
	db := t.c.GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("DB事务开启失败")
	}

	//1.获取新任务批次ID
	bid, q, a, err := dbTrans.Scalar(sql.GetNewTaskBatchID, map[string]interface{}{})
	if err != nil {
		dbTrans.Rollback()
		return nil, fmt.Errorf("获取新任务批次ID失败(err:%v),sql:(%s),输入参数:(%v)", err, q, a)
	}
	BID := types.GetInt64(bid)

	//2.设置任务状态及批次号
	row, q, a, err := dbTrans.Execute(sql.SetTaskDealing, map[string]interface{}{
		"batch_id": BID,
		"status":   task.TaskDealing,
		"exstatus": task.TaskStandingBy,
	})
	if err != nil {
		dbTrans.Rollback()
		return nil, fmt.Errorf("设置任务状态及批次号失败(err:%v),sql:(%s),输入参数:(%v),数据库受影响行数:(%d)", err, q, a, row)
	}
	if row == 0 {
		dbTrans.Rollback()
		return nil, nil
	}

	//3.获取该批次任务数据
	data, q, a, err := dbTrans.Query(sql.GetTaskByBatchID, map[string]interface{}{
		"batch_id": BID,
	})
	if err != nil {
		dbTrans.Rollback()
		return nil, fmt.Errorf("获取该批次任务数据失败(err:%v),sql:(%s),输入参数:(%v)", err, q, a)
	}
	if data.Len() <= 0 {
		dbTrans.Rollback()
		return nil, nil
	}

	dbTrans.Commit()
	return data, nil
}

//SetDealt 任务设为已处理
func (t *DbTask) SetDealt(TID int) error {
	db := t.c.GetRegularDB()

	row, q, a, err := db.Execute(sql.SetTaskDealt, map[string]interface{}{
		"id":       TID,
		"status":   task.TaskDealt,
		"exstatus": task.TaskDealing,
	})
	if err != nil || row == 0 {
		return fmt.Errorf("设置任务状态及批次号失败(err:%v),sql:(%s),输入参数:(%v),数据库受影响行数:(%d)", err, q, a, row)
	}

	return nil
}
