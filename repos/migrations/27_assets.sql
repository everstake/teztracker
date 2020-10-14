-- Already on conseil precision -> scale
CREATE TABLE tezos.registered_tokens
(
  id            integer not null
    constraint registered_tokens_pkey
      primary key,
  name          text    not null,
  contract_type text    not null,
  account_id    text    not null,
  precision             int
);

CREATE TABLE tezos.asset_operations
(
    token_id  integer,
    operation_id integer constraint asset_operations_pkey
      primary key,
    operation_group_hash varchar,
    sender varchar,
    receiver varchar,
    amount numeric,
    type varchar,
    data varchar,
    timestamp timestamp
)

CREATE VIEW tezos.asset_info AS
SELECT registered_tokens.*, accounts.balance, operations.timestamp, operations.source
FROM tezos.registered_tokens
       LEFT JOIN tezos.accounts
                 on registered_tokens.account_id = accounts.account_id
       LEFT JOIN tezos.operations ON registered_tokens.account_id = originated_contracts;
