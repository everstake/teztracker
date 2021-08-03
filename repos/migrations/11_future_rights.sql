--Future baking rights

CREATE OR REPLACE FUNCTION tezos.right_with_cycle() RETURNS trigger AS $right_with_cycle$
    declare
        blocksPerCycle integer;
    BEGIN

        IF NEW.cycle IS NULL THEN
        -- TODO check level
            IF NEW.block_level >= 1613825 THEN
             blocksPerCycle = 8192;
            ELSE
             blocksPerCycle = 4096;
            END IF;

            NEW.cycle := DIV(NEW.block_level - 1, blocksPerCycle);
        END IF;

        RETURN NEW;
    END;
$right_with_cycle$ LANGUAGE plpgsql;

CREATE TRIGGER baking_right_with_cycle
BEFORE INSERT OR UPDATE ON tezos.baking_rights
    FOR EACH ROW EXECUTE FUNCTION tezos.right_with_cycle();

CREATE INDEX IF NOT EXISTS baking_rights_level_priority_index
	ON tezos.baking_rights (block_level, priority);

UPDATE tezos.baking_rights SET cycle = DIV(block_level - 1, 8192) WHERE cycle is null and block_hash is null and block_level >= 1613825;
UPDATE tezos.baking_rights SET cycle = DIV(block_level - 1, 4096) WHERE cycle is null and block_hash is null and block_level < 1613825;


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
SELECT *
FROM tezos.endorsing_rights
WHERE block_level > (SELECT level FROM tezos.blocks ORDER BY level DESC LIMIT 1);

CREATE INDEX IF NOT EXISTS endorsing_rights_delegate_cycle_idx
    ON tezos.endorsing_rights USING btree (delegate,cycle);

CREATE TRIGGER endorsing_right_with_cycle
    BEFORE INSERT OR UPDATE ON tezos.endorsing_rights
    FOR EACH ROW EXECUTE FUNCTION tezos.right_with_cycle();

UPDATE tezos.endorsing_rights SET cycle = DIV(block_level - 1, 8192) WHERE cycle is null and block_hash is null and block_level >= 1613825;
UPDATE tezos.endorsing_rights SET cycle = DIV(block_level - 1, 4096) WHERE cycle is null and block_hash is null and block_level < 1613825;

CREATE OR REPLACE VIEW tezos.baker_future_endorsement_view as
    select delegate, cycle, count(1) as count
    from tezos.future_endorsement_rights_view
    group by delegate, cycle;
