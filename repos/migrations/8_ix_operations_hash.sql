CREATE INDEX IF NOT EXISTS ix_operations_hash
    ON operations USING btree
    (operation_group_hash ASC NULLS LAST);

