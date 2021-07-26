CREATE TABLE tezos.account_created_at
(
  account_id            varchar not null
    constraint account_created_at_pkey
      primary key,
  created_at          timestamp    not null
);

CREATE INDEX IF NOT EXISTS ix_account_created_at
    ON tezos.account_created_at USING btree
    (created_at DESC);

CREATE OR REPLACE FUNCTION insert_account_created_at()
  RETURNS TRIGGER LANGUAGE plpgsql
  AS $$
  BEGIN
      insert into tezos.account_created_at(account_id, created_at)
      VALUES (NEW.account_id, now());
  RETURN NULL;
  END $$;

CREATE TRIGGER account_created_at
  AFTER INSERT
  ON tezos.accounts
  FOR EACH ROW
EXECUTE PROCEDURE insert_account_created_at();

INSERT INTO tezos.account_created_at (account_id, created_at)
SELECT account_id , min(asof) FROM tezos.accounts_history GROUP BY account_id ON CONFLICT DO NOTHING;

CREATE VIEW account_list_view AS
SELECT accounts.*, created_at, blocks.timestamp last_active, aka.alias account_name, ka.alias as delegate_name
FROM "tezos"."accounts"
         inner join tezos.account_created_at act on accounts.account_id = act.account_id
         inner join tezos.blocks on accounts.block_id = blocks.hash
         left join tezos.known_addresses aka on accounts.account_id = aka.address
         left join tezos.known_addresses ka on accounts.delegate_value = ka.address;
