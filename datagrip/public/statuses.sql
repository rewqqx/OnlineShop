create table online_shop.order_statuses
(
    id   serial not null,
    status_name varchar,
    constraint order_statuses_pkey primary key (id),
    constraint order_statuses_name unique(status_name)
);

alter table online_shop.order_statuses
    owner to postgres;

insert into online_shop.order_statuses (status_name) values
                                                   ('Новый'),
                                                   ('Выдан'),
                                                   ('Отменен'),
                                                   ('Формируется'),
                                                   ('Готов к выдаче'),
                                                   ('Передан курьеру'),
                                                   ('Оплачен, готов к выдаче')