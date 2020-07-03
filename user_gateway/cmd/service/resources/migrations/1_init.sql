-- +migrate Up
create table "user" (
    id uuid,
    username varchar (32) unique,
    password varchar (128),
    primary key (id)
);