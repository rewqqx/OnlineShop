create table orders
(
    id               serial
        primary key,
    delivery_id      integer not null,
    payment_id       integer not null,
    total_price      integer,
    creation_date    timestamp default now(),
    ordered_item_ids integer[],
    user_id          integer
);

alter table orders
    owner to postgres;

