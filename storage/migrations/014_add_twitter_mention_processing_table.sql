-- +migrate Up
-- +migrate StatementBegin
create table twitter_mention_processing_runs(
       twitter_mention_processing_id varchar(36) primary key not null,
       range_start timestamp(0) not null,
       range_end timestamp(0) not null,
       created_at timestamp(0) not null
);
-- +migrate StatementEnd
