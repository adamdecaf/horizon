-- horizon postgres queries

-- find duplicated cities
select name, state_id
from cities
group by name, state_id
having count(name) > 1
;

-- remove cities that are invalid
delete from cities
where length(city_id) <> 36
;

-- remove states that are invalid
delete from states
where length(state_id) <> 36
;

-- most tweeting user
select u.id, u.name, count(t.twitter_user_id)
from twitter_users as u
inner join twitter_tweets t
on u.twitter_user_id = t.twitter_user_id
group by u.name, t.twitter_user_id
order by count(t.twitter_user_id) desc
-- having count(t.twitter_user_id) > 1
limit 10
;

-- find more shared urls
select count(url), url
from twitter_tweet_urls
group by url
order by count(url) desc
limit 20
;

-- tweets by minute for the last hour
select date_trunc('minute', created_at), count(tweet_id)
from twitter_tweets
where created_at > now() - interval '1 hour'
group by date_trunc('minute', created_at)
order by date_trunc('minute', created_at) asc
;
