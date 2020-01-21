CREATE INDEX IF NOT EXISTS ix_opeations_delegate_kind
    ON operations USING btree
    (delegate COLLATE pg_catalog."default", kind COLLATE pg_catalog."default")
      WHERE delegate IS NOT NULL;
