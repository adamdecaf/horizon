#!/bin/bash
set -e
ACTION=$1

case "$ACTION" in
    install)
        # Install the latest docker, ubuntu 14.04 only has 1.0.1 by default. Yes, on this date in history (2015-04-12)
        # ubuntu 14.04 is 5 minor versions behind on a package, shocking I know..
        ## FYI: It might be better to install via "exec curl -s https://get.docker.com | sh" because we had to switch to using aufs
        if [[ ! -e /etc/apt/sources.list.d/docker.list ]];
        then
            echo "=== Adding docker apt repository"
            sudo sh -c "echo deb https://get.docker.io/ubuntu docker main > /etc/apt/sources.list.d/docker.list"
            sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 36A1D7869245C8950F966E92D8576A8BA88D21E9
        fi

        # Tell ubuntu to install postgres 9.4, ubuntu 14.04 is stuck on 9.3
        if [[ ! -e /etc/apt/sources.list.d/pgdg.list ]];
        then
            echo "=== Adding postgres 9.4 apt repository"
            sudo sh -c "echo deb http://apt.postgresql.org/pub/repos/apt/ trusty-pgdg main > /etc/apt/sources.list.d/pgdg.list"
            wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add - sudo apt-get update
        fi

        echo "== Installing Base Packages"
        sudo apt-get update
        sudo apt-get install -y -q --force-yes curl dstat emacs24 git htop lxc-docker ntp ruby runit postgresql-client-9.4 screen sysstat vim wget

        echo "== Staring docker dameon"
        # Inspired from https://github.com/Banno/vagrant-mesos, under apache 2 license

        if [[ ! -h "/service" ]]
        then
            sudo ln -s /etc/service /service
        fi

        sudo mkdir -p /etc/sv/docker/log
        cat <<EOF >/etc/sv/docker/log/run
#!/bin/sh
exec svlogd -tt /var/log/docker
EOF
        chmod 755 /etc/sv/docker/log/run
        mkdir -p /var/log/docker
        cat <<EOF > /etc/sv/docker/run
#!/bin/sh
exec 2>&1
exec /usr/bin/docker -d --host=unix:///var/run/docker.sock --storage-driver=aufs
EOF
        chmod 755 /etc/sv/docker/run

        if [[ ! -h "/service/docker" ]]
        then
            ln -s /etc/sv/docker /service/
        fi
    ;;
    upgrade)
        echo "== Upgradaing Base Packages"
        apt-get update > /tmp/apt-update.log
        apt-get upgrade -y
    ;;
esac
