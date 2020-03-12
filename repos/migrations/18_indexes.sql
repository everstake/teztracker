-- +migrate Up

create index ix_operations_voting_proposal_source_kind_period
  on tezos.operations (proposal, source, kind, period)
  where ((kind::text = 'proposals'::text) or (kind::text = 'ballot'::text)) and proposal is not null;

create index ix_rolls_pkh_block_level
  on tezos.rolls (pkh, block_level);

create index ix_operations_double_endorsement_index
  on tezos.operations (operation_id)
  where ((kind)::text = 'double_endorsement_evidence'::text);

-- +migrate Down