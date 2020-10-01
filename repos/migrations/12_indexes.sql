//After sync

CREATE INDEX ix_operations_transactions_source_operation_id
    ON tezos.operations USING btree
    (source, operation_id)
      WHERE kind='transaction';

CREATE INDEX ix_operations_transactions_destination_operation_id
    ON tezos.operations USING btree
    (destination, operation_id)
      WHERE kind='transaction';

CREATE INDEX ix_operations_endorsements_delegate_block_level
    ON tezos.operations USING btree
    (delegate, block_level)
      WHERE kind='endorsement';

CREATE INDEX ix_operations_endorsements_operation_id
    ON tezos.operations USING btree
    (operation_id)
      WHERE kind='endorsement';

CREATE INDEX ix_operations_delegations_operation_id
    ON tezos.operations USING btree
    (operation_id)
      WHERE kind='delegation';

CREATE INDEX ix_operations_endorsements_operation_id
    ON tezos.operations USING btree
    (operation_id)
      WHERE kind='endorsement';

CREATE INDEX ix_accounts_balance
    ON tezos.accounts USING btree
    (balance);

CREATE INDEX ix_operations_double_endorsement_index
  ON tezos.operations (operation_id)
  WHERE ((kind)::text = 'double_endorsement_evidence'::text);

CREATE INDEX accounts_history_account_id_index
  ON tezos.accounts_history (account_id);

CREATE INDEX accounts_account_id_acc_index
	ON tezos.accounts (account_id)
WHERE account_id LIKE 'tz%';

CREATE INDEX accounts_account_id_kt_index
	ON tezos.accounts (account_id)
WHERE account_id LIKE 'KT1%';

CREATE INDEX accounts_delegate_value_index
  ON tezos.accounts (delegate_value)
  WHERE delegate_value IS NOT NULL AND balance > 0;

CREATE index ix_operations_endorsements_level
  ON tezos.operations (level)
  WHERE ((kind)::text = 'endorsement'::text);

CREATE index concurrently balance_updates_source_hash_index
  ON tezos.balance_updates (source_hash)
  WHERE category='rewards' AND source='block';

CREATE index concurrently ix_balance_updates_operation_group_hash_deposits
  ON tezos.balance_updates (operation_group_hash, category)
  WHERE category = 'deposits' AND operation_group_hash IS NOT NULL;

CREATE index balance_updates_operation_group_hash_index
	ON tezos.balance_updates (operation_group_hash);

CREATE INDEX concurrently ix_balance_updates_operation_group_hash_rewards
  ON tezos.balance_updates (operation_group_hash, category)
  WHERE category = 'rewards' AND operation_group_hash IS NOT NULL;

CREATE INDEX accounts_history_asof_index
	ON tezos.accounts_history (asof);

CREATE INDEX accounts_history_block_level_index
	ON tezos.accounts_history (block_level DESC);

CREATE INDEX accounts_history_delegate_value_index
	ON tezos.accounts_history (delegate_value)
	WHERE delegate_value IS NOT NULL;

CREATE INDEX IF NOT EXISTS ix_operations_originated_contract
    ON tezos.operations USING btree
    (originated_contracts)
      WHERE originated_contracts IS NOT NULL;

CREATE INDEX ix_operations_voting_proposal_source_kind_period
  on tezos.operations (proposal, source, kind, period)
  where ((kind::text = 'proposals'::text) or (kind::text = 'ballot'::text)) and proposal is not null;