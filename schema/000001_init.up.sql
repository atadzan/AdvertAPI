CREATE TABLE adverts (
    id bigserial not null unique,
    title varchar(255) not null,
    description text,
    category varchar(255) not null,
    location varchar(255) not null,
    phone_number int,
    price int not null,
    publish_date timestamp not null,
    views int
);

CREATE TABLE images
(
    id bigserial not null unique,
    path varchar(255) not null,
    advert_id int REFERENCES adverts(id)
);
