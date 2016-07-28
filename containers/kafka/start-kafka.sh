#!/bin/bash

if [[ "$DELAY_KAFKA_STARTUP" -gt 0 ]]; then
  echo "Delaying Kafka startup by $DELAY_KAFKA_STARTUP seconds..."
  sleep ${DELAY_KAFKA_STARTUP}s
fi

export KAFKA_ADVERTISED_HOST_NAME=${KAFKA_ADVERTISED_HOST_NAME:-$HOST}
export KAFKA_ADVERTISED_HOST_NAME=${KAFKA_ADVERTISED_HOST_NAME:-$(ifconfig eth0 | grep -Eo '[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}' | head -n1)}
export KAFKA_ZOOKEEPER_CONNECT=${KAFKA_ZOOKEEPER_CONNECT:-$(echo $ZOOKEEPER_PORT_2181_TCP | sed -e 's|tcp://||')$KAFKA_ZOOKEEPER_CHROOT_PATH}

if [[ -n "$KAFKA_HEAP_OPTS" ]]; then
    sed -r -i "s/^(export KAFKA_HEAP_OPTS)=\"(.*)\"/\1=\"$KAFKA_HEAP_OPTS\"/g" $KAFKA_HOME/bin/kafka-server-start.sh
    unset KAFKA_HEAP_OPTS
fi

for VAR in `env`
do
  if [[ $VAR =~ ^KAFKA_ && ! $VAR =~ ^KAFKA_HOME ]]; then
    kafka_name=`echo "$VAR" | sed -r "s/KAFKA_(.*)=.*/\1/g" | tr '[:upper:]' '[:lower:]' | tr _ .`
    env_var=`echo "$VAR" | sed -r "s/(.*)=.*/\1/g"`
    if egrep -q "(^|^#)$kafka_name" $KAFKA_HOME/config/server.properties; then
        sed -r -i "s@(^|^#)($kafka_name)=(.*)@\2=${!env_var}@g" $KAFKA_HOME/config/server.properties #note that no config values may contain an '@' char
    else
        echo "$kafka_name=${!env_var}" >> $KAFKA_HOME/config/server.properties
    fi
  fi
done

env | sort
cat $KAFKA_HOME/config/server.properties

exec $KAFKA_HOME/bin/kafka-server-start.sh $KAFKA_HOME/config/server.properties
