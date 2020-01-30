create view tezos.blocks_counter_view as select blocks.baker,count(1) as blocks
from tezos.blocks group by baker;


CREATE VIEW tezos.baker_endorsement_view AS
select delegates.pkh as baker, staking_balance, en.count as endorsements from tezos.delegates
  inner join (select baker, SUM(count) as count from tezos.endorsements_view group by baker) as en ON delegates.pkh = en.baker
WHERE deactivated=false;

CREATE MATERIALIZED VIEW tezos.baker_view_new AS
select
       case WHEN bcvbaker IS NOT NULL THEN bcvbaker ELSE bevbaker END as account_id, staking_balance, endorsements, blocks
from (
        select bcv.baker as bcvbaker, bev.baker as bevbaker, COALESCE(staking_balance,0) as staking_balance, COALESCE(endorsements,0) as endorsements, COALESCE(blocks,0) as blocks
        from tezos.baker_endorsement_view as bev
        full outer join tezos.blocks_counter_view as bcv on bev.baker = bcv.baker
     ) as r
where (bcvbaker is not null or bevbaker is not null);

CREATE UNIQUE INDEX unique_index ON tezos.baker_view (account_id);

