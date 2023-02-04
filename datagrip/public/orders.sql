create table online_shop.orders
(
    id                serial  not null,
    display_number    varchar not null,
    user_id           integer not null,
    status_id         integer    default 1,
    cancel_reason     varchar,
    payment_id        integer not null,
    total_price       numeric(2) default 1.00,
    creation_date     timestamp  default now(),
    modification_date timestamp  default now(),
    constraint orders_pkey primary key (id),
    constraint orders_display_num unique (display_number),
    constraint orders_status_fk foreign key (status_id)
        references online_shop.order_statuses (id)
        on delete no action
        on update no action
        not deferrable,
    constraint orders_user_fk foreign key (user_id)
        references online_shop.users (id)
        on delete no action
        on update no action
        not deferrable
);

alter table online_shop.orders
    owner to postgres;


