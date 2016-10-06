-- +migrate Up
create table cities(
       city_id varchar(36) primary key not null,
       name varchar(128) not null,
       created_at timestamp not null default now()
);
