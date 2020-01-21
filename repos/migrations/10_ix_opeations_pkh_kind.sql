
CREATE INDEX IF NOT EXISTS ix_opeations_pkh_kind
    ON operations USING btree
    (pkh COLLATE pg_catalog."default", kind COLLATE pg_catalog."default")
      WHERE pkh IS NOT NULL;
