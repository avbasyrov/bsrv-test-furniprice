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
    upvote integer default 0 not null,
    created timestamp not null
);

create table categories
(
    id serial not null
        constraint categories_pk
            primary key,
    title varchar(255) not null
);

INSERT INTO public.categories (id, title) VALUES (DEFAULT, 'music');
INSERT INTO public.categories (id, title) VALUES (DEFAULT, 'programming');
INSERT INTO public.categories (id, title) VALUES (DEFAULT, 'funny');
INSERT INTO public.categories (id, title) VALUES (DEFAULT, 'news');
INSERT INTO public.categories (id, title) VALUES (DEFAULT, 'fashion');
INSERT INTO public.categories (id, title) VALUES (DEFAULT, 'videos');

alter table posts
    add category_id int;

alter table posts
    add constraint posts_categories_id_fk
        foreign key (category_id) references categories
            on update restrict on delete restrict;

alter table posts alter column category_id set not null;

INSERT INTO public.posts (id, score, views, title, url, upvote, created, category_id) VALUES ('3ebe76b7-a295-4526-be01-61f3a83c7d67', 1, 2, 'First', 'example.com', 33, '2020-08-02 02:59:43.000000', 3);
INSERT INTO public.posts (id, score, views, title, url, upvote, created, category_id) VALUES ('b6e84291-e8b8-43a0-9160-aa1fba96ccb4', 13, 54, '2nd funny title', 'utl', 11, '2020-08-02 04:17:38.000000', 3);
