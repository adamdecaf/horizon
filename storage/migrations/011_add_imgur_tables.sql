-- +migrate Up
-- +migrate StatementBegin
create table imgur_subreddit_galleries(
       subreddit_gallery_id varchar(36) primary key not null,
       subreddit varchar(128) not null,
       created_at timestamp(0) not null
);

create table imgur_subreddit_galleries_crawls(
       subreddit_crawl_id varchar(36) primary key not null,
       subreddit_gallery_id varchar(36) not null,
       successful boolean not null,
       crawled_at timestamp(0) not null
);

create table imgur_images(
       image_id varchar(36) primary key not null,
       imgur_image_id varchar(36) not null,
       storage_hash varchar(40) not null,
       crawled_at timestamp(0) not null
);
-- +migrate StatementEnd
