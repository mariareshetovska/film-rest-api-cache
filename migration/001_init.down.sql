-- create database cinema_db;

\c cinema_db

CREATE TABLE if not exists films
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    release_year timestamp
);

CREATE TABLE if not exists response_time_log
(
    id serial not null unique,
    request varchar(255) not null unique,
    time_db int,
    time_redis int,
    time_memory int
);

insert into films (id, title, description, release_year) values (1, 'Academy Dinosaur', 'Epic Drama of a Feminist And a Mad Scientist who must Battle a Teacher in The Canadian Rockies', '01-01-2006');

insert into films (id, title, description, release_year) values (2, 'Ace Goldfinger', 'A Astounding Epistle of a Database Administrator And a Explorer who must Find a Car in Ancient China', '01-01-2006');

insert into films (id, title, description, release_year) values (3, 'Grosse Wonderful', 'A Epic Drama of a Cat And a Explorer who must Redeem a Moose in Australia', '01-01-2006');


