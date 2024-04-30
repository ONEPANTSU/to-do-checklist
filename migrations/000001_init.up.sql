create table users
(
    id            serial primary key,
    username      varchar(255) not null unique,
    email         varchar(255) not null unique,
    hashed_password varchar(256) not null
);

create table todo_lists
(
    id          serial primary key,
    title       varchar(256) not null,
    description varchar(512)
);

create table user_list
(
    id      serial primary key,
    user_id int not null references users (id) on delete cascade,
    list_id int not null references todo_lists (id) on delete cascade
);

create table todo_items
(
    id          serial primary key,
    title       varchar(256) not null,
    description varchar(512),
    completed   boolean      not null default false,
    list_id     int          not null references todo_lists (id) on delete cascade
);