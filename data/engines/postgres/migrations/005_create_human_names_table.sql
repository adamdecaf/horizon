-- +migrate Up
-- +migrate StatementBegin
-- Names
create table human_names(
       name_id varchar(36) primary key not null,
       name varchar(250) not null,
       created_at timestamp not null default now()
);

create or replace function human_name_search_trigger() returns trigger as
$$
begin
  new.name_tsvector := setweight(to_tsvector(coalesce(new.name)), 'A');
  return new;
end
$$ language plpgsql;

create or replace function add_human_names_ts_vector() returns void as
$$
begin
if not exists(select 1 from pg_attribute where attrelid = (select oid from pg_class where relname = 'human_names') and attname = 'name_tsvector') then
  alter table human_names add column name_tsvector tsvector;
  create index human_name_tsvector_idx on human_names using gin(name_tsvector);

  update human_names set name_tsvector = setweight(to_tsvector(coalesce(name)), 'A');

  create trigger human_name_update_trigger before insert or update
  on human_names for each row execute procedure human_name_search_trigger();
end if;
end;
$$ language plpgsql;

select add_human_names_ts_vector();
drop function add_human_names_ts_vector();

-- Nicknames
create table human_nicknames(
       nickname_id varchar(36) primary key not null,
       nickname varchar(250) not null,
       created_at timestamp not null default now()
);

create or replace function human_nickname_search_trigger() returns trigger as
$$
begin
  new.nickname_tsvector := setweight(to_tsvector(coalesce(new.nickname)), 'A');
  return new;
end
$$ language plpgsql;

create or replace function add_human_nicknames_ts_vector() returns void as
$$
begin
if not exists(select 1 from pg_attribute where attrelid = (select oid from pg_class where relname = 'human_nicknames') and attname = 'nickname_tsvector') then
  alter table human_nicknames add column nickname_tsvector tsvector;
  create index human_nickname_tsvector_idx on human_nicknames using gin(nickname_tsvector);

  update human_nicknames set nickname_tsvector = setweight(to_tsvector(coalesce(nickname)), 'A');

  create trigger human_nickname_update_trigger before insert or update
  on human_nicknames for each row execute procedure human_nickname_search_trigger();
end if;
end;
$$ language plpgsql;

select add_human_nicknames_ts_vector();
drop function add_human_nicknames_ts_vector();
-- +migrate StatementEnd
