//Refactor
CREATE VIEW tezos.rewards_counter AS
    SELECT baker,
           staking_balance,
           br.cycle,
           cbv.count  as baking_count,
           cbv.stolen    stolen_baking,
           cbv.reward    baking_reward,
           cev.count     endorsements_count,
           cev.reward    endorsements_reward,
           fbrv.count as future_baking_count,
           fev.count     future_endorsement_count
    FROM tezos.baking_rewards as br
           left join tezos.baker_future_baking_rights_view fbrv on br.baker = fbrv.delegate and br.cycle = fbrv.cycle
           left join tezos.baker_cycle_bakings_view cbv on br.baker = cbv.delegate and br.cycle = cbv.cycle
           left join tezos.baker_cycle_endorsements_view cev on br.baker = cev.delegate and br.cycle = cev.cycle
           left join tezos.baker_future_endorsement_view fev on br.baker = fev.delegate and br.cycle = fev.cycle
ORDER BY br.cycle desc;