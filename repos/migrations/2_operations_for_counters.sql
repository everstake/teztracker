-- +migrate Up

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
       case when operations.kind = 'double_endorsement_evidence' then 1 else 0 end as isdouble_endorsement_evidence
from tezos.operations;

-- +migrate Down
