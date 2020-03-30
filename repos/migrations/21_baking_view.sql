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