create table tezos.users
(
    email      varchar                 not null,
    account_id varchar                 not null,
    created_at timestamp default now() not null
);


create table tezos.user_addresses
(
    account_id          varchar not null,
    delegations_enabled boolean,
    transfers_enabled   boolean
);

create
    unique index user_address_uindex
    on tezos.users (account_id);


create table tezos.user_notes
(
    id         serial  not null
        constraint user_notes_pk
            primary key,
    account_id varchar not null,
    alias      varchar,
    text       varchar not null
);

create
    index user_notes_user_id_index
    on tezos.user_notes (account_id);

