-- +migrate Up
-- +migrate StatementBegin
create table twitter_users(
       twitter_user_id varchar(36) primary key not null,
       name varchar(256) not null,
       screen_name varchar(256) not null,
       created_at timestamp(0) not null
);

create table twitter_tweets(
       tweet_id varchar(36) primary key not null,
       twitter_user_id varchar(36) not null,
       text varchar(160) not null,
       created_at timestamp(0) not null
);
create index twitter_tweets_user_id_idx on twitter_tweets using btree(twitter_user_id);

create table twitter_tweet_urls(
       twitter_tweet_url_id varchar(36) primary key not null,
       tweet_id varchar(36) not null,
       url varchar(1024) not null
);
create index twitter_urls_tweet_id_idx on twitter_tweet_urls using btree(twitter_tweet_url_id);
-- +migrate StatementEnd
