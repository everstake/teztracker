CREATE TABLE tezos.delegators_by_cycle(
    account_id varchar NOT NULL,
    cycle integer NOT NULL,
    level integer NOT NULL,
    delegate_value varchar NOT NULL,
    balance numeric NOT NULL,
    PRIMARY KEY (account_id, cycle, delegate_value)
);

CREATE TABLE tezos.baking_rewards
(
  baker           varchar NOT NULL,
  cycle           integer NOT NULL,
  delegators      integer,
  staking_balance bigint,
  constraint baking_rewards_pk
  PRIMARY KEY (baker, cycle)
);

CREATE VIEW tezos.baker_delegators_by_cycle AS
SELECT delegate_value baker, cycle, sum(balance) staking_balance, count(1) delegators
from tezos.delegators_by_cycle
group by delegate_value, cycle;

//After sync
CREATE OR REPLACE FUNCTION tezos.baking_rewards()
 RETURNS trigger LANGUAGE plpgsql
AS $$
    BEGIN
        INSERT INTO tezos.baking_rewards
        SELECT delegate_value, NEW.snp_cycle, count(1), sum(balance)
        FROM (SELECT account_id, max(block_level) block_level
              FROM tezos.accounts_history
                     LEFT JOIN tezos.snapshots ON NEW.snp_cycle = snp_cycle
              where block_level <= snp_block_level
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

CREATE OR REPLACE FUNCTION tezos.delegators_by_cycle()
 RETURNS trigger LANGUAGE plpgsql
AS $$
BEGIN
  insert into tezos.delegators_by_cycle
  SELECT s.account_id, NEW.snp_cycle, s.block_level, delegate_value, balance
  from (SELECT account_id, max(block_level) block_level
        from tezos.accounts_history
               left join tezos.snapshots on NEW.snp_cycle = snp_cycle
        where block_level <= snp_block_level
        group by account_id) s
         left join tezos.accounts_history ah on s.account_id = ah.account_id and s.block_level = ah.block_level
  where delegate_value is not null
    and s.account_id <> delegate_value
    and balance > 0;
RETURN NEW;
END $$;

CREATE TRIGGER delegators_by_cycle_insert
  AFTER INSERT
  ON tezos.snapshots
  FOR EACH ROW
EXECUTE PROCEDURE tezos.delegators_by_cycle();

CREATE OR REPLACE VIEW tezos.frozen_endorsement_rewards as
SELECT delegate,sum(reward) rewards, sum(count) count
from tezos.baker_cycle_endorsements_view
where cycle >= (SELECT meta_cycle from tezos.blocks order by level desc limit 1) - 5 group by delegate;

CREATE OR REPLACE VIEW tezos.frozen_baking_rewards as
SELECT delegate,sum(reward) rewards, sum(count) count
from tezos.baker_cycle_bakings_view
where cycle >= (SELECT meta_cycle from tezos.blocks order by level desc limit 1) - 5 group by delegate;


//New

CREATE TABLE tezos.baker_cycle_endorsements
 AS (SELECT * FROM baker_cycle_endorsements_view);

alter table tezos.baker_cycle_endorsements
	add constraint baker_cycle_endorsements_pk
		primary key (delegate, cycle);

CREATE OR REPLACE VIEW tezos.baker_current_cycle_endorsements_view AS
    SELECT delegate, cycle, sum(reward) reward, sum(missed) missed, count(1) count
    FROM tezos.baker_endorsements
    WHERE cycle = (select meta_cycle from tezos.blocks order by level desc limit 1)
    GROUP BY delegate, cycle;

CREATE TABLE tezos.baker_cycle_bakings
 AS (SELECT * FROM baker_cycle_bakings_view);

alter table tezos.baker_cycle_bakings
	add constraint baker_cycle_bakings_pk
		primary key (delegate, cycle);

CREATE OR REPLACE VIEW tezos.baker_current_cycle_bakings_view AS
    select cycle, delegate, avg(priority) avg_priority, sum(reward) reward, sum(baked) count, sum(missed) missed, sum(stolen) stolen, sum(fees) fees
    from tezos.baker_bakings
    WHERE cycle = (select meta_cycle from tezos.blocks order by level desc limit 1)
    GROUP BY delegate, cycle;