package sql

const QueryAllUpChannelAccountAndIDByUser = `
select
wm_concat(channel_no) channel_no
from
down_channel_info
where
(agent_id=@agent_id and #agent=decode(#agent,1,1,-1)) 
	or 
(channel_no=@channel_no and #channel=decode(#channel,1,1,-1))
`
