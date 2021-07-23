CREATE OR REPLACE VIEW tezos.blocks_counter_view AS
SELECT baker, count(1) as blocks
FROM tezos.blocks
GROUP BY baker;

create or replace view tezos.baker_delegations_view as
select delegate_value as baker, count(1) as active_delegations
from tezos.accounts
where delegate_value is not null
  and account_id != delegate_value
  and balance > 0
group by delegate_value;

CREATE TABLE tezos.baker_baking_since
(
  pkh            varchar not null
    constraint baker_baking_since_pkey
      primary key,
  baking_since          timestamp    not null
);

CREATE OR REPLACE FUNCTION insert_baker_baking_since()
  RETURNS TRIGGER LANGUAGE plpgsql
  AS $$
  BEGIN
      insert into tezos.baker_baking_since(pkh, baking_since)
      VALUES (NEW.pkh, now());
  RETURN NULL;
  END $$;

CREATE TRIGGER baker_baking_since
  AFTER INSERT
  ON tezos.bakers
  FOR EACH ROW
EXECUTE PROCEDURE insert_baker_baking_since();

INSERT INTO tezos.baker_baking_since(pkh, baking_since)
SELECT pkh, min(asof)  from tezos.bakers_history group by pkh;

CREATE OR REPLACE VIEW tezos.baker_endorsement_view as
SELECT delegate,
       count(1) AS endorsements
FROM tezos.baker_endorsements
where missed = 0
GROUP BY delegate;

CREATE MATERIALIZED VIEW tezos.baker_view AS
SELECT b.pkh account_id,
       balance,
       frozen_balance,
       staking_balance,
       delegated_balance,
       COALESCE(bdv.active_delegations, 0) as active_delegations,
       fer.rewards                         as frozen_endorsement_rewards,
       fer.count                              endorsement_count,
       fbr.rewards                            frozen_baking_rewards,
       fbr.count                              baking_count,
       TRUNC(staking_balance / 8000 / 1000000, 0) as rolls,
       COALESCE(bev.endorsements, (0)::bigint) AS endorsements,
       COALESCE(bcv.blocks, (0)::bigint)       AS blocks,
       baking_since
FROM tezos.bakers b
LEFT JOIN baker_endorsement_view bev on b.pkh = bev.delegate
LEFT JOIN tezos.blocks_counter_view bcv on b.pkh = bcv.baker
LEFT JOIN baker_baking_since bbs on b.pkh = bbs.pkh
LEFT JOIN tezos.baker_delegations_view as bdv on b.pkh = bdv.baker
LEFT JOIN tezos.frozen_baking_rewards as fbr on b.pkh = fbr.delegate
LEFT JOIN tezos.frozen_endorsement_rewards as fer on b.pkh = fer.delegate
where deactivated IS FALSE;

CREATE UNIQUE INDEX unique_index ON tezos.baker_view (account_id);

