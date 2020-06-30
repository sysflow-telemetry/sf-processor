#!/bin/bash
# Usage: build.sh version release target-image:tag
#make build
docker build --build-arg BUILD_NUMBER=$2 --build-arg VERSION=$1 --build-arg RELEASE=$2 --target runtime -f Dockerfile.processor -t $3 .
