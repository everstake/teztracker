create table tezos.users
(
    email      varchar                 not null,
    account_id varchar(36)             not null
        constraint users_pk
            primary key,
    created_at timestamp default now() not null
);

create table tezos.user_addresses
(
    account_id          varchar(36) not null,
    delegations_enabled boolean,
    transfers_enabled   boolean,
    address             varchar
);

create index user_addresses_account_id_index
    on tezos.user_addresses (account_id);

create index user_addresses_address_index
    on tezos.user_addresses (address);


create table tezos.user_notes
(
    id         serial      not null
        constraint user_notes_pk
            primary key,
    account_id varchar(36) not null,
    alias      varchar,
    text       varchar     not null
);

create
    index user_notes_user_id_index
    on tezos.user_notes (account_id);

