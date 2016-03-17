-- +migrate Up
-- +migrate StatementBegin
create table employers(
       employer_id varchar(36) primary key not null,
       name varchar(250) not null,
       created_at timestamp not null default now()
);

create table employers_addresses(
       employer_id varchar(36) not null,
       normalized_address_id varchar(36) not null,
       created_at timestamp not null default now()
);

create index employers_addresses_id_idx on employers_addresses using btree(employer_id);
create index employers_addresses_normalized_address_id_idx on employers_addresses using btree(normalized_address_id);

create or replace function employer_name_search_trigger() returns trigger as
$$
begin
  new.name_tsvector := setweight(to_tsvector(coalesce(new.name)), 'A');
  return new;
end
$$ language plpgsql;

create or replace function add_employer_names_ts_vector() returns void as
$$
begin
if not exists(select 1 from pg_attribute where attrelid = (select oid from pg_class where relname = 'employers') and attname = 'name_tsvector') then
  alter table employers add column name_tsvector tsvector;
  create index employer_name_tsvector_idx on employers using gin(name_tsvector);

  update employers set name_tsvector = setweight(to_tsvector(coalesce(name)), 'A');

  create trigger employer_name_update_trigger before insert or update
  on employers for each row execute procedure employer_name_search_trigger();
end if;
end;
$$ language plpgsql;

select add_employer_names_ts_vector();
drop function add_employer_names_ts_vector();
-- +migrate StatementEnd
