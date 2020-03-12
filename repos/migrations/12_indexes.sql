-- +migrate Up

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

CREATE INDEX ix_accounts_balance
    ON tezos.accounts USING btree
    (balance);

-- +migrate Down
