#!/bin/bash

set -e

## Setup Docker Machine

docker-machine --debug create \
               --engine-storage-driver overlay \
               --virtualbox-disk-size "30000" \
               --virtualbox-cpu-count 8 \
               --virtualbox-memory "5120" \
               --driver virtualbox \
               horizon-vm  2> /tmp/horizon-docker-machine-create.log
