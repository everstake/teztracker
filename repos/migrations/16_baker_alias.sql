CREATE TABLE tezos.baker_alias (
    address        varchar,
    name           varchar,
    CONSTRAINT address PRIMARY KEY(address)
);

CREATE INDEX future_baking_rights_level_priority_index
	ON tezos.future_baking_rights (level, priority);