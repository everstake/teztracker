CREATE TABLE tezos.future_baking_rights (
    cycle integer NOT NULL,
    level integer NOT NULL,
    delegate character varying NOT NULL,
    priority integer NOT NULL,
    estimated_time timestamp without time zone NOT NULL,
    PRIMARY KEY (cycle, level, delegate, priority)
);

CREATE INDEX future_baking_rights_delegate_idx
    ON tezos.future_baking_rights USING btree (delegate);

CREATE INDEX future_baking_rights_level_priority_index
	ON tezos.future_baking_rights (level, priority);

CREATE VIEW tezos.baker_future_baking_rights_view AS
SELECT cycle, delegate, avg(priority) avg_priority, sum(zero_priority) AS count
FROM (SELECT delegate, cycle, priority, CASE WHEN priority = 0 THEN 1 ELSE 0 END AS zero_priority
      FROM tezos.future_baking_rights
      WHERE level > (SELECT level FROM tezos.blocks ORDER BY level DESC limit 1)
        AND priority <= 10
     ) s
GROUP BY cycle, delegate;

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

CREATE VIEW tezos.baker_future_endorsement_view as
    select delegate, cycle, sum(array_length(slots, 1)) as count
    from tezos.future_endorsement_rights
    where level > (select level from tezos.blocks order by level desc limit 1)
    group by delegate, cycle;