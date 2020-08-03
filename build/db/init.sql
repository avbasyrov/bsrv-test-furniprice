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

INSERT INTO public.users (password, admin, login) VALUES ('12345678', true, 'basyrov');

create table posts
(
    id uuid default gen_random_uuid() not null
        constraint posts_pk
            primary key,
    views integer default 0 not null,
    title varchar(255) not null,
    url varchar(255) not null,
    created timestamp not null,
    author_id integer default 1 not null,
    category varchar(20) default 'news'::character varying not null,
    is_link boolean default false not null,
    text text default ''::text not null
);

INSERT INTO public.posts (id, score, views, title, url, upvote, created) VALUES ('3ebe76b7-a295-4526-be01-61f3a83c7d67', 1, 2, 'First', 'example.com', 33, '2020-08-02 02:59:43.000000');
INSERT INTO public.posts (id, score, views, title, url, upvote, created) VALUES ('b6e84291-e8b8-43a0-9160-aa1fba96ccb4', 13, 54, '2nd funny title', 'utl', 11, '2020-08-02 04:17:38.000000');

create index posts_category_index
    on posts (category);
