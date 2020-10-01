CREATE VIEW tezos.blocks_counter_view AS
SELECT baker, count(1) as blocks
FROM tezos.blocks
GROUP BY baker;

CREATE or replace view tezos.baker_endorsement_view as
SELECT bakers.pkh  AS baker,
       bakers.staking_balance,
       en.count       AS endorsements,
       en.block_level as first_endorsement_level,
       bakers.balance,
       bakers.frozen_balance
FROM (tezos.bakers
       JOIN (SELECT endorsements_view.baker,
                    sum(endorsements_view.count) AS count,
                    min(block_level)             as block_level
             FROM tezos.endorsements_view
             GROUP BY endorsements_view.baker) en ON (((bakers.pkh)::text = (en.baker)::text)))
WHERE (bakers.deactivated = false);

create or replace view tezos.blocks_counter_view as
SELECT blocks.baker,
       count(1)   AS blocks,
       min(level) as first_baking_block
FROM tezos.blocks
GROUP BY blocks.baker;


create or replace view tezos.baker_delegations_view as
select delegate_value as baker, count(1) as active_delegations
from tezos.accounts
where delegate_value is not null
  and account_id != delegate_value
  and balance > 0
group by delegate_value;

create materialized view tezos.baker_view as
SELECT w.*, COALESCE(bdv.active_delegations, 0) as active_delegations
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
       left join tezos.baker_delegations_view as bdv on account_id = bdv.baker;

CREATE UNIQUE INDEX unique_index ON tezos.baker_view (account_id);

