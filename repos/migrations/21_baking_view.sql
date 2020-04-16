CREATE TABLE tezos.baker_bakings(
  cycle integer,
  delegate varchar,
  level integer,
  priority integer,
  baked integer,
  reward integer,
  missed integer,
  stolen integer,
  PRIMARY KEY (delegate,cycle,level));

CREATE VIEW tezos.baker_cycle_bakings_view as
select cycle,
       delegate,
       avg(priority) avg_priority,
       sum(reward)          reward,
       sum(baked)        count,
       sum(missed)       missed,
       sum(stolen)       stolen
from tezos.baker_bakings
group by cycle, delegate;

CREATE OR REPLACE FUNCTION tezos.baker_bakings()
RETURNS trigger
LANGUAGE plpgsql
AS
$$
BEGIN
  insert into tezos.baker_bakings
select cycle,
       br.delegate,
       br.level,
       CASE WHEN baker = br.delegate THEN bl.priority ELSE NULL END        as priority,
       CASE WHEN baker = br.delegate THEN 1 ELSE 0 END                     as baked,
       CASE WHEN baker = br.delegate THEN bu.change ELSE 0 END             as reward,
       CASE WHEN bl.priority > br.priority THEN 1 ELSE 0 END               as missed,
       CASE WHEN br.priority > 0 and baker = br.delegate THEN 1 ELSE 0 END as stolen
from tezos.baking_rights br
       left join tezos.blocks bl on (br.level = bl.meta_level)
       left join tezos.balance_updates bu on source_hash = hash
where category = 'rewards'
  and change > 0
  and source = 'block'
  and bl.level = NEW.meta_level-2
  and (baker = br.delegate or bl.priority > br.priority or (br.priority > 0 and baker = br.delegate));
  RETURN NEW;
END
$$;

CREATE TRIGGER baker_bakings_insert
  AFTER INSERT
  ON tezos.blocks
  FOR EACH ROW
EXECUTE PROCEDURE tezos.baker_bakings();

alter table tezos.future_baking_rights
	add cycle integer;


alter table tezos.future_baking_rights
    add constraint future_baking_rights_pk
	    primary key (cycle, level, delegate, priority);

CREATE MATERIALIZED VIEW tezos.future_baking_rights_materialized_view as
select cycle, delegate, avg(priority) avg_priority, sum(zero_priority) as count
from (select delegate, cycle, priority, case when priority = 0 then 1 else 0 end as zero_priority
      from tezos.future_baking_rights
      where level > (select level from tezos.blocks order by level desc limit 1)
        and priority <= 10
     ) s
group by cycle, delegate;

create index concurrently ix_balance_updates_operation_group_hash_rewards
  on tezos.balance_updates (operation_group_hash, category) where category = 'rewards' and operation_group_hash IS NOT NULL;