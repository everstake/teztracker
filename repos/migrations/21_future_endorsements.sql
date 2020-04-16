CREATE TABLE tezos.future_endorsement_rights (
    level integer NOT NULL,
    cycle integer NOT NULL,
    delegate character varying NOT NULL,
    slots integer[] NOT NULL,
    estimated_time timestamp without time zone NOT NULL,
    PRIMARY KEY (level, delegate, cycle)
);


CREATE INDEX future_endorsement_rights_delegate_cycle_idx
    ON tezos.future_endorsement_rights USING btree (delegate,cycle);