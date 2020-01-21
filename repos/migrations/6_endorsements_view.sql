CREATE OR REPLACE VIEW endorsements_view AS
SELECT delegate as baker, json_array_length(slots::json) as count, priority+1.0 as priority, block_level
FROM operations 
INNER JOIN blocks ON blocks.level = operations.level 
WHERE kind = 'endorsement';