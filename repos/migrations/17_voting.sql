CREATE TABLE tezos.rolls (
    pkh varchar,
    rolls integer NOT NULL,
    block_level  int not null,
    cycle int not null,
    voting_period  int not null
);

--TODO Add indexes

CREATE OR REPLACE VIEW tezos.voting_view AS
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
       inner join tezos.rolls on (s.source = rolls.pkh and rolls.voting_period = s.period);

CREATE TABLE tezos.baker_voting AS SELECT * from tezos.voting_view;

CREATE OR REPLACE FUNCTION tezos.baker_voting()
RETURNS trigger
LANGUAGE plpgsql
AS
$$
    DECLARE
    rollsN integer;
    BEGIN
    --     Skip non endorsement operations
        IF (NEW.kind != 'proposals' AND NEW.kind != 'ballot' ) THEN
            RETURN NEW;
        END IF;

        select rolls into rollsN from tezos.rolls where (rolls.pkh = NEW.source AND rolls.voting_period = NEW.period);

        INSERT INTO tezos.baker_voting(period, proposal, kind, source, block_level, ballot, rolls)
        VALUES(
               NEW.period,
               unnest(regexp_split_to_array(replace(replace(NEW.proposal :: text, '[', ''), ']', ''), ',')),
               NEW.kind,
               NEW.source,
               NEW.block_level,
               coalesce(ballot, 'yay'),
               rollsN
               );
    RETURN NEW;
    END
$$;

CREATE TRIGGER baker_voting_insert
  AFTER INSERT
  ON tezos.operations
  FOR EACH ROW
EXECUTE PROCEDURE tezos.baker_voting();

CREATE OR REPLACE VIEW tezos.proposal_stat_view AS
SELECT sum(rolls) AS rolls, count(1) AS bakers, min(block_level) AS block_level, proposal, period, kind, ballot
FROM tezos.baker_voting
GROUP BY proposal, period, kind, ballot;

CREATE OR REPLACE VIEW tezos.voting_participation AS
select period, count(1) bakers, sum(rolls) rolls FROM
(select  DISTINCT period, source, rolls  from tezos.baker_voting) s GROUP BY period;

CREATE OR REPLACE VIEW tezos.period_stat_view AS
SELECT vp.rolls AS rolls, vp.bakers AS bakers, block_level, s.period, kind
FROM (SELECT min(block_level) AS block_level, period, kind
      FROM tezos.proposal_stat_view
      GROUP BY period, kind) AS s
       left join tezos.voting_participation AS vp on s.period = vp.period;

CREATE TABLE tezos.voting_period
(
  id          integer,
  type        varchar,
  start_block integer,
  end_block   integer,
  start_time  timestamp,
  end_time    timestamp
);

CREATE OR REPLACE VIEW tezos.period_total_stat_view AS
SELECT psv.period, sum(r.rolls) AS total_rolls, count(1) AS total_bakers
FROM tezos.period_stat_view AS psv
     LEFT JOIN tezos.rolls AS r ON psv.period = r.voting_period
GROUP BY psv.period;

CREATE TABLE tezos.voting_proposal
(
  hash      varchar,
  title        varchar,
  short_description varchar,
  proposal_file varchar,
  proposer varchar,
  is_main bool default false not null,
  period integer not null
);