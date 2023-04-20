create table online_shop.payments
(
    id            serial
        primary key,
    payment_value integer,
    type_id       integer,
    status_id     integer,
    creation_date timestamp
);

alter table payments
    owner to postgres;

