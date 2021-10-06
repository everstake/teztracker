CREATE TABLE tezos.nft_contracts
(
    id integer constraint nft_contracts_pkey primary key,
    name  varchar,
    contract_type  varchar,
    account_id  varchar,
    swap_contract varchar,
    description  varchar,
    ledger_big_map integer,
    tokens_big_map integer,
    operations_num integer,
    last_height integer,
    last_update_height integer
);

create unique index nft_contracts_account_id_uindex
	on nft_contracts (account_id);

create unique index nft_contracts_ledger_big_map_uindex
	on tezos.nft_contracts (ledger_big_map);

create unique index nft_contracts_tokens_big_map_uindex
	on tezos.nft_contracts (tokens_big_map);

CREATE VIEW tezos.nft_contracts_view AS
SELECT * FROM tezos.nft_contracts
    LEFT JOIN
    (SELECT big_map_id, count(1) nfts_number FROM tezos.big_map_contents GROUP BY big_map_id) s ON s.big_map_id = nft_contracts.tokens_big_map
ORDER BY id DESC;

CREATE TABLE tezos.nft_tokens
(
  contract_id integer,
  token_id integer,
  name varchar,
  description varchar,
  decimals integer,
  category varchar,
  amount integer,
  last_price numeric,
  issued_by varchar,
  is_for_sale boolean,
  ipfs_source varchar,
  created_at timestamp ,
  last_active_at timestamp
);

alter table tezos.nft_tokens
	add constraint nft_tokens_pk
		primary key (contract_id, token_id);

CREATE VIEW tezos.nft_tokens_ledger_view AS
select big_map_id, SPLIT_PART(key, ' ', 2) address, SPLIT_PART(key, ' ', 3) token_id, value from big_map_contents where value  :: integer > 0;