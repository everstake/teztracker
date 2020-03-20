create index balance_updates_operation_group_hash_index
	on tezos.balance_updates (operation_group_hash);

CREATE OR REPLACE VIEW tezos.operations_for_counters AS
select block_level,
       amount,
       fee,
       case when operations.kind = 'endorsement' then 1 else 0 end                 as isendorsement,
       case when operations.kind = 'proposals' then 1 else 0 end                   as isproposals,
       case when operations.kind = 'seed_nonce_revelation' then 1 else 0 end       as isseed_nonce_revelation,
       case when operations.kind = 'delegation' then 1 else 0 end                  as isdelegation,
       case when operations.kind = 'transaction' then 1 else 0 end                 as istransaction,
       case when operations.kind = 'activate_account' then 1 else 0 end            as isactivate_account,
       case when operations.kind = 'ballot' then 1 else 0 end                      as isballot,
       case when operations.kind = 'origination' then 1 else 0 end                 as isorigination,
       case when operations.kind = 'reveal' then 1 else 0 end                      as isreveal,
       case when operations.kind = 'double_baking_evidence' then 1 else 0 end      as isdouble_baking_evidence,
       case when operations.kind = 'double_endorsement_evidence' then 1 else 0 end as isdouble_endorsement_evidence,
       consumed_gas
from tezos.operations;

CREATE OR REPLACE VIEW tezos.block_aggregation_view
    AS
    SELECT operations.block_level AS level,
    COALESCE(sum(operations.amount), 0::numeric) AS volume,
    COALESCE(sum(operations.fee), 0::numeric) AS fees,
    sum(operations.isendorsement) AS endorsements,
    sum(operations.isproposals) AS proposals,
    sum(operations.isseed_nonce_revelation) AS seed_nonce_revelations,
    sum(operations.isdelegation) AS delegations,
    sum(operations.istransaction) AS transactions,
    sum(operations.isactivate_account) AS activate_accounts,
    sum(operations.isballot) AS ballots,
    sum(operations.isorigination) AS originations,
    sum(operations.isreveal) AS reveals,
    sum(operations.isdouble_baking_evidence) AS double_baking_evidences,
    sum(operations.isdouble_endorsement_evidence) AS double_endorsement_evidences,
    count(1) - sum(operations.isendorsement) as number_of_operations,
    COALESCE(sum(operations.consumed_gas), 0::numeric) AS gas_used
   FROM tezos.operations_for_counters operations
  GROUP BY operations.block_level;
