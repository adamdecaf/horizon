-- +migrate Up
create table normalized_addresses(
       address_id varchar(36) primary key not null,
       address varchar(2048) not null,

       -- parsed parts
       city_id varchar(36),
       postal_code smallint,
       country_id varchar(36)
);
