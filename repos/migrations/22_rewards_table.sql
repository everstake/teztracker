CREATE OR REPLACE FUNCTION tezos.baking_rewards()
 RETURNS trigger LANGUAGE plpgsql
AS $$
BEGIN
insert into tezos.baking_rewards
select delegate_value, NEW.snp_cycle, count(1), sum(balance)
from (select account_id, max(block_level) block_level
      from tezos.accounts_history
             left join tezos.snapshots on NEW.snp_cycle = snp_cycle
      where cycle <= (NEW.snp_cycle - 7)
        and block_level <= snp_block_level
      group by account_id) s
       left join tezos.accounts_history ah on s.account_id = ah.account_id and s.block_level = ah.block_level
where delegate_value is not null
  and balance > 0
group by delegate_value;
RETURN NEW;
END $$;

CREATE TRIGGER baker_rewards_insert
  AFTER INSERT
  ON tezos.snapshots
  FOR EACH ROW
EXECUTE PROCEDURE tezos.baking_rewards();

CREATE OR REPLACE VIEW tezos.frozen_endorsement_rewards as
select delegate,sum(reward) rewards, sum(count) count
from tezos.baker_cycle_endorsements_view  where cycle >= (select meta_cycle from tezos.blocks order by level desc limit 1) - 5 group by delegate;

 CREATE OR REPLACE VIEW tezos.frozen_baking_rewards as
select delegate,sum(reward) rewards, sum(count) count
from tezos.baker_cycle_bakings_view where cycle >= (select meta_cycle from tezos.blocks order by level desc limit 1) - 5 group by delegate;

--Add frozen rewards and count
drop materialized view tezos.baker_view;

create materialized view tezos.baker_view as
SELECT w.*, COALESCE(bdv.active_delegations, 0) as active_delegations, fer.rewards as frozen_endorsement_rewards, fer.count endorsement_count, fbr.rewards frozen_baking_rewards, fbr.count baking_count
FROM (
       SELECT CASE
                WHEN (r.bcvbaker IS NOT NULL) THEN r.bcvbaker
                ELSE r.bevbaker
                END AS account_id,
              r.staking_balance,
              TRUNC(staking_balance/8000/1000000,0) as rolls,
              r.frozen_balance,
              r.endorsements,
              r.blocks,
              b.timestamp as baking_since,
              r.balance
       FROM (SELECT bcv.baker                                          AS bcvbaker,
                    bev.baker                                          AS bevbaker,
                    COALESCE(bev.staking_balance, (0)::numeric)        AS staking_balance,
                    COALESCE(bev.frozen_balance,  (0)::numeric)        AS frozen_balance,
                    COALESCE(bev.balance, (0)::numeric)                AS balance,
                    COALESCE(bev.endorsements, (0)::bigint)            AS endorsements,
                    COALESCE(bcv.blocks, (0)::bigint)                  AS blocks,
                    LEAST(first_endorsement_level, first_baking_block) as first_block
             FROM (tezos.baker_endorsement_view bev
                    FULL JOIN tezos.blocks_counter_view bcv ON (((bev.baker)::text = (bcv.baker)::text)))
            ) r
       LEFT JOIN tezos.blocks as b on r.first_block = b.level
       WHERE ((r.bcvbaker IS NOT NULL) OR (r.bevbaker IS NOT NULL))
     ) as w
       left join tezos.baker_delegations_view as bdv on account_id = bdv.baker
        left join tezos.frozen_baking_rewards as fbr on account_id = fbr.delegate
left join tezos.frozen_endorsement_rewards as fer on account_id = fer.delegate;

create unique index unique_index
  on tezos.baker_view (account_id);


