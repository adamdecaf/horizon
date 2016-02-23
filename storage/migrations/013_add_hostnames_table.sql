-- +migrate Up
create table hostnames(
       hostname_id varchar(36) primary key not null,
       value varchar(256) not null
);
