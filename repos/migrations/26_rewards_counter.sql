CREATE VIEW tezos.rewards_counter AS
    SELECT baker,
        staking_balance,
        br.cycle,
        COALESCE (bcb.count, ccb.count) as baking_count,
        COALESCE (bcb.stolen, ccb.stolen) as stolen_baking,
        COALESCE (bcb.reward,  ccb.reward) as baking_rewards,
        COALESCE (bcb.missed,  ccb.missed) as missed_baking,
        COALESCE (bcb.fees,  ccb.fees) as fees,
        fbrv.count as future_baking_count,
        fev.count as future_endorsement_count,
        COALESCE (cev.count, cce.count)     endorsements_count,
        COALESCE (cev.reward,  cce.reward) as endorsement_rewards,
        COALESCE (cev.missed,  cce.missed) as missed_endorsements
    FROM tezos.baking_rewards as br
        left join tezos.baker_cycle_endorsements cev on br.baker = cev.delegate and br.cycle = cev.cycle
        left join tezos.baker_cycle_bakings bcb on br.baker = bcb.delegate and br.cycle = bcb.cycle
        left join tezos.baker_current_cycle_endorsements_view cce on br.baker = cce.delegate and br.cycle = cce.cycle
        left join tezos.baker_current_cycle_bakings_view ccb on br.baker = ccb.delegate and br.cycle = ccb.cycle
        left join tezos.baker_future_baking_rights_view fbrv on br.baker = fbrv.delegate and br.cycle = fbrv.cycle
        left join tezos.baker_future_endorsement_view fev on br.baker = fev.delegate and br.cycle = fev.cycle
ORDER BY br.cycle desc;


