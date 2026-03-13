update
	t_agent a
set 
	basic_commission = sub.basic_commission
from (
    select
        agent_code,
        avg((commission / premium) * 100) as basic_commission
    from
		t_policy
    group by 
		agent_code
) sub
where a.agent_code = sub.agent_code;

select * 
from t_agent;