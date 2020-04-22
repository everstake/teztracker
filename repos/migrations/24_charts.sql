CREATE VIEW tezos.blocks_delay AS
select blocks.timestamp ,blocks.timestamp - prev.timestamp block_delay
from tezos.blocks
  inner join tezos.blocks prev on blocks.level = prev.level + 1;

CREATE VIEW tezos.delegations_view AS
select operations.*, acch.balance delegation_amount
from tezos.operations
  left join tezos.accounts_history acch on operations.source = acch.account_id and operations.block_level = acch.block_level
where kind = 'delegation';

CREATE VIEW tezos.block_priority_counter_view AS
select timestamp, CASE WHEN priority = 0 THEN 1 ELSE 0 END as zero_priority,CASE WHEN priority = 1 THEN 1 ELSE 0 END as first_priority, CASE WHEN priority >= 2 THEN 1 ELSE 0 END as second_priority
from tezos.blocks;

create index accounts_history_asof_index
	on tezos.accounts_history (asof);
