-- +migrate Up
-- +migrate StatementBegin
create table emails(
       email_id varchar(36) primary key not null,
       email varchar(128) not null,
       domain varchar(128) not null,
       created_at timestamp(0) not null
);

create table humans_emails(
       human_id varchar(36) not null,
       email_id varchar(36) not null,
       created_at timestamp(0) not null,
       primary key (human_id, email_id)
);

create table employers_emails(
       employer_id varchar(36) not null,
       email_id varchar(36) not null,
       created_at timestamp(0) not null,
       primary key (employer_id, email_id)
);
-- +migrate StatementEnd
