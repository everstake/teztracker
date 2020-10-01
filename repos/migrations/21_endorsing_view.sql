CREATE TABLE tezos.baker_endorsements(
  cycle integer,
  delegate varchar,
  level integer,
  slot integer,
  reward integer,
  missed integer,
  PRIMARY KEY (delegate,cycle,level, slot));

CREATE OR REPLACE VIEW tezos.baker_cycle_endorsements_view AS
    SELECT delegate, cycle, sum(reward) reward, sum(missed) missed, count(1) count
    FROM tezos.baker_endorsements
    GROUP BY delegate, cycle;

// After sync
CREATE OR REPLACE FUNCTION tezos.baker_endorsements()
RETURNS trigger
LANGUAGE plpgsql
AS
$$
BEGIN
  insert into tezos.baker_endorsements
  select er.cycle,
         er.delegate,
         er.level,
         er.slot,
         op.reward,
         CASE WHEN op.delegate is null THEN 1 ELSE 0 END as missed
  from tezos.endorsing_rights er
         left join
       (select op.cycle,
               op.delegate,
               op.level,
               json_array_elements_text(slots :: json) as elem,
               change / json_array_length(slots :: json)  reward
        from tezos.operations op
               left join tezos.balance_updates bu
                         on (op.operation_group_hash = bu.operation_group_hash and category = 'rewards')
        where op.kind = 'endorsement'
          and op.level = NEW.meta_level-2) as op on er.level = op.level and op.elem = er.slot::varchar
  where er.level = NEW.meta_level-2;
  RETURN NEW;
END
$$;

CREATE TRIGGER baker_endorsements_insert
  AFTER INSERT
  ON tezos.blocks
  FOR EACH ROW
EXECUTE PROCEDURE tezos.baker_endorsements();