-- +migrate Up
create table normalized_urls(
       url_id varchar(36) primary key not null,
       url varchar(2048) not null,

       scheme varchar(10) not null,

       -- authority
       username varchar(128),
       password varchar(128),
       host varchar(1024) not null,
       port integer not null,

       path varchar(2048) not null,
       query varchar(2048),
       fragment varchar(1024)
);
