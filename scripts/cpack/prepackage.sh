#!/bin/sh
set -e
DIR=$(pwd)
rm -rf build && mkdir -p build/bin build/resources
cp ../../driver/sfprocessor $DIR/build/bin/.
cp -r ../service/systemd $DIR/build/service
cp -r ../../resources/pipelines $DIR/build/resources/.
cp -r ../../resources/policies $DIR/build/resources/.
cp -r ../../LICENSE.md $DIR/build/.
cp -r ../../README.md $DIR/build/.
cd $DIR

