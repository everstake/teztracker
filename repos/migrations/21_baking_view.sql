CREATE TABLE tezos.baker_bakings(
  cycle integer,
  delegate varchar,
  level integer,
  priority integer,
  baked integer,
  reward integer,
  fees numeric,
  missed integer,
  stolen integer,
  PRIMARY KEY (delegate,cycle,level));

CREATE OR REPLACE VIEW tezos.baker_cycle_bakings_view as
select cycle,
       delegate,
       avg(priority) avg_priority,
       sum(reward)       reward,
       sum(baked)        count,
       sum(missed)       missed,
       sum(stolen)       stolen,
       sum(fees)         fees
from tezos.baker_bakings
group by cycle, delegate;

//After sync
CREATE OR REPLACE FUNCTION tezos.baker_bakings()
RETURNS trigger
LANGUAGE plpgsql
AS
$$
BEGIN
insert into tezos.baker_bakings (cycle,delegate,level,priority,baked,reward,fees,missed,stolen)
       select meta_cycle,
       br.delegate,
       br.level,
       CASE WHEN baker = br.delegate THEN bl.priority ELSE NULL END        as priority,
       CASE WHEN baker = br.delegate THEN 1 ELSE 0 END                     as baked,
       CASE WHEN baker = br.delegate THEN bu.change ELSE 0 END             as reward,
       CASE WHEN baker = br.delegate THEN bav.fees ELSE 0 END              as fees,
       CASE WHEN bl.priority > br.priority THEN 1 ELSE 0 END               as missed,
       CASE WHEN br.priority > 0 and baker = br.delegate THEN 1 ELSE 0 END as stolen
from tezos.baking_rights br
       left join tezos.blocks bl on (br.level = bl.meta_level)
       left join tezos.balance_updates bu on source_hash = hash
       left join tezos.block_aggregation_view bav on bav.level = bl.level
where category = 'rewards'
  and change > 0
  and source = 'block'
  and bl.level = NEW.meta_level-5
  and (baker = br.delegate or bl.priority > br.priority or (br.priority > 0 and baker = br.delegate));
  IF NEW.meta_cycle_position <= 5 THEN
    INSERT INTO tezos.baker_cycle_bakings (SELECT * FROM tezos.baker_cycle_bakings_view
    where tezos.baker_cycle_bakings_view.cycle = NEW.meta_level-1)
    ON CONFLICT ON CONSTRAINT baker_cycle_bakings_pk
    DO UPDATE SET avg_priority = excluded.avg_priority, reward = excluded.reward, count = excluded.count, missed = excluded.missed,stolen = excluded.stolen,fees = excluded.fees;
  END IF;
  RETURN NEW;
END
$$;

CREATE TRIGGER baker_bakings_insert
  AFTER INSERT
  ON tezos.blocks
  FOR EACH ROW
EXECUTE PROCEDURE tezos.baker_bakings();
