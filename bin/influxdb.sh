#!/bin/bash

apt-get install -y influxdb-client > /dev/null

docker-ip() {
  docker inspect --format '{{ .NetworkSettings.IPAddress }}' "$@" 2> /dev/null
}

INFLUXDB_CONTAINER_ID=`docker ps | grep influxdb1 | cut -f1`
INFLUX_IP=`docker-ip $INFLUX_CONTAINER_ID`

influx -execute 'show databases'
influx -execute 'create database "metrics"' | tee > /dev/null
influx -execute 'create user metrics'
# influx -execute 'grant all privileges to metrics'
