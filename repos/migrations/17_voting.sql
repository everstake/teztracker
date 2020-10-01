CREATE VIEW tezos.voting_view AS
SELECT period, proposal, source, rolls, kind, ballot, s.block_level AS block_level
FROM (SELECT period,
             unnest(regexp_split_to_array(replace(replace(proposal :: text, '[', ''), ']', ''), ',')) AS proposal,
             kind,
             source,
             min(block_level)                                                                         AS block_level,
             min(coalesce(ballot, 'yay'))                                                             AS ballot
      FROM tezos.operations
      where (kind = 'proposals' or kind = 'ballot')
        and proposal is not null
      GROUP BY proposal, source, kind, period) AS s
       inner join tezos.rolls on (s.source = rolls.pkh and rolls.block_level = s.block_level);

CREATE VIEW tezos.proposal_stat_view AS
SELECT sum(rolls) AS rolls, count(1) AS bakers, min(block_level) AS block_level, proposal, period, kind, ballot
FROM tezos.voting_view
GROUP BY proposal, period, kind, ballot;

CREATE VIEW tezos.double_voting_by_period AS
SELECT  p.period, sum(p.rolls)/2 AS rolls, count(1)/2 AS bakers
FROM tezos.voting_view AS p inner join tezos.voting_view AS w on (p.period=w.period and p.source=w.source and p.proposal<>w.proposal)
GROUP BY p.period;

CREATE VIEW tezos.period_stat_view AS
SELECT s.rolls - coalesce(v.rolls, 0) AS rolls, s.bakers - coalesce(v.bakers, 0) AS bakers, block_level, s.period, kind
FROM (SELECT sum(rolls) AS rolls, sum(bakers) AS bakers, min(block_level) AS block_level, period, kind
      FROM tezos.proposal_stat_view
      GROUP BY period, kind) AS s
       left join tezos.double_voting_by_period AS v on s.period = v.period;

CREATE TABLE tezos.voting_period
(
  id          integer,
  type        varchar,
  start_block integer,
  end_block   integer,
  start_time  timestamp,
  end_time    timestamp
);

CREATE VIEW tezos.period_total_stat_view AS
SELECT psv.period, sum(r.rolls) AS total_rolls, count(1) AS total_bakers
FROM tezos.period_stat_view AS psv
     LEFT JOIN tezos.bakers_history AS r ON psv.block_level = r.block_level
GROUP BY psv.period;

CREATE TABLE tezos.voting_proposal
(
  hash      varchar,
  title        varchar,
  short_description varchar,
  proposal_file varchar,
  proposer varchar
);