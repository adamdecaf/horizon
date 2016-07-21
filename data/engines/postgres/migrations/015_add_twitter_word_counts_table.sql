create table twitter_hourly_word_counts (
       word varchar(32) not null,
       count smallint not null,
       hour timestamp not null
);

create index twitter_hourly_word_counts_hour_idx on twitter_hourly_word_counts using btree(hour);
