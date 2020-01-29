CREATE MATERIALIZED VIEW tezos.baker_view AS
select delegates.pkh as account_id, staking_balance,count(1) as endorsements from tezos.delegates
  inner join tezos.operations ON delegates.pkh = operations.delegate
WHERE kind = 'endorsement' and deactivated=false
group by delegates.pkh,staking_balance;

CREATE UNIQUE INDEX unique_index ON tezos.baker_view (account_id);