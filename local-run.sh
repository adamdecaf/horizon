#!/bin/bash

## Storage
export STORAGE_USER=horizon
export STORAGE_USERNAME=192.168.59.103
export STORAGE_PASSWORD=e06b4ed2b382f68

## Reddit
export REDDIT_CRAWLER_ENABLED=no
export REDDIT_USERNAME=adamdecaf
export REDDIT_PASSWORD=""

## Insert data
## todo: Add a 'check' value which would insert only if there's no data in the table
export INSERT_RAW_STATES=no
export INSERT_RAW_CITIES=no
export INSERT_RAW_COUNTRIES=no

## Start the postgres image
docker-compose up -d postgres 2> /dev/null

exec ./horizon
