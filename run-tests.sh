#!/bin/bash
## Storage
export STORAGE_USER=horizon
export STORAGE_USERNAME=192.168.59.103
export STORAGE_PASSWORD=e06b4ed2b382f68

## Start the postgres image
docker-compose up -d postgres 2> /dev/null

exec go test -v ./...
