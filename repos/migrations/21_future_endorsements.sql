CREATE TABLE tezos.future_endorsement_rights (
    level integer NOT NULL,
    cycle integer NOT NULL,
    delegate character varying NOT NULL,
    slots integer[] NOT NULL,
    PRIMARY KEY (level, delegate, cycle)
);


CREATE INDEX future_baking_rights_delegate_idx
    ON tezos.future_baking_rights USING btree (delegate,cycle);