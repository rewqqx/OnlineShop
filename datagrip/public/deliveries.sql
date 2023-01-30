create table deliveries
(
    id            serial
        primary key,
    address_id    integer,
    status_id     integer,
    creation_date timestamp,
    update_date   timestamp,
    target_date   timestamp,
    type_id       integer
);

alter table deliveries
    owner to postgres;

