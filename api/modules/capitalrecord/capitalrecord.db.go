package capitalrecord

import (
	"fmt"
	"gsms/GSMSMerchant/api/modules/const/sql"
	"gsms/GSMSMerchant/api/modules/const/usertype"
	"gsms/GSMSMerchant/api/modules/member"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

// DBCapitalrecord DBCapitalrecord 对象
type DBCapitalrecord struct {
	c component.IContainer
}

// NewDBCapitalrecord 构建 DBCapitalrecord 对象
func NewDBCapitalrecord(c component.IContainer) *DBCapitalrecord {
	return &DBCapitalrecord{
		c: c,
	}
}

//QueryRecordList 获取佣金流水.
func (d *DBCapitalrecord) QueryRecordList(userInfo *member.LoginAdminState, input map[string]interface{}) (list db.QueryRows, count int, err error) {
	dbConn := d.c.GetRegularDB("gsms_delivery_v2")

	if strings.EqualFold(userInfo.AccType, usertype.Channel) { //渠道
		input["channel_no"] = userInfo.RefAccID
		return QueryChannelList(dbConn, input)
	}
	if strings.EqualFold(userInfo.AccType, usertype.Agent) { //代理
		input["agent_id"] = userInfo.RefAccID
		return QueryAgentList(dbConn, input)
	}

	// data, q, a, err := dbConn.Scalar(sql.QueryRecordListCount, input)
	// if err != nil {
	// 	return nil, 0, fmt.Errorf("查询资金（佣金）流水数据失败:err:%+v sql:%s params:%+v", err, q, a)
	// }
	// fmt.Println("data count:", data)
	// if types.GetInt(data, -1) < 1 {
	// 	return nil, 0, nil
	// }

	// list, q, a, err = dbConn.Query(sql.QueryRecordList, input)
	// if err != nil {
	// 	return nil, 0, fmt.Errorf("查询资金（佣金）流水数据失败:err:%+v sql:%s params:%+v", err, q, a)
	// }
	// return list, types.GetInt(data, -1), nil
	return
}

func QueryAgentList(dbConn db.IDB, input map[string]interface{}) (list db.QueryRows, count int, err error) {

	data, q, a, err := dbConn.Scalar(sql.QueryRecordListCount, input)
	if err != nil {
		return nil, 0, fmt.Errorf("QueryAgentList 查询资金（佣金）流水数据失败:err:%+v sql:%s params:%+v", err, q, a)
	}
	fmt.Println("data count:", data)
	if types.GetInt(data, -1) < 1 {
		return nil, 0, nil
	}

	list, q, a, err = dbConn.Query(sql.QueryRecordList, input)
	if err != nil {
		return nil, 0, fmt.Errorf("QueryAgentList 查询资金（佣金）流水数据失败:err:%+v sql:%s params:%+v", err, q, a)
	}
	return list, types.GetInt(data, -1), nil
}

func QueryChannelList(dbConn db.IDB, input map[string]interface{}) (list db.QueryRows, count int, err error) {
	data, q, a, err := dbConn.Scalar(sql.QueryChannelListCount, input)
	if err != nil {
		return nil, 0, fmt.Errorf("QueryChannelList 查询渠道资金流水数据失败:err:%+v sql:%s params:%+v", err, q, a)
	}
	fmt.Println("data count:", data, q, a)
	if types.GetInt(data, -1) < 1 {
		return nil, 0, nil
	}

	list, q, a, err = dbConn.Query(sql.QueryChannelList, input)
	if err != nil {
		return nil, 0, fmt.Errorf("QueryChannelList 查询渠道资金流水数据失败:err:%+v sql:%s params:%+v", err, q, a)
	}
	return list, types.GetInt(data, -1), nil
}
