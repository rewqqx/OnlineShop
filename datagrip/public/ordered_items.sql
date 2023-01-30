create table ordered_items
(
    id            serial
        primary key,
    item_id       integer not null,
    quantity      integer,
    default_price integer,
    discount      integer,
    final_price   integer
);

alter table ordered_items
    owner to postgres;

