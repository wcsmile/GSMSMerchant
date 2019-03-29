package main

import (
	mem "gsms/GSMSMerchant/api/modules/member"

	"github.com/micro-plat/hydra/context"
)

//handling 请求预处理(每个请求执行前执行)，用于检查用户权限，缓存用户状态，服务分发等
func (r *assistantapi) handling() {
	//每个请求执行前执行
	r.Handling(func(ctx *context.Context) (rt interface{}) {
		ctx.Log.Info("----------handing-----------:", ctx.Service)

		//是否配置jwt
		jwt, err := ctx.Request.GetJWTConfig() //获取jwt配置

		if err != nil {
			return err
		}
		if jwt.IsExcluded(ctx.Service) {
			return nil
		}

		// 缓存用户信息
		var m mem.LoginAdminState
		if err = ctx.Request.GetJWT(&m); err != nil {
			return context.NewError(context.ERR_FORBIDDEN, err)
		}
		if err = mem.SaveMember(ctx, &m); err != nil {
			return err
		}

		return nil
	})
}
