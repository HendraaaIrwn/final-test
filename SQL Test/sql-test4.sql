update
	t_policy
set
	policy_due_date = (
    date_trunc('month', policy_submit_date + interval '30 days')
    + interval '1 month - 1 day'
)::date;

select
    policy_number,
    policy_submit_date,
    policy_due_date
from
	t_policy
order by policy_number;