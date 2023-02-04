create table online_shop.users
(
    id              serial not null,
    user_name       varchar,
    user_surname    varchar,
    user_patronymic varchar,
    phone           varchar not null,
    birthdate       timestamp,
    password_hash   varchar,
    mail            varchar,
    role_id         integer,
    constraint users_pkey primary key (id),
    constraint users_phone unique (phone),
    constraint users_mail unique (mail)
);

alter table online_shop.users
    owner to postgres;

