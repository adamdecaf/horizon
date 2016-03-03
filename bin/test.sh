#!/bin/bash

## Storage
export STORAGE_USER=horizon
export STORAGE_HOSTNAME=192.168.99.100
export STORAGE_PASSWORD=e06b4ed2b382f68
export STORAGE_PORT=5432

## Start the postgres image
docker-compose up -d postgres 2> /dev/null

exec go test -v ./...
