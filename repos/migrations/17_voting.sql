create index ix_operations_voting_proposal_source_kind_period
  on tezos.operations (proposal, source, kind, period)
  where ((kind::text = 'proposals'::text) or (kind::text = 'ballot'::text)) and proposal is not null;

create index ix_rolls_pkh_block_level
  on tezos.rolls (pkh, block_level);

CREATE VIEW tezos.voting_view AS
select period, proposal, source, rolls, kind, ballot, s.block_level as block_level
from (select period,
             unnest(regexp_split_to_array(replace(replace(proposal :: text, '[', ''), ']', ''), ',')) as proposal,
             kind,
             source,
             min(block_level)                                                                         as block_level,
             min(coalesce(ballot, 'yay'))                                                             as ballot
      from tezos.operations
      where (kind = 'proposals' or kind = 'ballot')
        and proposal is not null
      group by proposal, source, kind, period) as s
       inner join tezos.rolls on (s.source = rolls.pkh and rolls.block_level = s.block_level);

CREATE VIEW tezos.proposal_stat_view AS
select sum(rolls) as rolls, count(1) as bakers, min(block_level) as block_level, proposal, period, kind, ballot
from tezos.voting_view
group by proposal, period, kind, ballot;

Create view tezos.double_voting_by_period as
select  p.period, sum(p.rolls)/2 as rolls, count(1)/2 as bakers
from tezos.voting_view as p inner join tezos.voting_view as w on (p.period=w.period and p.source=w.source and p.proposal<>w.proposal)
group by p.period;

CREATE VIEW tezos.period_stat_view AS
select s.rolls - coalesce(v.rolls, 0) as rolls, s.bakers - coalesce(v.bakers, 0) as bakers, block_level, s.period, kind
from (select sum(rolls) as rolls, sum(bakers) as bakers, min(block_level) as block_level, period, kind
      from tezos.proposal_stat_view
      group by period, kind) as s
       left join tezos.double_voting_by_period as v on s.period = v.period;

CREATE table tezos.voting_period
(
  id          integer,
  type        varchar,
  start_block integer,
  end_block   integer,
  start_time  timestamp,
  end_time    timestamp
);

CREATE VIEW tezos.period_total_stat_view AS
select psv.period, sum(r.rolls) as total_rolls, count(1) as total_bakers
from tezos.period_stat_view as psv
       left join tezos.rolls as r on psv.block_level = r.block_level
group by psv.period;

--
create index ix_operations_double_endorsement_index
  on tezos.operations (operation_id)
  where ((kind)::text = 'double_endorsement_evidence'::text);
