// +build prod

package main

//install 用于配置生产环境参数，这些参数使用以#开头的变量命名，当执行 install 命令时会引导安装人员进行参数设置
func (s *assistantapi) install() {
	s.IsDebug = false

	s.Conf.SetInput(`#db_connection_string`, `数据库连接串`, `username/password@host`)
	s.Conf.API.SetMainConf(`{"address":":9092"}`)

	s.Conf.API.SetSubConf("app", `{
		"channel_id":"wechat"
	}`)

	s.Conf.API.SetSubConf("header", `
				{
					"Access-Control-Allow-Origin": "*",
					"Access-Control-Allow-Methods": "GET,POST,PUT,DELETE,PATCH,OPTIONS",
					"Access-Control-Allow-Headers": "__jwt__",
					"Access-Control-Allow-Credentials": "true",
					"Access-Control-Expose-Headers": "__jwt__"
				}
			`)

	s.Conf.API.SetSubConf("auth", `
		{
			"jwt": {
				"exclude": [
					"/industryapi/dictionary/**",	
					"/industryapi/wechat/**",
					"/industryapi/notify/**"
				],
				"expireAt": 36000,
				"mode": "HS512",
				"name": "__jwt__",
				"secret": "ef1a8839cb511780903ff6d5d79cf8f8",
				"domain": "offical.cdqykj.cn"
			}
		}
		`)

	s.Conf.Plat.SetVarConf("db", "db", `{			
			"provider":"ora",
			"connString":"#db_connection_string",
			"maxOpen":200,
			"maxIdle":10,
			"lifeTime":600		
	}`)

	s.Conf.Plat.SetVarConf("cache", "cache", `
		{
			"proto":"redis",
			"addrs":[
					#redis_server
			],
			"db":1,
			"dial_timeout":10,
			"read_timeout":10,
			"write_timeout":10,
			"pool_size":100
	}
		`)

	s.Conf.Plat.SetVarConf("queue", "queue", `
	{
		"proto":"redis",
		"addrs":[
				#redis_server
		],
		"db":1,
		"dial_timeout":10,
		"read_timeout":10,
		"write_timeout":10,
		"pool_size":100
}
	`)
}
