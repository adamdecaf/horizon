#!/bin/bash
set -e

echo "Getting deps for horizon"

# Go deps
for dep in ChimeraCoder/anaconda jzelinskie/geddit lib/pq rubenv/sql-migrate satori/go.uuid mvdan/xurls
do
    echo "Updating $dep"
    cd $GOPATH/src/github.com/$dep && \
        git pull origin master && \
        go build .
done

echo "Finihsed getting deps"
