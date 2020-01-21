CREATE INDEX IF NOT EXISTS ix_opeations_public_key_kind
    ON operations USING btree
    (public_key COLLATE pg_catalog."default", kind COLLATE pg_catalog."default")
      WHERE public_key IS NOT NULL;