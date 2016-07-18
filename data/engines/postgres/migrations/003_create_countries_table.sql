-- +migrate Up
create table countries(
       country_id varchar(36) primary key not null,
       name varchar(128) not null,
       created_at timestamp not null default now()
);
