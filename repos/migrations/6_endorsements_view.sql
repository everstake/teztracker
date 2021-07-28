CREATE OR REPLACE VIEW tezos.endorsements_view AS
SELECT delegate as baker, number_of_slots, priority + 1.0 as priority, block_level
FROM tezos.operations
       INNER JOIN tezos.blocks ON blocks.level = operations.level
WHERE kind = 'endorsement';