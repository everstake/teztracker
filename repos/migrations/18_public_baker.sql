CREATE TABLE tezos.public_bakers(
	baker_charges_transaction_fee boolean,
	baker_name varchar,
	baker_offchain_registry_url varchar,
	baker_pays_from_accounts varchar[],
	delegate varchar  constraint public_bakers_pkey primary key,
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
	is_hidden boolean default false,
	media text
);