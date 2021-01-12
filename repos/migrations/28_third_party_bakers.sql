create table tezos.third_party_bakers
(
    provider           varchar,
    number             integer,
    name               varchar,
    address            varchar,
    yield              double precision,
    staking_balance    numeric,
    fee                double precision,
    available_capacity numeric,
    efficiency         double precision,
    payout_accuracy    varchar
);

create index third_party_bakers_address_index
    on tezos.third_party_bakers (address);