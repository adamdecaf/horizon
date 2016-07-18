-- +migrate Up
-- +migrate StatementBegin
create table states(
       state_id varchar(36) primary key not null,
       name varchar(128) not null,
       abbreviation varchar(5) not null,
       created_at timestamp not null default now()
);

alter table cities add column state_id varchar(36) not null;
-- +migrate StatementEnd
