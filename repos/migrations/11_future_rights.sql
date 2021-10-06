--Future baking rights

CREATE OR REPLACE FUNCTION tezos.cycle_by_level(bigint) RETURNS INTEGER AS $cycle_by_level$
    BEGIN
        if $1 <= 1589248 THEN
            RETURN DIV($1 - 1, 4096);
        END IF;

        RETURN 388 + DIV($1 - 1589248 - 1, 8192);
    END;
$cycle_by_level$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION tezos.right_with_cycle() RETURNS trigger AS $right_with_cycle$
    BEGIN

        IF NEW.cycle IS NULL THEN
            NEW.cycle := tezos.cycle_by_level(NEW.block_level);
        END IF;

        RETURN NEW;
    END;
$right_with_cycle$ LANGUAGE plpgsql;

CREATE TRIGGER baking_right_with_cycle
BEFORE INSERT OR UPDATE ON tezos.baking_rights
    FOR EACH ROW EXECUTE FUNCTION tezos.right_with_cycle();

CREATE INDEX IF NOT EXISTS baking_rights_level_priority_index
	ON tezos.baking_rights (block_level, priority);

UPDATE tezos.baking_rights SET cycle = tezos.cycle_by_level(block_level) WHERE cycle is null and block_hash is null;

CREATE OR REPLACE VIEW tezos.future_baking_rights_view AS
SELECT *
FROM tezos.baking_rights
WHERE block_level > (SELECT level FROM tezos.blocks ORDER BY level DESC LIMIT 1);

CREATE OR REPLACE VIEW tezos.baker_future_baking_rights_view AS
SELECT cycle, delegate, avg(priority) avg_priority, sum(zero_priority) AS count
FROM (SELECT delegate, cycle, priority, CASE WHEN priority = 0 THEN 1 ELSE 0 END AS zero_priority
      FROM tezos.future_baking_rights_view
      WHERE priority <= 10
     ) s
GROUP BY cycle, delegate;


--Future endorsing rights
CREATE OR REPLACE VIEW tezos.future_endorsement_rights_view AS
SELECT delegate, block_level, min(er.block_hash) block_hash, array_agg(slot) slots, min(estimated_time) estimated_time, min(cycle) "cycle",
       min(governance_period) governance_period, min(endorsed_block) , min(invalidated_asof) invalidated_asof, min(fork_id) fork_id
FROM tezos.endorsing_rights er
WHERE block_level > (SELECT level FROM tezos.blocks ORDER BY level DESC LIMIT 1)
GROUP BY er.delegate, er.block_level;

CREATE INDEX IF NOT EXISTS endorsing_rights_delegate_cycle_idx
    ON tezos.endorsing_rights USING btree (delegate,cycle);

CREATE TRIGGER endorsing_right_with_cycle
    BEFORE INSERT OR UPDATE ON tezos.endorsing_rights
    FOR EACH ROW EXECUTE FUNCTION tezos.right_with_cycle();

UPDATE tezos.endorsing_rights SET cycle = tezos.cycle_by_level(block_level) WHERE cycle is null and block_hash is null;

CREATE OR REPLACE VIEW tezos.baker_future_endorsement_view as
    select delegate, cycle, count(1) as count
    from tezos.endorsing_rights
    WHERE block_level > (SELECT level FROM tezos.blocks ORDER BY level DESC LIMIT 1)
    group by delegate, cycle;
