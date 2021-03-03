create table tezos.daily_stats
(
    key   varchar(255) not null,
    date  date         not null,
    value bigint       not null
);

create
index daily_stats_date_index
    on tezos.daily_stats (date);

create
unique index daily_stats_key_date_uindex
    on tezos.daily_stats (key, date);

create table tezos.storage
(
    key   varchar(255) not null
        constraint storage_pk
            primary key,
    value text         not null
);

create
unique index storage_key_uindex
    on tezos.storage (key);