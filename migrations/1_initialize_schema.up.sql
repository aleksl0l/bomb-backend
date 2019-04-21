create table if not exists "user"
(
    id            serial primary key,
    username      text not null unique,
    password_hash text not null
);

create table if not exists "profile"
(
    id    serial primary key,
    email text unique
);

create table if not exists "token"
(
    id      serial primary key,
    token   text not null
);

alter table "profile"
    add column user_id integer references "user" (id) on delete cascade;

alter table "token"
    add column user_id integer unique references "user" (id) on delete cascade;