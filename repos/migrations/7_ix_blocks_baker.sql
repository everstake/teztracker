CREATE INDEX IF NOT EXISTS ix_blocks_baker
    ON blocks USING btree
    (baker ASC NULLS LAST);
