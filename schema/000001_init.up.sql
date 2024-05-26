CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE books
(
    id          serial        not null unique,
    id_name     varchar(255)  not null unique,
    book_name   varchar(255)  not null,
    book_author varchar(255)  not null,
    book_genres varchar(255)  not null,
    book_image  varchar(255)  not null,
    description varchar(1023) not null
);

CREATE TABLE genres
(
    id    serial       not null unique,
    genre varchar(255) not null unique,
);