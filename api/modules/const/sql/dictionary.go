package sql

const QueryAllDictionary = `
select t.id, t.name, t.value, t.type, t.sort_id
  from base_dictionary_info t
 where t.status = 0
`
