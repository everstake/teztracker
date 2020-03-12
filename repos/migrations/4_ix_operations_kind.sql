-- +migrate Up
CREATE INDEX IF NOT EXISTS ix_operations_kind
    ON tezos.operations USING btree
    (kind ASC NULLS LAST);

-- +migrate Down
