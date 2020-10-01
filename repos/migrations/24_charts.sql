CREATE VIEW tezos.blocks_delay AS
select blocks.timestamp, blocks.timestamp - prev.timestamp block_delay
from tezos.blocks
       inner join tezos.blocks prev on blocks.level = prev.level + 1;

CREATE VIEW tezos.delegations_view AS
select operations.*, acch.balance delegation_amount
from tezos.operations
       left join tezos.accounts_history acch
                 on operations.source = acch.account_id and operations.block_level = acch.block_level
where kind = 'delegation';

CREATE OR REPLACE VIEW tezos.block_priority_counter_view AS
select
       meta_cycle,
       timestamp,
       CASE WHEN priority = 0 THEN 1 ELSE 0 END  as zero_priority,
       CASE WHEN priority = 1 THEN 1 ELSE 0 END  as first_priority,
       CASE WHEN priority = 2 THEN 1 ELSE 0 END as second_priority,
       CASE WHEN priority >= 3 THEN 1 ELSE 0 END as third_priority
from tezos.blocks;

CREATE TABLE tezos.whale_accounts_periods
(
  day            timestamp without time zone,
  whale_accounts integer,
  PRIMARY KEY (day)
);

CREATE OR REPLACE FUNCTION insert_whale_stat(data integer) RETURNS integer AS
$$
BEGIN
  INSERT INTO tezos.whale_accounts_periods
  select date_trunc('day', to_timestamp(data)), count(1)
  from (select account_id, max(block_level) block_level
        from tezos.accounts_history
        where asof <= to_timestamp(data)
        group by account_id) r
         left join tezos.accounts_history ah on r.account_id = ah.account_id
    and r.block_level = ah.block_level
  where balance >= 500000000000;
  RETURN null;
END;
$$ LANGUAGE plpgsql;