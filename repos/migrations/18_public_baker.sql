-- +migrate Up

CREATE TABLE tezos.public_bakers(
	baker_name varchar,
	baker_offchain_registry_url varchar,
	baker_pays_from_accounts varchar[],
	delegate varchar,
	min_delegation integer,
	min_payout  integer,
	open_for_delegation boolean,
	over_delegation_threshold integer,
	payout_delay integer,
	payment_config_mask varchar,
	payout_frequency integer,
	reporter_account varchar[],
	split integer,
	subtract_payouts_less_than_min boolean,
	subtract_rewards_from_uninvited_delegation boolean,
	last_update_id integer,
	baker_charges_transaction_fee boolean,
	is_hidden boolean default false,
	PRIMARY KEY (delegate)
);

create index ix_operations_endorsements_level
  on tezos.operations (level)
  where ((kind)::text = 'endorsement'::text);

create index accounts_delegate_value_index
  on tezos.accounts (delegate_value)
  where delegate_value is not null and balance > 0;

-- +migrate Down
