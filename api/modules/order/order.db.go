package order

import (
	"fmt"
	"gsms/GSMSMerchant/api/modules/const/sql"
	"gsms/GSMSMerchant/api/modules/const/usertype"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//IDbOrder 方法
type IDbOrder interface {
	Query(inputData *QueryOrderInput) (data db.QueryRows, count int, err error)
	QueryChannel() (data db.QueryRows, err error)
}

// DbOrder 对象
type DbOrder struct {
	c component.IContainer
}

// NewDbOrder 实例化
func NewDbOrder(c component.IContainer) *DbOrder {
	return &DbOrder{
		c: c,
	}
}

// Query 查询订单信息
func (d *DbOrder) Query(inputData *QueryOrderInput) (data db.QueryRows, count int, err error) {
	db1 := d.c.GetRegularDB("gsms_delivery_v2")
	db3 := d.c.GetRegularDB("gsms_common_v2")
	var channelDatas []types.XMap

	// 通过当前渠道账户(channel)的渠道ID获取渠道信息
	if strings.EqualFold(inputData.AccType, usertype.Channel) {
		rows, q, a, err := db3.Query(sql.QueryAccountBychannel, map[string]interface{}{
			"ref_branch_id": inputData.RefAccID,
			"channel_no":    inputData.ChannelNo,
		})
		if err != nil {
			return nil, 0, fmt.Errorf("通过当前渠道账户(channel)的渠道ID获取渠道信息发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
		}
		if len(rows) == 0 {
			return nil, 0, nil
		}
		channelDatas = rows
	}

	// 通过当代理商账户(agent)的渠道ID获取渠道信息
	if strings.EqualFold(inputData.AccType, usertype.Agent) {
		rows, q, a, err := db3.Query(sql.QueryAccountByAgent, map[string]interface{}{
			"ref_branch_id": inputData.RefAccID,
			"channel_no":    inputData.ChannelNo,
		})
		if err != nil {
			return nil, 0, fmt.Errorf("通过当代理商账户(agent)的渠道ID获取渠道信息发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
		}
		if len(rows) == 0 {
			return nil, 0, nil
		}
		channelDatas = rows
	}

	channelNoArry := []string{}
	for _, v := range channelDatas {
		channelNoArry = append(channelNoArry, v.GetString("channel_no"))
	}
	channelNos := strings.Join(channelNoArry, ",")

	params := map[string]interface{}{
		"order_no":       inputData.OrderNo,
		"pi":             inputData.Pi,
		"ps":             inputData.Ps,
		"payment_status": inputData.PaymentStatus,
		"channelnos":     channelNos,
		//"branch_type":    inputData.AccType,
	}
	total, sq, sa, err := db1.Scalar(sql.QueryOrderInfoCount, params)
	if err != nil || types.GetInt(total) < 0 {
		return nil, 0, fmt.Errorf("获取查询订单条数发生错误(err:%v),sql:%s,输入参数:%v", err, sq, sa)
	}

	datas, q1, a1, err1 := db1.Query(sql.QueryOrderInfo, params)
	if err != nil {
		return nil, 0, fmt.Errorf("获取查询订单信息发生错误(err:%v),sql:%s,输入参数:%v,", err1, q1, a1)
	}

	// 添加渠道名字
	for i := 0; i < len(datas); i++ {
		for j := 0; j < len(channelDatas); j++ {
			if strings.EqualFold(datas[i].GetString("down_channel_no"), channelDatas[j].GetString("channel_no")) {
				datas[i]["channel_name"] = channelDatas[j].GetString("channel_name")
			}
		}
	}

	return datas, types.GetInt(total), nil
}

// QueryChannel 查询订单来源(渠道ID)
func (d *DbOrder) QueryChannel() (data db.QueryRows, err error) {
	db := d.c.GetRegularDB("gsms_common_v2")

	// 查询订单来源(渠道ID)
	channelDatas, q, a, err := db.Query(sql.QueryChannel, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("查询订单来源(渠道ID)信息发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	if len(channelDatas) == 0 {
		return nil, nil
	}
	return channelDatas, nil
}
