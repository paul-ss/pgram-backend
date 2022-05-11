-- create database dbdb
-- 	with owner api_user
-- 	encoding 'utf8'
-- 	LC_COLLATE = 'C'
--     LC_CTYPE = 'C'
--     TEMPLATE template0
-- ;

CREATE EXTENSION citext;


create table if not exists users (
    id serial not null primary key,
    nickname citext not null unique,
    first_name text default null,
    last_name text default null,
    about text default null,
    email text not null unique,
    password text not null,
    image text default null
);


create table if not exists groups (
    id serial not null primary key,
    name text not null
);


create table if not exists posts (
    id bigserial not null primary key,
    user_id integer not null,
    group_id integer default null,
    content text default null,
    created timestamp with time zone default now(),
    image text default null,

    -- trigger
    likes integer default 0,
    comments integer default 0,

    foreign key (user_id) references users(id),
    foreign key (group_id) references groups(id)
);

