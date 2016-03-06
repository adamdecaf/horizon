#!/bin/bash

docker-ip() {
  docker inspect --format '{{ .NetworkSettings.IPAddress }}' "$@" 2> /dev/null
}

GRAFANA_CONTAINER_ID=`docker ps | grep grafana | cut -f1`
GRAFANA_IP=`docker-ip $GRAFANA_CONTAINER_ID`

grafana-token() {
  curl -i -X GET http://admin:admin@$GRAFANA_IP:3000/
}	

GRAFANA_TOKEN=`grafana-token $GRAFANA_IP`
