create table online_shop.order_items
(
    id            serial  not null,
    order_id      integer not null,
    item_id       integer not null,
    quantity      integer,
    default_price numeric(2) default 1.00,
    discount      integer    default 0,
    final_price   numeric(2) generated always as ( default_price * (100 - discount) / 100 ) stored,
    constraint order_items_default_price check (default_price > 0.00),
    constraint order_items_final_price check (final_price > 0.00),
    constraint order_items_pkey primary key (id),
    constraint order_items_item_fkey foreign key (item_id)
        references online_shop.items (id)
        on delete no action
        on update no action
        not deferrable,
    constraint order_items_order_fkey foreign key (order_id)
        references online_shop.orders (id)
        on delete no action
        on update no action
        not deferrable
);

alter table online_shop.order_items
    owner to postgres;

