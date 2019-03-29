package sql

const QueryCardListByCardNo = `
select
	t.card_no,
	t.status,
	t.has_first_recharge
from
oil_card t 
where
t.card_no in (select column_value from table(strsplit(@card_no)))
`
