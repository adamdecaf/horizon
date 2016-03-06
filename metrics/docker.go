package metrics

// todo
// Add docker golang client to grab containers and ips



// docker-ip() {
//   docker inspect --format '{{ .NetworkSettings.IPAddress }}' "$@" 2> /dev/null
// }

// GRAFANA_CONTAINER_ID=`docker ps | grep grafana | cut -f1`
// GRAFANA_IP=`docker-ip $GRAFANA_CONTAINER_ID`

// grafana-token() {
//   curl -i -X GET http://admin:admin@$GRAFANA_IP:3000/
// }

// GRAFANA_TOKEN=`grafana-token $GRAFANA_IP`

// apt-get install -y influxdb-client > /dev/null

// docker-ip() {
//   docker inspect --format '{{ .NetworkSettings.IPAddress }}' "$@" 2> /dev/null
// }

// INFLUXDB_CONTAINER_ID=`docker ps | grep influxdb1 | cut -f1`
// INFLUX_IP=`docker-ip $INFLUX_CONTAINER_ID`

// influx -execute 'show databases'
// influx -execute 'create database "metrics"' | tee > /dev/null
// influx -execute 'create user metrics'
// # influx -execute 'grant all privileges to metrics'
