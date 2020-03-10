CREATE TABLE tezos.public_bakers(
	baker_charges_transaction_fee boolean,
	baker_name varchar,
	baker_offchain_registry_url varchar,
	baker_pays_from_accounts varchar[],
	delegate varchar,
	min_delegation integer,
	min_payout  integer,
	open_for_delegation boolean,
	over_delegation_threshold integer,
	payout_delay integer,
	payment_config_mask varchar,
	payout_frequency integer,
	reporter_account varchar[],
	split integer,
	subtract_payouts_less_than_min boolean,
	subtract_rewards_from_uninvited_delegation boolean,
	last_update_id integer,
	baker_charges_transaction_fee boolean,
	is_hidden boolean default set false
	PRIMARY KEY (delegate)
)

create index ix_operations_endorsements_level
  on tezos.operations (level)
  where ((kind)::text = 'endorsement'::text);

create or replace view tezos.baker_endorsement_view as
SELECT delegates.pkh  AS baker,
       delegates.staking_balance,
       en.count       AS endorsements,
       en.block_level as first_endorsement_level,
       delegates.balance,
       delegates.frozen_balance
FROM (tezos.delegates
       JOIN (SELECT endorsements_view.baker,
                    sum(endorsements_view.count) AS count,
                    min(block_level)             as block_level
             FROM tezos.endorsements_view
             GROUP BY endorsements_view.baker) en ON (((delegates.pkh)::text = (en.baker)::text)))
WHERE (delegates.deactivated = false);

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

create index accounts_delegate_value_index
  on tezos.accounts (delegate_value)
  where delegate_value is not null and balance > 0;

drop materialized view tezos.baker_view;

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

create unique index unique_index
  on tezos.baker_view (account_id);

---
drop table tezos.baker_alias;
