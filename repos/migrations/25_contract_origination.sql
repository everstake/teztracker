CREATE INDEX IF NOT EXISTS ix_operations_originated_contract
    ON tezos.operations USING btree
    (originated_contracts)
      WHERE originated_contracts IS NOT NULL;