CREATE TABLE users
(
    id            serial       not null unique,
    phone_code    varchar(10)  not null,
    phone_hash    varchar(255) not null unique,
    name          varchar(50),
    surname       varchar(50),
    email         varchar(50),
    birthday      date,
    value         int          default 0,
    message_key   varchar(255) not null unique
);

CREATE TABLE categories
(
    id   serial       not null unique,
    guid varchar(255) not null unique,
    name varchar(255) not null unique
);

CREATE TABLE sub_categories
(
    id          serial       not null unique,
    parent_guid varchar(255) references categories (guid) not null,
    guid        varchar(255)                              not null unique,
    name        varchar(255)                              not null unique
);

CREATE TABLE items
(
    id             serial        not null unique,
    cat_guid       varchar(255)  not null,
    sub_cat_guid   varchar(255)  default '',
    guid           varchar(255)  not null unique,
    name           varchar(255)  not null unique,
    description    text,
    thimbnails_pic varchar(255)  default ''
);

CREATE TABLE types
(
    id          serial                               not null unique,
    parent_guid varchar(255) references items (guid) not null,
    guid        varchar(255)                         not null unique,
    name        varchar(255),
    price       int,
    type_pic    varchar(255) default ''
);

CREATE TABLE actions
(
    id                 serial        not null unique,
    action_guid        varchar(255)  not null,
    action_name        varchar(255)  not null,
    action_start_date  date          not null default now(),
    action_expiry_date date          not null,
    action_picture     varchar(255)  default '',
    description        varchar(255)  default ''
);