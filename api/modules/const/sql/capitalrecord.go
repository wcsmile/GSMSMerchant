package sql

const QueryRecordListCount = `
select count(1)
  from gr_order_agent_fund_changes t
 where t.create_time >= to_date(@start_time, 'yy-mm-dd hh24:mi:ss')
   and t.create_time <= to_date(@end_time, 'yy-mm-dd hh24:mi:ss')
   and t.agent_id = @agent_id
   and t.business_type = nvl(@business_type,t.business_type)
   and t.id = nvl(@record_no, t.id)
`

const QueryRecordList = `
select TAB1.*
	from (select L.*
		from (select rownum LINENUM, R.*
			from (
				select 
				t.id,
				t.agent_id,
				t.order_no,
				t.down_channel_id,
				to_char(t.create_time, 'yyyy-mm-dd hh24:mi:ss') create_time,
				to_char(t.amount, 'fm99999999990.00') amount,
				to_char(t.balance, 'fm99999999990.00') balance,
				t.business_type,
				t.change_type
			from gr_order_agent_fund_changes t
			where t.create_time >= to_date(@start_time, 'yy-mm-dd hh24:mi:ss')
			and t.create_time <= to_date(@end_time, 'yy-mm-dd hh24:mi:ss')
			and t.agent_id = @agent_id
			and t.business_type = nvl(@business_type,t.business_type)
			and t.id = nvl(@record_no, t.id)
		) R
	where rownum <= @pi * @ps) L
where L.LINENUM > @ps * (@pi - 1)) TAB1
`

const QueryChannelListCount = `
select count(1)
  from view_fd_down_fund t
where t.create_time >= to_date(@start_time, 'yy-mm-dd hh24:mi:ss')
   and t.create_time <= to_date(@end_time, 'yy-mm-dd hh24:mi:ss')
   and t.channel_no = @channel_no
   and t.change_type = nvl(@business_type, t.change_type)
   and t.record_id = nvl(@record_no, t.record_id)
`

const QueryChannelList = `
select TAB1.*
	from (select L.*
		from (select rownum LINENUM, R.*
			from (
				select 
				t.record_id id,
				t.fd_order_id,
				t.channel_no,
				t.account_id,
				t.trade_order_no,
				t.ext_order_no,
				t.order_source,
				to_char(t.create_time, 'yyyy-mm-dd hh24:mi:ss') create_time,
				t.order_time,
				t.change_type channel_type,
				to_char(t.change_amount, 'fm99999999990.00') amount,
				to_char(t.balance, 'fm99999999990.00') balance,
				t.memo,
				t.service_fee
				from view_fd_down_fund t
			where t.create_time >= to_date(@start_time, 'yy-mm-dd hh24:mi:ss')
			and t.create_time <= to_date(@end_time, 'yy-mm-dd hh24:mi:ss')
			and t.channel_no = @channel_no
			and t.change_type = nvl(@business_type,t.change_type)
			and t.record_id = nvl(@record_no, t.record_id)
		) R
	where rownum <= @pi * @ps) L
where L.LINENUM > @ps * (@pi - 1)) TAB1
`
