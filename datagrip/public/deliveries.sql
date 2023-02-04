create table online_shop.deliveries
(
    id          serial  not null,
    order_id    integer not null,
    address_id  integer not null,
    target_date timestamp,
    type_id     integer,
    constraint deliveries_pkey primary key (id),
    constraint deliveries_order_fkey foreign key (order_id)
        references online_shop.orders (id)
        on delete no action
        on update no action
        not deferrable
);

alter table online_shop.deliveries
    owner to postgres;

