create table online_shop.items
(
    id          serial not null,
    item_name   varchar,
    price       integer default 100,
    description varchar,
    image_ids   integer[],
    constraint items_pkey primary key (id),
    constraint items_name unique (item_name),
    constraint items_price check (price > 0)
);

alter table online_shop.items
    owner to postgres;

ALTER TABLE online_shop.items ADD COLUMN tag varchar;

