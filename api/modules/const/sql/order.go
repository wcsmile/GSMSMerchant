package sql

// QueryOrderInfoCount QueryOrderInfoCount
const QueryOrderInfoCount = `
	select count(*)
					from gr_order_main a
						 left join gr_order_main_ext b on b.order_no = a.order_no
					where 
					b.order_type ='1' and
					a.down_channel_no in (select column_value from table(f_strsplit(@channelnos, ','))) and
					a.order_no = nvl(@order_no,a.order_no) and 
					a.payment_status = nvl(@payment_status,a.payment_status)
`

// QueryOrderInfo QueryOrderInfo
const QueryOrderInfo = `
	select TAB1.* 
	from (select L.* 
		from (select rownum as rn,R.*
			from (select 
					a.order_no,
					b.order_source,
					a.recharge_account_id,
					a.business_type,
					a.total_face,
					a.product_face,
					a.payment_status,
					a.recharge_status,
					a.down_channel_no,
					'渠道商' branch_type
				from gr_order_main a
					 left join gr_order_main_ext b on b.order_no = a.order_no
				where 
					  b.order_type ='1' and
					  a.down_channel_no in (select column_value from table(f_strsplit(@channelnos, ','))) and
					  a.order_no like '%'||@order_no||'%' and 
					  a.payment_status = nvl(@payment_status,a.payment_status)
				order by a.order_no desc) R 
			where rownum <= @pi * @ps) L 
	where L.rn > (@pi - 1) * @ps) TAB1
`

// QueryAccountBychannel 通过当前账户的渠道ID获取渠道账户
const QueryAccountBychannel = `select 
			a.channel_no ,a.channel_name
			from 
				down_channel_info a 
			where 
				a.channel_no = nvl(@channel_no,a.channel_no) and
				a.channel_no = @ref_branch_id 
`

// QueryAccountByAgent 通过当前账户的渠道ID获取渠道账户
const QueryAccountByAgent = `select 
			a.channel_no ,a.channel_name
			from 
				down_channel_info a 
			where 
				a.channel_no = nvl(@channel_no,a.channel_no) and
				a.channel_no in (select b.channel_no from down_channel_info b where b.agent_id = @ref_branch_id)
`

// QueryChannel 查询订单来源(渠道ID)
const QueryChannel = `select 
				a.channel_no ,a.channel_name 
			from 
				down_channel_info a 
				`
