CREATE TABLE tezos.baker_endorsements(
  cycle integer,
  delegate varchar,
  level integer,
  slot integer,
  reward integer,
  missed integer,
  PRIMARY KEY (delegate,cycle,level, slot));

CREATE OR REPLACE VIEW tezos.baker_cycle_endorsements_view AS
    SELECT delegate, cycle, reward, missed, count - missed count from
    (SELECT delegate, cycle, sum(reward) reward, sum(missed) missed, count(1) count
    FROM tezos.baker_endorsements
    GROUP BY delegate, cycle) s;

CREATE TABLE tezos.baker_cycle_endorsements
 AS (SELECT * FROM tezos.baker_cycle_endorsements_view);

alter table tezos.baker_cycle_endorsements
	add constraint baker_cycle_endorsements_pk
		primary key (delegate, cycle);

CREATE OR REPLACE VIEW tezos.baker_current_cycle_endorsements_view AS
    SELECT delegate, cycle, sum(reward) reward, sum(missed) missed, count(1) count
    FROM tezos.baker_endorsements
    WHERE cycle = (select meta_cycle from tezos.blocks order by level desc limit 1)
    GROUP BY delegate, cycle;


CREATE OR REPLACE FUNCTION tezos.endorsement_reward(integer, integer)
   RETURNS integer
   LANGUAGE plpgsql
AS $$
    declare
        reward integer;
    BEGIN

    	case
		  when $1 >= 388 then
             reward = 78125;
		  when $1 >= 208 then
             reward = 1250000;
		  else
	        RETURN 2000000;
		 end case;

        IF $2 > 0 THEN
            reward = reward / 1.5;
        END IF;
    RETURN reward;
END $$;


CREATE OR REPLACE FUNCTION tezos.baker_endorsement()
RETURNS trigger
LANGUAGE plpgsql
AS
$$
    DECLARE
    priority integer;
    BEGIN
    --     Skip non endorsement operations
        IF (NEW.kind != 'endorsement' AND NEW.kind != 'endorsement_with_slot' ) THEN
            RETURN NEW;
        END IF;

        select blocks.priority into priority from tezos.blocks where level = NEW.block_level;

        insert into tezos.baker_endorsements(cycle, delegate, level, slot, reward, missed)
        VALUES (
        NEW.cycle :: integer,
        NEW.delegate,
        NEW.block_level -1, --meta_level
        json_array_elements_text(NEW.slots :: json) :: integer,
        tezos.endorsement_reward(NEW.cycle, priority),
        0)
        ON CONFLICT ON CONSTRAINT baker_endorsements_pkey
        DO UPDATE SET reward = excluded.reward, missed = excluded.missed;
    RETURN NEW;
    END
$$;

CREATE TRIGGER baker_endorsement_insert
  AFTER INSERT
  ON tezos.operations
  FOR EACH ROW
EXECUTE PROCEDURE tezos.baker_endorsement();

CREATE OR REPLACE FUNCTION tezos.baker_endorsements_block()
RETURNS trigger
LANGUAGE plpgsql
AS
$$
BEGIN

  insert into tezos.baker_endorsements(cycle, delegate, level, slot, reward, missed)
  select er.cycle,
         er.delegate,
         er.block_level,
         er.slot,
         op.reward,
         CASE WHEN op.delegate is null THEN 1 ELSE 0 END as missed
  from tezos.endorsing_rights er
         left join
       (select op.cycle,
               op.delegate,
               op.level,
               json_array_elements_text(slots :: json) as elem,
               tezos.endorsement_reward(op.cycle, bu.priority) as reward
        from tezos.operations op
               left join tezos.blocks bu
                         on (op.block_level = bu.level)
        where (op.kind = 'endorsement' OR op.kind = 'endorsement_with_slot')
        ) as op on er.block_level = op.level and op.elem = er.slot::varchar
        where er.block_level = NEW.meta_level-7 ON CONFLICT DO NOTHING;

  IF NEW.meta_cycle_position <= 9 THEN
   INSERT INTO tezos.baker_cycle_endorsements (SELECT * FROM tezos.baker_cycle_endorsements_view
    where tezos.baker_cycle_endorsements_view.cycle = NEW.meta_cycle-1)
    ON CONFLICT ON CONSTRAINT baker_cycle_endorsements_pk
    DO UPDATE SET reward = excluded.reward, missed = excluded.missed, count = excluded.count;
  END IF;
  RETURN NEW;
END
$$;

CREATE TRIGGER baker_cycle_endorsements_block_insert
  AFTER INSERT
  ON tezos.blocks
  FOR EACH ROW
EXECUTE PROCEDURE tezos.baker_endorsements_block();