#!/bin/bash
set -e

echo "Getting deps for horizon"

# Go deps
for dep in adamdecaf/go-whois ChimeraCoder/anaconda jzelinskie/geddit ivpusic/grpool lib/pq rubenv/sql-migrate satori/go.uuid mvdan/xurls
do
    echo "Updating $dep"
    cd $GOPATH/src/github.com/$dep && \
        git pull origin master && \
        go build .
done

echo "Finihsed getting deps"
