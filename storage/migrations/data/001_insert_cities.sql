-- +migrate Up
-- +migrate StatementBegin
create or replace function upsert_city(i varchar(36), n varchar(128)) returns void as
$$
begin

if not exists(select 1 from cities where city_id = i) then
  insert into cities (city_id, name, created_at) values (i, n, now());
else
  update cities set name = n where city_id = i;
end if;

end;
$$ language plpgsql;

-- Insert rows
-- select upsert_city('05b3d923-792c-4399-9d93-5c062eda7c01', 'Des Moines');

-- +migrate StatementEnd
