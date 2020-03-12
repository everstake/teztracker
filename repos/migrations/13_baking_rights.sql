-- +migrate Up

Create table tezos.operation_counters(
    cnt_id SERIAL PRIMARY KEY,
    cnt_last_op_id int not null,
    cnt_operation_type varchar(100) NOT NULL,
    cnt_count bigint not null,
    cnt_created_at timestamp with time zone NULL DEFAULT NULL,
    CONSTRAINT operation_counters_last_op_foreign FOREIGN KEY (cnt_last_op_id) REFERENCES tezos.operations (operation_id)
);

CREATE TABLE tezos.future_baking_rights (
    level integer NOT NULL,
    delegate character varying NOT NULL,
    priority integer NOT NULL,
    estimated_time timestamp without time zone NOT NULL,
    PRIMARY KEY (level, priority)
);

CREATE INDEX future_baking_rights_delegate_idx
    ON tezos.future_baking_rights USING btree (delegate);

alter table tezos.blocks ADD UNIQUE (level);

CREATE TABLE tezos.snapshots (
    snp_cycle integer PRIMARY KEY ,
    snp_block_level integer NOT NULL,
    snp_rolls integer NOT NULL,
    CONSTRAINT snapshots_block_foreign FOREIGN KEY (snp_block_level) REFERENCES tezos.blocks (level)
);

CREATE TABLE tezos.double_baking_evidences (
    operation_id integer PRIMARY KEY,
    dbe_block_hash character varying NOT NULL,
    dbe_block_level integer NOT NULL,
    dbe_denounced_level integer NOT NULL,
    dbe_offender character varying NOT NULL,
    dbe_priority integer NOT NULL,
    dbe_evidence_baker character varying NOT NULL,
    dbe_baker_reward numeric NOT NULL,
    dbe_lost_deposits numeric NOT NULL,
    dbe_lost_rewards numeric NOT NULL,
    dbe_lost_fees numeric NOT NULL,
    CONSTRAINT double_baking_evidences_block_foreign FOREIGN KEY (dbe_block_level) REFERENCES tezos.blocks (level),
    CONSTRAINT double_baking_evidences_denounced_block_foreign FOREIGN KEY (dbe_denounced_level) REFERENCES tezos.blocks (level),
    CONSTRAINT double_baking_evidences_operation_foreign FOREIGN KEY (operation_id) REFERENCES tezos.operations (operation_id)
);

-- +migrate Down