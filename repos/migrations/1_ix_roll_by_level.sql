CREATE INDEX CONCURRENTLY IF NOT EXISTS  roll_by_level
    ON rolls USING btree
    (block_level DESC NULLS LAST);