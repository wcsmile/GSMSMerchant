package main

import (
	"fmt"
	"gsms/GSMSMerchant/api/modules/common"
	mem "gsms/GSMSMerchant/api/modules/member"
	"gsms/GSMSMerchant/api/servers/mgrserver/services/branch"
	"gsms/GSMSMerchant/api/servers/mgrserver/services/businesstype"
	"gsms/GSMSMerchant/api/servers/mgrserver/services/capitalrecord"
	"gsms/GSMSMerchant/api/servers/mgrserver/services/dictionary"
	"gsms/GSMSMerchant/api/servers/mgrserver/services/member"
	"gsms/GSMSMerchant/api/servers/mgrserver/services/order"
	"gsms/GSMSMerchant/api/servers/mgrserver/services/product"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/component"
)

//init 检查app，数据库，缓存，队列等配置是否正确，并注册服务
func (r *assistantapi) init() {
	r.Initializing(func(c component.IContainer) error {

		common.Container = c
		common.IsDebug = r.IsDebug

		//获取配置
		var conf mem.Conf
		if err := c.GetAppConf(&conf); err != nil {
			return err
		}
		if b, err := govalidator.ValidateStruct(&conf); !b {
			return fmt.Errorf("app 配置文件有误:%v", err)
		}
		mem.SaveConf(c, &conf)

		return nil
	})

	r.Micro("/dictionary/query", dictionary.NewDictionaryHandler, "*") //字典获取
	r.Micro("/member/login", member.NewLoginHandler, "*")              //登录系统
	r.Micro("/member/menu/get", member.NewMenuHandler, "*")            //获取菜单
	r.Micro("/sys/get", member.NewInfoHandler, "*")                    //获取系统信息
	r.Micro("/member/update/pwd", member.NewUpdateHandler, "*")        //修改登录密码

	//----------------------门店操作------------------------//
	r.Micro("/branch", branch.NewBranchManagerHandler, "*") //获取门店
	r.Micro("/order", order.NewOrderManagerHandler, "*")    //订单列表

	r.Micro("/product/purchase", product.NewProductPurchaseHandler, "*")   //商品采购
	r.Micro("/sys/businesstype", businesstype.NewBusinessTypeHandler, "*") //业务类型
	r.Micro("/product", product.NewProductHandler, "*")                    //业务类型

	r.Micro("/capitalrecord/list", capitalrecord.NewCapitalRecordHandler, "*") //资金流水

}
