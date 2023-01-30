create table users
(
    id            serial
        primary key,
    user_name     varchar,
    user_surname  varchar,
    birthdate     timestamp,
    password_hash varchar,
    mail          varchar,
    role_id       integer
);

alter table users
    owner to postgres;

