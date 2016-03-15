-- +migrate Up
-- +migrate StatementBegin
create table phone_numbers(
       phone_number_id varchar(36) primary key not null,
       number varchar(15) not null,
       country_code varchar(5) not null,
       created_at timestamp(0) not null
);

create table humans_phone_numbers(
       human_id varchar(36) not null,
       phone_number_id varchar(36) not null,
       created_at timestamp(0) not null,
       primary key (human_id, phone_number_id)
);

create table employers_phone_numbers(
       employer_id varchar(36) not null,
       phone_number_id varchar(36) not null,
       created_at timestamp(0) not null,
       primary key (employer_id, phone_number_id)
);
-- +migrate StatementEnd
