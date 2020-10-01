CREATE INDEX IF NOT EXISTS ix_operations_pkh_not_null
    ON operations USING btree
    (pkh COLLATE pg_catalog."default", kind COLLATE pg_catalog."default")
      WHERE pkh IS NOT NULL;
