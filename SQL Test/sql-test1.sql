select
	p.*
from
	t_policy p
join
	t_client c on p.client_number = c.client_number
where p.policy_submit_date > DATE '2018-01-15'
  and extract(month from c.birth_date) = 9;