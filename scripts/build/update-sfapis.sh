#!/bin/bash
# Usage: update-sfapis.sh [tag|branch]
DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
for gomod in $(find $DIR/../.. -name go.mod); do    
    match=$(grep -w github.com/sysflow-telemetry/sf-apis/go $gomod)
    if [ -n "$match" ]; then    
        echo "Updating $gomod"
        CDIR=$(pwd)
        cd $(dirname $gomod)
        go get -u github.com/sysflow-telemetry/sf-apis/go@$1
        go mod tidy
        cd $CDIR
    fi
done