create table tezos.users
(
    email      varchar                 not null,
    username   varchar                 not null,
    account_id varchar(36)             not null
        constraint users_pk
            primary key,
    verified   boolean                 not null,
    created_at timestamp default now() not null
);

create table tezos.user_addresses
(
    account_id            varchar(36) not null,
    delegations_enabled   boolean not null,
    in_transfers_enabled  boolean not null,
    out_transfers_enabled boolean not null,
    address               varchar
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


create table tezos.email_verifications
(
    account_id varchar                 not null,
    email      varchar                 not null,
    token      varchar                 not null,
    verified   boolean   default false not null,
    created_at timestamp default now() not null,
    sent       boolean   default false not null
);

create index email_verifications_account_id_index
    on tezos.email_verifications (account_id);

create unique index email_verifications_token_uindex
    on tezos.email_verifications (token);

create index email_verifications_sent_index
    on tezos.email_verifications (sent);


