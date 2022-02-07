#!/bin/sh
set -e

DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $DIR && rm -rf build && mkdir -p build
cp -a $DIR/../../bin $DIR/build/bin
cp -a $DIR/../service/systemd $DIR/build/service
mkdir -p $DIR/build/resources
cp -a $DIR/../../resources/pipelines $DIR/build/resources/.
cp -a $DIR/../../resources/policies $DIR/build/resources/.
cp $DIR/../../LICENSE.md $DIR/build/.
cp $DIR/../../README.md $DIR/build/.
cd $DIR

