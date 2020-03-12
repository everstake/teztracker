-- +migrate Up

CREATE OR REPLACE VIEW tezos.endorsements_view AS
SELECT delegate as baker, json_array_length(slots::json) as count, priority+1.0 as priority, block_level
FROM tezos.operations
INNER JOIN tezos.blocks ON blocks.level = operations.level
WHERE kind = 'endorsement';

-- +migrate Down