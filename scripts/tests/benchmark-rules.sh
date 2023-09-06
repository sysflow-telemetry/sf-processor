#!/bin/bash

tm=timeout
if ! command -v $tm &> /dev/null
then
    tm=gtimeout
fi

DURATION=$1
CONFIG=$2
RULES=$3
TRACES=$4
OUTDIR=$5

mkdir -p $OUTDIR

n=0
while [ "$n" -lt "$RULES" ]; do
    export POLICYENGINE_BENCH_RULEINDEX=$n    
    echo "Benchmarking rule at index $POLICYENGINE_BENCH_RULEINDEX"        
    ( $tm $DURATION ../../driver/sfprocessor -perflog -log=quiet -config=$CONFIG -driver=file $TRACES > $OUTDIR/rate_rule_$POLICYENGINE_BENCH_RULEINDEX.out ) &
    sleep 90
    n=$(($n + 1))
done