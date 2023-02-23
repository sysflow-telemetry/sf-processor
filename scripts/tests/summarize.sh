#!/bin/bash

tm=timeout
if ! command -v $tm &> /dev/null
then
    tm=gtimeout
fi

DIR=${1:-"."} 
VALUES=""
LABELS=""

for n in {1..10}
do    
    export POLICYENGINE_BENCH_RULESSAMPLESIZE=$(( $n*5 ))
    AVG=$(cat $DIR/rate_$POLICYENGINE_BENCH_RULESSAMPLESIZE.out | grep "Policy engine rate" | awk 'NR>2 {print $NF}' | awk '{ total += $1; count++ } END { print total/(1000*count) }')
    VALUES="$VALUES $AVG"
    LABELS="$LABELS $POLICYENGINE_BENCH_RULESSAMPLESIZE"
done

echo $VALUES
echo $LABELS
