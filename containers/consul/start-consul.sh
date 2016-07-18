#!/bin/sh

COMMAND="/bin/consul agent --data-dir $CONSUL_DATA_DIR -server"

if [[ ! -z $CONSUL_SEED_NODE ]];
then
    COMMAND=$COMMAND" -join $CONSUL_SEED_NODE"
else
    # If there is no seed, assume it can be defacto master
    COMMAND=$COMMAND" -bootstrap-expect 1"
fi

COMMAND=$COMMAND" -node $CONSUL_NODE_NAME"

exec $COMMAND
