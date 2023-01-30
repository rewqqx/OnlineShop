create table items
(
    id          serial
        primary key,
    item_name   varchar,
    price       integer,
    description varchar,
    image_ids   integer[]
);

alter table items
    owner to postgres;

