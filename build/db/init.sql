CREATE EXTENSION pgcrypto;

create table posts
(
    id uuid default gen_random_uuid() not null
        constraint posts_pk
            primary key,
    score integer default 0 not null,
    views integer default 0 not null,
    title varchar(255) not null,
    url varchar(255) not null,
    "upvotePercentage" integer default 0 not null,
    created timestamp not null
);
