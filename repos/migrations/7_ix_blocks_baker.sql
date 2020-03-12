-- +migrate Up

CREATE INDEX IF NOT EXISTS ix_blocks_baker
    ON tezos.blocks USING btree
    (baker ASC NULLS LAST);

-- +migrate Down
