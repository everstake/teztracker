CREATE TABLE tezos.asset_operations
(
    block_level integer,
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


CREATE UNIQUE index registered_tokens_account_id_uindex
	ON tezos.registered_tokens (account_id);

alter table tezos.asset_operations
	add block_level int default 0 not null;

CREATE INDEX asset_operations_token_id_index
	ON tezos.asset_operations (token_id);

create index asset_operations_timestamp_index
	on tezos.asset_operations (timestamp desc);

alter table tezos.asset_operations
	add block_level int default 0 not null;

create index asset_operations_block_level_index
	on tezos.asset_operations (block_level desc);
