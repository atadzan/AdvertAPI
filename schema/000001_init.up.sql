CREATE TABLE adverts (
    id bigserial not null unique primary key,
    title varchar(255) not null,
    description text,
    category varchar(255) not null,
    location varchar(255) not null,
    phone_number varchar(12),
    price int not null,
    publish_date timestamp not null,
    views int
);

CREATE TABLE images
(
    id bigserial not null unique,
    fname varchar(255) not null,
    fsize int8 not null,
    ftype varchar(255) not null,
    path varchar(255) not null,
    advert_id int4 REFERENCES adverts(id)
);

CREATE TABLE users
(
    id serial not null unique,
    username varchar(255) not null,
    password_hash varchar(255) not null,
    phone_number varchar(12) not null,
    created_at timestamp(0),
    updated_at timestamp(0),
    fav_list int[]
);