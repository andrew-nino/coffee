CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE coffee_lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users_lists
(
    id      serial                                             not null unique,
    user_id int references users (id)   on delete cascade      not null,
    list_id int references coffee_lists (id) on delete cascade not null
);

CREATE TABLE coffee_items
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    done        boolean      not null default false
);


CREATE TABLE lists_items
(
    id      serial                                            not null unique,
    item_id int references coffee_items(id) on delete cascade not null,
    list_id int references coffee_lists(id) on delete cascade not null
);

CREATE TABLE client
(
    guid              varchar(255) not null unique,
    key_authorization varchar(255) not null unique,
    accaunt_name      varchar(255) not null unique
);

CREATE TABLE categories
(
    id   serial       not null unique,
    guid varchar(255) not null unique,
    name varchar(255) not null unique
);

CREATE TABLE sub_categories
(
    id   serial       not null unique,
    parent_guid varchar(255) references categories (guid) not null,
    guid        varchar(255)                              not null unique,
    name        varchar(255)                              not null unique
);

CREATE TABLE items
(
    id           serial                                        not null unique,
    cat_guid     varchar(255) references categories (guid)     not null,
    guid         varchar(255)                                  not null unique,
    name         varchar(255)                                  not null unique,
    description  varchar(255)
);

CREATE TABLE types
(
    id          serial                               not null unique,
    parent_guid varchar(255) references items (guid) not null,
    guid        varchar(255)                         not null unique,
    name        varchar(255),
    price       int,
    type_pic    varchar(255)
);