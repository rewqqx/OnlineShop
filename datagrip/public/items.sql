create table online_shop.items
(
    id          serial not null,
    item_name   varchar,
    price       numeric(2) default 1.00,
    description varchar,
    image_ids   integer[],
    constraint items_pkey primary key (id),
    constraint items_name unique (item_name),
    constraint items_price check (price > 0.00)
);

alter table online_shop.items
    owner to postgres;

