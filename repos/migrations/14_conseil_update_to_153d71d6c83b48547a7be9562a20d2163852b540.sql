CREATE TABLE tezos.processed_chain_events (
    event_level numeric,
    event_type char varying,
    PRIMARY KEY (event_level, event_type)
);

ALTER TABLE tezos.accounts_checkpoint
    ADD COLUMN  cycle integer;

ALTER TABLE tezos.accounts_history
    DROP COLUMN script,
    DROP COLUMN manager,
    DROP COLUMN spendable,
    DROP COLUMN delegate_setable,
    ADD COLUMN  cycle integer;


ALTER TABLE tezos.baking_rights
    ALTER COLUMN block_hash TYPE character varying,
    ALTER COLUMN estimated_time TYPE timestamp without time zone,
    ADD COLUMN cycle integer,
    ADD COLUMN governance_period integer;

ALTER TABLE tezos.endorsing_rights
    ALTER COLUMN block_hash TYPE character varying,
    ALTER COLUMN estimated_time TYPE  timestamp without time zone,
    ADD COLUMN cycle integer,
    ADD COLUMN governance_period integer;

CREATE INDEX ix_balance_updates_op_group_hash ON tezos.balance_updates USING btree (operation_group_hash);