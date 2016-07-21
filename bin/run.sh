#!/bin/bash

## Storage
export STORAGE_USER=horizon
export STORAGE_HOSTNAME=192.168.99.100
export STORAGE_PASSWORD=e06b4ed2b382f68
export STORAGE_PORT=5432

## Metrics
export STDOUT_REPORTING_ENABLED=no
export LIBRATO_REPORTING_ENABLED=yes
export LIBRATO_OWNER_EMAIL=adamkshannon@gmail.com
export LIBRATO_API_TOKEN=592e3708ffa92bdd1c8d70ea3d95fca3b9f7b11812f028ed054a4314fa7fb50f
export LIBRATO_INSTANCE_HOSTNAME=horizon2

## Reddit
export REDDIT_CRAWLER_ENABLED=no
export REDDIT_USERNAME=adamdecaf
export REDDIT_PASSWORD=""

## Twitter
export TWITTER_PUBLIC_CRAWLER_ENABLED=no
export TWITTER_USER_CRAWLER_ENABLED=no
export TWITTER_MENTION_PROCESSOR_ENABLED=no
export TWITTER_WORD_COUNT_REPROCESSOR_ENABLED=no
export TWITTER_CONSUMER_KEY=""
export TWITTER_CONSUMER_SECRET=""
export TWITTER_ACCESS_TOKEN=""

## Insert data
export INSERT_RAW_STATES=no
export INSERT_RAW_CITIES=no
export INSERT_RAW_COUNTRIES=no
export INSERT_HOSTNAMES=no

## Start the docker images
docker-compose up -d postgres 2> /dev/null

exec ./horizon
