#!/bin/bash

tm=timeout
if ! command -v $tm &> /dev/null
then
    tm=gtimeout
fi

DURATION=$1
CONFIG=$2
TRACES=$3
OUTDIR=$4

mkdir -p $OUTDIR

for n in {1..10}
do    
    export POLICYENGINE_BENCH_RULESETSIZE=$(( $n*5 ))    
    echo "Benchmarking with $POLICYENGINE_BENCH_RULESETSIZE rules"        
    ( $tm $DURATION ../../driver/sfprocessor -perflog -log=quiet -config=$CONFIG -driver=file $TRACES > $OUTDIR/rate_$POLICYENGINE_BENCH_RULESETSIZE.out ) &
    sleep 90
done