-- +migrate Up

CREATE INDEX IF NOT EXISTS  roll_by_level
    ON tezos.rolls USING btree
    (block_level DESC NULLS LAST);

-- +migrate Down