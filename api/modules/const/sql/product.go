package sql

const QueryProductPurchaseListCount = `
select
count(1)
from
gr_order_main t 
inner join gr_order_main_ext e on (
	t.order_no=e.order_no
and e.order_type=4
and e.need_logistics=nvl(@need_logistics,e.need_logistics)
)
where
t.create_time >= to_date(@start_time,'yyyy-mm-dd hh24:mi:ss')
and t.create_time <= to_date(@end_time,'yyyy-mm-dd hh24:mi:ss')
and t.down_channel_no in
(select column_value from table(f_strsplit(@down_channel_no)))
`

const QueryProductPurchaseList = `
select TAB1.*
	from (select L.*
		from (select rownum LINENUM, R.*
			from (
				select
					t.order_no,
					t.product_num,
					t.down_product_no,
					t.business_type,
					t.product_standard,
					t.product_face,
					t.order_status,
					e.need_logistics,
					e.order_source,
					to_char(ac.finish_time, 'yyyy-MM-dd hh24:mi:ss') delivery_time,
					ac.logistics_company,
					ac.logistics_order_no
				from
					gr_order_main t 
				inner join gr_order_main_ext e on (
						t.order_no=e.order_no
						and e.order_type=4
						and e.need_logistics=nvl(@need_logistics,e.need_logistics)
					)
				left join gr_order_main_ext_accopen ac on ac.order_no=t.order_no
				where
					t.create_time >= to_date(@start_time,'yyyy-mm-dd hh24:mi:ss')
					and t.create_time <= to_date(@end_time,'yyyy-mm-dd hh24:mi:ss')
					and t.down_channel_no in (select column_value from table(f_strsplit(@down_channel_no)))
				order by t.create_time desc
		) R
	where rownum <= @pi * @ps) L
where L.LINENUM > @ps * (@pi - 1)) TAB1
`

const QueryProductPurchaseListCountWithCardNo = `
select
count(1)
from
gr_order_main t 
inner join gr_order_main_ext e on (
	t.order_no=e.order_no
and e.order_type=4
and e.need_logistics=nvl(@need_logistics,e.need_logistics)
)
inner join gr_delivery_main d on (t.order_no=d.order_no and d.recharge_account_id=@card_no and and d.delivery_status=0)
where
t.create_time >= to_date(@start_time,'yyyy-mm-dd hh24:mi:ss')
and t.create_time <= to_date(@end_time,'yyyy-mm-dd hh24:mi:ss')
and t.down_channel_no in
(select column_value from table(f_strsplit(@down_channel_no)))
`

const QueryProductPurchaseListWithCardNo = `
select TAB1.*
	from (select L.*
		from (select rownum LINENUM, R.*
			from (
				select
					t.order_no,
					t.product_num,
					t.down_product_no,
					t.business_type,
					t.product_standard,
					t.product_face,
					t.order_status,
					e.need_logistics,
					e.order_source,
					to_char(ac.finish_time, 'yyyy-MM-dd hh24:mi:ss') delivery_time,
					ac.logistics_company,
					ac.logistics_order_no
				from
					gr_order_main t 
				inner join gr_order_main_ext e on (
						t.order_no=e.order_no
						and e.order_type=4
						and e.need_logistics=nvl(@need_logistics,e.need_logistics)
					)
				inner join gr_delivery_main d on (t.order_no=d.order_no and d.recharge_account_id=@card_no and d.delivery_status=0)
				left join gr_order_main_ext_accopen ac on ac.order_no=t.order_no
				where
					t.create_time >= to_date(@start_time,'yyyy-mm-dd hh24:mi:ss')
					and t.create_time <= to_date(@end_time,'yyyy-mm-dd hh24:mi:ss')
					and t.down_channel_no in (select column_value from table(f_strsplit(@down_channel_no)))
				order by t.create_time desc
		) R
	where rownum <= @pi * @ps) L
where L.LINENUM > @ps * (@pi - 1)) TAB1
`

const QueryDeliveryCardNoCountByOrderNo = `
select
count(1)
from
gr_delivery_main
where
order_no=@order_no
and delivery_status=0
`

const QueryDeliveryCardNoListByOrderNo = `
select TAB1.*
	from (select L.*
		from (select rownum LINENUM, R.*
			from (
				select
				recharge_account_id
				from
				gr_delivery_main
				where
				order_no=@order_no
				and delivery_status=0
) R
where rownum <= @pi * @ps) L
where L.LINENUM > @ps * (@pi - 1)) TAB1
`

const QueryAllDeliveryCardNoListByOrderNo = `
select
				recharge_account_id
				from
				gr_delivery_main
				where
				order_no=@order_no
				and delivery_status=0
`

const QueryDownChannelProductsByChannelNo = `
select
	product_no,
	down_channel_no,
	standard,
	face,
	deduct_discount
from
down_product_details
where
down_channel_no in (select column_value from table(strsplit(@down_channel_no)))
`

const QueryProductCardListCount = `
select
count(1)
from
gr_delivery_main t 
inner join gr_order_main m on (
	m.order_no=t.order_no
	and m.down_channel_no in (select column_value from table(f_strsplit(@down_channel_no)))
)
inner join gr_order_main_ext e on (
	t.order_no=e.order_no
and e.order_type=4
and e.order_source=nvl(@order_source,e.order_source)
)
where
t.create_time >= to_date(@start_time,'yyyy-mm-dd hh24:mi:ss')
and t.create_time <= to_date(@end_time,'yyyy-mm-dd hh24:mi:ss')
and t.recharge_account_id=nvl(@recharge_account_id,t.recharge_account_id)
and t.delivery_status<>90
and t.recharge_account_id<>'-'
`

const QueryPoductCardList = `
select TAB1.*
	from (select L.*
		from (select rownum LINENUM, R.*
			from (
				select
				t.recharge_account_id,
				to_char(t.create_time, 'yyyy-MM-dd hh24:mi:ss') create_time,
				m.business_type,
				m.down_product_no,
				e.order_source
				from
				gr_delivery_main t 
				inner join gr_order_main m on (
					m.order_no=t.order_no
					and m.down_channel_no in (select column_value from table(f_strsplit(@down_channel_no)))
				)
				inner join gr_order_main_ext e on (
					t.order_no=e.order_no
				and e.order_type=4
				and e.order_source=nvl(@order_source,e.order_source)
				)
				where
				t.create_time >= to_date(@start_time,'yyyy-mm-dd hh24:mi:ss')
				and t.create_time <= to_date(@end_time,'yyyy-mm-dd hh24:mi:ss')
				and t.recharge_account_id=nvl(@recharge_account_id,t.recharge_account_id)
				and t.delivery_status<>90
				and t.recharge_account_id<>'-'
				order by t.create_time desc
		) R
	where rownum <= @pi * @ps) L
where L.LINENUM > @ps * (@pi - 1)) TAB1
`

const QueryProductRechargeListCount = `
select
count(1)
from
down_product_details t 
where
t.down_channel_no in (select column_value from table(strsplit(@down_channel_no)))
`

const QueryProductRechargeList = `
select
t.product_no,
t.down_shelf_id,
t.business_type,
t.province_no,
decode(t.province_no,'*','全国',c1.chinese_name) province_name,
t.city_no,
decode(t.city_no,'*','全省',c2.chinese_name) city_name,
t.face,
to_char(t.deduct_discount, 'fm9999990.00') deduct_discount,
t.status
from
down_product_details t 
left join sys_canton_info c1 on(t.province_no=c1.canton_code and c1.grade=1) 
left join sys_canton_info c2 on (t.city_no=c2.canton_code and c2.grade=2)
where
t.down_channel_no in (select column_value from table(strsplit(@down_channel_no)))
`
