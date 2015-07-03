#!/bin/bash
set -e

function set_host_in_etc_hosts() {
    ip=$1
    host=$2

    if ! grep $ip /etc/hosts 1> /dev/null;
    then
        sudo /bin/bash -c "echo $ip $host >> /etc/hosts"
    fi
}

for i in 1 2 3 4
do
    set_host_in_etc_hosts "192.168.150.1$i" "horz-node-$i.horizon"
done
