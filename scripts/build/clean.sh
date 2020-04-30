#!/bin/bash
# Usage: clean.sh target-image:tag
docker rmi $1
docker image prune -af
make clean
