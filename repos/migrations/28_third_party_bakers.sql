CREATE TABLE tezos.third_party_bakers
(
    provider  varchar,
    number integer,
    name varchar,
    address varchar,
    yield float,
    staking_balance numeric,
    fee float,
    available_capacity numeric,
    efficiency float,
    payout_accuracy varchar
);