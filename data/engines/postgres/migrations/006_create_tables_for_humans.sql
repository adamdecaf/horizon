-- +migrate Up
-- +migrate StatementBegin
create table humans(
       human_id varchar(36) primary key not null,
       created_at timestamp not null default now()
);

-- Names
create table humans_names(
       human_id varchar(36) not null,
       human_name_id varchar(36) not null,
       created_at timestamp not null default now()
);

create index humans_names_id_idx on humans_names using btree(human_id);
create index humans_names_idx on humans_names using btree(human_name_id);

create table humans_nicknames(
       human_id varchar(36) not null,
       human_nickname_id varchar(36) not null,
       created_at timestamp not null default now()
);

create index humans_nicknames_id_idx on humans_nicknames using btree(human_id);
create index humans_nicknames_idx on humans_nicknames using btree(human_nickname_id);

-- Employers
create table humans_employers(
       human_id varchar(36) not null,
       employer_id varchar(36) not null,
       created_at timestamp not null default now()
);

create index humans_employers_id_idx on humans_employers using btree(human_id);
create index humans_employers_idx on humans_employers using btree(employer_id);

-- Addresses
create table humans_addresses(
       human_id varchar(36) not null,
       normalized_address_id varchar(36) not null,
       created_at timestamp not null default now()
);

create index humans_addresses_id_idx on humans_addresses using btree(human_id);
create index humans_addresses_idx on humans_addresses using btree(normalized_address_id);
-- +migrate StatementEnd
