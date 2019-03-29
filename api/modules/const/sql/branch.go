package sql

// QueryAgentBranchInfoCount 查询代理商门店列表条数
const QueryAgentBranchInfoCount = `select count(1)
								  from 
								  	ag_branch_info a
								  where 
										a.agent_id =@agent_id 
										&branch_id
										&status
										`

// QueryAgentBranchInfo 查询代理商门店列表信息
const QueryAgentBranchInfo = `select TAB1.*
						from (select L.*
							from (select rownum LINENUM, R.*
								from (select
										a.branch_id,
										a.agent_id,
										a.branch_name,
										a.contact_name,
										a.contact_tel,
										a.province,
										a.city,
										a.district,
										a.address,
										a.status,
										b.chipcard_count,
										b.recharge_count,
										to_char(b.cumulative_standard,'fm99999990.00') cumulative_standard,
										to_char(a.create_time, 'yyyy-MM-dd hh24:mi:ss') as createdtime
									  from ag_branch_info a
										left join rpt_up_channel_branch_census b on a.branch_id = b.branch_id
									  where 
									  	a.agent_id =@agent_id and
										a.branch_id =nvl(@branch_id,a.branch_id)
										&status
									  order by create_time) R
								where rownum <= @pi * @ps) L
							where L.LINENUM > @ps * (@pi - 1)) TAB1`

// QueryChannelBranchInfoCount 查询渠道门店列表条数
const QueryChannelBranchInfoCount = `select count(1)
								  from 
									up_channel_branch_info a
								  where 
								 		a.up_channel_id in (select column_value from table(f_strsplit(@channel_ids, ',')))
										&branch_id
										&status
										`

// QueryChannelBranchInfo 查询渠道门店列表信息
const QueryChannelBranchInfo = `select TAB1.*
						from (select L.*
							from (select rownum LINENUM, R.*
								from (select
										a.branch_id,
										a.up_channel_id,
										a.branch_name,
										a.contact_name,
										a.contact_tel,
										a.province,
										a.city,
										a.district,
										a.address,
										a.status,
										b.chipcard_count,
										b.recharge_count,
										to_char(b.cumulative_standard,'fm99999990.00') cumulative_standard,
										to_char(a.create_time, 'yyyy-MM-dd hh24:mi:ss') as createdtime
									  from up_channel_branch_info a
										left join rpt_up_channel_branch_census b on a.branch_id = b.branch_id
									  where 
										a.up_channel_id in (select column_value from table(f_strsplit(@channel_ids, ','))) and
										a.branch_id =nvl(@branch_id,a.branch_id)
										&status
									  order by create_time) R
								where rownum <= @pi * @ps) L
							where L.LINENUM > @ps * (@pi - 1)) TAB1`

// EditChannelStatusInfo 编辑渠道门店状态
const EditChannelStatusInfo = `update up_channel_branch_info a set a.status = @status where a.branch_id =@branch_id`

// EditAgentStatusInfo 编辑代理门店状态
const EditAgentStatusInfo = `update ag_branch_info a set a.status = @status where a.branch_id =@branch_id`

// AddAgentBranchInfo 添加代理门店信息
const AddAgentBranchInfo = `insert into ag_branch_info
(
	branch_id,
	branch_name,
	agent_id,
	contact_name,
	contact_tel,
	province,
	city,
	district,
	address
)
values
(
	seq_branch_id.nextval,
	@branch_name,
	@channel_id,
	@contact_name,
	@contact_tel,
	@province,
	@city,
	@district,
	@address
)`

// IsExitAgentBranchNameInfo 当前代理商渠道下面是否存在相同门店
const IsExitAgentBranchNameInfo = `select count(1) count from ag_branch_info where branch_name=@branch_name and agent_id=@agent_id`

// IsExitsChannelBranchNameInfo 当前渠道商渠道下面是否存在相同门店
const IsExitsChannelBranchNameInfo = `select count(1) count from up_channel_branch_info where branch_name=@branch_name and up_channel_id=@up_channel_id`

// AddChannelBranchInfo 添加渠道门店信息
const AddChannelBranchInfo = `insert into up_channel_branch_info
(
	branch_id,
	branch_name,
	up_channel_id,
	contact_name,
	contact_tel,
	province,
	city,
	district,
	address
)
values
(
	seq_branch_id.nextval,
	@branch_name,
	@channel_id,
	@contact_name,
	@contact_tel,
	@province,
	@city,
	@district,
	@address
)`

// QueryProvince 查询省信息
const QueryProvince = `select a.canton_code value,a.chinese_name name from base_canton_info a where a.status =0 and  grade=1 order by a.sort_id`

// QueryCityByProvince 查询市信息
const QueryCityByProvince = `select a.canton_code value,a.chinese_name name
							from 
								base_canton_info a
							where  a.status =0 and a.grade=2 and a.parent_code = (
									select b.canton_code from base_canton_info b where b.status=0 and b.grade=1 and b.canton_code = @province
								)
								order by a.sort_id`

// QueryDistrictByCity 查询区信息
const QueryDistrictByCity = `select a.canton_code value,a.chinese_name name 
							from base_canton_info a
							where  
							a.status =0 and a.grade=3 and a.parent_code = (
								select b.canton_code from base_canton_info b where b.status=0 and b.grade=2 and b.canton_code = @city
							)
							order by a.sort_id`

// QueryAgentDetailInfo 查询代理门店详细信息
const QueryAgentDetailInfo = `select * from ag_branch_info where branch_id=@branch_id`

// QueryChannelDetailInfo 查询渠道门店详细信息
const QueryChannelDetailInfo = `select * from up_channel_branch_info where branch_id=@branch_id`

// UpdateAgentBranchInfo 更新代理门店信息
const UpdateAgentBranchInfo = `update ag_branch_info a set
								a.contact_name= @contact_name,
								a.contact_tel= @contact_tel,
								a.province= @province,
								a.city= @city,
								a.district= @district,
								a.address= @address
							 where a.branch_id =@branch_id`

// UpdateChannelBranchInfo 更新渠道门店信息
const UpdateChannelBranchInfo = `update up_channel_branch_info a set
								a.contact_name= @contact_name,
								a.contact_tel= @contact_tel,
								a.province= @province,
								a.city= @city,
								a.district= @district,
								a.address= @address
							 where a.branch_id =@branch_id`

const GetAllBranchesByUser = `
select
	t.branch_id,
	t.branch_name
from
	up_channel_branch_info t
where
	t.up_channel_id in (select column_value from table(f_strsplit(@up_channel_id)))
	and t.status=0
`

const GetBranchIDsByName = `
select
wm_concat(branch_id)
from
up_channel_branch_info a
where 
a.branch_name like '%'||@branch_name||'%'
`

const GetAllAgentBranchesByChannelID = `
select
	t.branch_id,
	t.branch_name
from
	ag_branch_info t
where
	t.agent_id = @channel_id
`
const GetAllChannelBranchesByChannelID = `
select
	t.branch_id,
	t.branch_name
from
	up_channel_branch_info t
where
	t.up_channel_id = @up_channel_id
`
