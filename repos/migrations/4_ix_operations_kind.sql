CREATE INDEX IF NOT EXISTS ix_operations_kind
    ON operations USING btree
    (kind ASC NULLS LAST);

??