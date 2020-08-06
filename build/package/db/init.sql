CREATE EXTENSION pgcrypto;

create table sessions
(
    id bigserial not null
        constraint sessions_pk
            primary key,
    user_id int not null
);


create table users
(
    id serial not null
        constraint users_pk
            primary key,
    password varchar(255) not null,
    admin boolean default false not null,
    login varchar(255) not null
);


create table votes
(
    post_id uuid not null,
    user_id int not null,
    vote smallint not null,
    constraint votes_pk
        primary key (post_id, user_id)
);

create unique index users_login_uindex
    on users (login);

INSERT INTO public.users (password, admin, login) VALUES ('$2a$08$MEXZpgo5BEloLRuicte/yObhVELHOdsS7W6nQdpT3w/lO8i1rylOu', true, 'admin');

create table posts
(
    id uuid default gen_random_uuid() not null
        constraint posts_pk
            primary key,
    views integer default 0 not null,
    title varchar(255) not null,
    url varchar(255) default ''::character varying not null,
    created timestamp default now() not null,
    author_id integer default 1 not null,
    category varchar(20) default 'news'::character varying not null,
    is_link boolean default false not null,
    text text default ''::text not null
);

INSERT INTO public.posts (id, views, title, url, created) VALUES ('3ebe76b7-a295-4526-be01-61f3a83c7d67', 2, 'First', 'example.com', '2020-08-02 02:59:43.000000');
INSERT INTO public.posts (id, views, title, url, created) VALUES ('b6e84291-e8b8-43a0-9160-aa1fba96ccb4', 54, '2nd funny title', 'utl', '2020-08-02 04:17:38.000000');

create index posts_category_index
    on posts (category);

create table comments
(
    id uuid default gen_random_uuid() not null
        constraint comments_pk
            primary key,
    post_id uuid not null,
    body text not null,
    author_id int not null
);

alter table comments
    add created timestamp without time zone default NOW() not null;

create index comments_post_id_index
    on comments (post_id);
