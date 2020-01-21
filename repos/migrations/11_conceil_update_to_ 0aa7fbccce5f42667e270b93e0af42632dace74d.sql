CREATE TABLE accounts_history (
    account_id character varying NOT NULL,
    block_id character varying NOT NULL,
    counter integer,
    script character varying,
    storage character varying,
    balance numeric NOT NULL,
    block_level numeric DEFAULT '-1'::integer NOT NULL,
    manager character varying, -- retro-compat from protocol 5+
    spendable boolean, -- retro-compat from protocol 5+
    delegate_setable boolean, -- retro-compat from protocol 5+
    delegate_value char varying, -- retro-compat from protocol 5+
    asof timestamp with time zone NOT NULL
);

CREATE TABLE baking_rights (
    block_hash character varying NOT NULL,
    level integer NOT NULL,
    delegate character varying NOT NULL,
    priority integer NOT NULL,
    estimated_time timestamp without time zone NOT NULL,
    PRIMARY KEY (level, delegate),
    CONSTRAINT fk_block_hash FOREIGN KEY (block_hash) REFERENCES blocks(hash) NOT VALID
);

CREATE TABLE endorsing_rights (
    block_hash character varying NOT NULL,
    level integer NOT NULL,
    delegate character varying NOT NULL,
    slot integer NOT NULL,
    estimated_time timestamp without time zone NOT NULL,
    PRIMARY KEY (level, delegate, slot),
    CONSTRAINT fk_block_hash FOREIGN KEY (block_hash) REFERENCES blocks(hash) NOT VALID
);
ALTER TABLE accounts
ALTER COLUMN manager DROP NOT NULL,
ALTER COLUMN spendable DROP NOT NULL,
ALTER COLUMN delegate_setable DROP NOT NULL,
ALTER COLUMN counter DROP NOT NULL,
ALTER COLUMN balance DROP NOT NULL,
ALTER COLUMN block_level DROP NOT NULL;


ALTER TABLE accounts_checkpoint 
    ADD COLUMN  asof timestamp with time zone;

UPDATE accounts_checkpoint SET asof=NOW();

ALTER TABLE accounts_checkpoint 
    ALTER COLUMN asof SET NOT NULL;


ALTER TABLE balance_updates
    ADD COLUMN operation_group_hash character varying;



ALTER TABLE fees
    ADD COLUMN cycle integer,
    ADD COLUMN level integer;


ALTER TABLE operations 
    ADD COLUMN branch character varying,
    ADD COLUMN number_of_slots integer,
    ADD COLUMN cycle integer,
    ADD COLUMN proposal character varying,
    ADD COLUMN ballot character varying,
    ADD COLUMN internal boolean, 
    ADD COLUMN period integer;

update operations set internal = false;

ALTER TABLE operations
    ALTER COLUMN internal SET  NOT NULL;


CREATE INDEX baking_rights_level_idx ON baking_rights USING btree (level);
CREATE INDEX endorsing_rights_level_idx ON endorsing_rights USING btree (level);
CREATE INDEX fki_fk_block_hash ON baking_rights USING btree (block_hash);
CREATE INDEX fki_fk_block_hash2 ON endorsing_rights USING btree (block_hash);
CREATE INDEX ix_operation_groups_block_level ON operation_groups USING btree (block_level);
CREATE INDEX ix_rolls_block_id ON rolls USING btree (block_id);
CREATE INDEX ix_rolls_block_level ON rolls USING btree (block_level);

-- Conceil no longer uses those tables, so we can drop those. But let's leave it just in case we need it to repopulate new fields from other tables.
-- DROP TABLE ballots;
-- DROP TABLE delegated_contracts;
-- DROP TABLE proposals;