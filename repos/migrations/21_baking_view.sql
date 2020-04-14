create materialized view tezos.baking_materialized_view as
select cycle, delegate, avg(baking_priority) avg_priority , sum(reward) reward, sum(is_baked) count, sum(is_missed) missed, sum(is_stolen) stolen
from (select cycle,
             br.delegate,
             CASE WHEN baker = br.delegate THEN bu.change ELSE 0 END             as reward,
             CASE WHEN baker = br.delegate THEN bl.priority ELSE NULL END        as baking_priority,
             CASE WHEN baker = br.delegate THEN 1 ELSE 0 END                     as is_baked,
             CASE WHEN bl.priority > br.priority THEN 1 ELSE 0 END               as is_missed,
             CASE WHEN br.priority > 0 and baker = br.delegate THEN 1 ELSE 0 END as is_stolen
      from tezos.baking_rights br
             left join tezos.blocks bl on (br.level = bl.meta_level)
             left join tezos.balance_updates bu on source_hash = hash
        where category = 'rewards'
        and change > 0
        and source = 'block'
     ) as s
group by cycle, delegate;

alter table tezos.future_baking_rights
	add cycle integer;


alter table tezos.future_baking_rights
    add constraint future_baking_rights_pk
	    primary key (cycle, level, delegate, priority);

CREATE MATERIALIZED VIEW tezos.future_baking_rights_materialized_view as
select cycle, delegate, avg(priority), sum(zero_priority) as count, sum(zero_priority) * 40 as rewards
from (select delegate, cycle, priority, case when priority = 0 then 1 else 0 end as zero_priority
      from tezos.future_baking_rights
      where cycle  > (select meta_cycle from tezos.blocks order by level desc limit 1)) s
group by cycle, delegate;

create index concurrently ix_balance_updates_operation_group_hash_rewards
  on tezos.balance_updates (operation_group_hash, category) where category = 'rewards' and operation_group_hash IS NOT NULL;

CREATE TABLE tezos.baker_endorsements(
  cycle integer,
  delegate varchar,
  level integer,
  slot integer,
  reward integer,
  missed integer,
  PRIMARY KEY (delegate,cycle,level));


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