select
	p.*
from
	t_policy p
join
	t_agent a on p.agent_code = a.agent_code
where p.policy_status = 'INFORCE'
  and a.agent_office = 'JAKARTA'