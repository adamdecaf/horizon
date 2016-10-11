#!/bin/bash

# Define all projects
all_projects=(geo human internet reddit twitter)

cmd=$1
proj=$2

projects=()
if [[ -z "$proj" ]];
then
    projects=${all_projects[*]}
else
    projects=("$proj")
fi

for project in ${projects[*]}
do
    wd=$(pwd)
    cd "$project"
    case "$cmd" in
        'build')
            echo "Building $project"
            CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o "$wd"/bin/horizon-"$project"-linux .
            GOOS=darwin GOARCH=386 go build -o "$wd"/bin/horizon-"$project"-osx .
        ;;
        'test')
            echo "Testing $project"
            go test -v ./...
        ;;
        'vet')
            echo "Vetting $project"
            go tool vet .
        ;;
    esac
    cd -
done
