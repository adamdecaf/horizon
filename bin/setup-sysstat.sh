#!/bin/bash
set -e

if grep -v "ENABLED=\"true\"" /etc/default/sysstat;
then
    echo "== Installing and setting up sysstat"
    apt-get install -y -q sysstat
    sed -i 's/ENABLED="false"/ENABLED="true"/' /etc/default/sysstat
    sed -i 's/5-55\/10/*/' /etc/cron.d/sysstat
    /etc/init.d/sysstat restart
fi
