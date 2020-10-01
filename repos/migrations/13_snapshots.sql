ALTER TABLE tezos.blocks
    ADD UNIQUE (level);

CREATE TABLE tezos.snapshots (
    snp_cycle integer PRIMARY KEY ,
    snp_block_level integer NOT NULL,
    snp_rolls integer NOT NULL,
    CONSTRAINT snapshots_block_foreign FOREIGN KEY (snp_block_level) REFERENCES tezos.blocks (level)
);