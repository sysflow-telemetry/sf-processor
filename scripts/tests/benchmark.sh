#!/bin/bash

tm=timeout
if ! command -v $tm &> /dev/null
then
    tm=gtimeout
fi

DURATION=$1
CONFIG=$2
TRACES=$3

for n in {1..10}
do    
    export POLICYENGINE_BENCH_RULESSAMPLESIZE=$(( $n*5 ))    
    echo "Benchmarking with $POLICYENGINE_BENCH_RULESSAMPLESIZE rules"        
    ( $tm $DURATION ../../driver/sfprocessor -perflog -log=quiet -config=$CONFIG -driver=file $TRACES > rate_$POLICYENGINE_BENCH_RULESSAMPLESIZE.out ) &
    sleep $(( $DURATION + 30 ))
done