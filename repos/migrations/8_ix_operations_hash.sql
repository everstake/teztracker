-- +migrate Up

CREATE INDEX IF NOT EXISTS ix_operations_hash
    ON tezos.operations USING btree
    (operation_group_hash ASC NULLS LAST);

-- +migrate Down

