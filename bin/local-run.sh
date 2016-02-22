#!/bin/bash

## Storage
export STORAGE_USER=horizon
export STORAGE_HOSTNAME=192.168.99.100
export STORAGE_PASSWORD=e06b4ed2b382f68
export STORAGE_PORT=5432

## Reddit
export REDDIT_CRAWLER_ENABLED=no
export REDDIT_USERNAME=adamdecaf
export REDDIT_PASSWORD=""

## Twitter
export TWITTER_CRAWLER_ENABLED=yes
export TWITTER_CONSUMER_KEY=""
export TWITTER_CONSUMER_SECRET=""
export TWITTER_ACCESS_TOKEN=""
export TWITTER_ACCESS_SECRET=""

## Insert data
export INSERT_RAW_STATES=no
export INSERT_RAW_CITIES=no
export INSERT_RAW_COUNTRIES=no

## Start the postgres image
docker-compose up -d postgres 2> /dev/null

exec ./horizon
