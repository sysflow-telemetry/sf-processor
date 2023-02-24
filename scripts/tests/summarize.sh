#!/bin/bash

tm=timeout
if ! command -v $tm &> /dev/null
then
    tm=gtimeout
fi

DIR=${1:-"."} 
VALUES=""
LABELS=""

for f in `ls -v $DIR`
do   
    n=$(v=${f%.*} && printf "%s\n" "${v##*_}")
    AVG=$(cat $DIR/$f | grep "Policy engine rate" | awk 'NR>2 {print $NF}' | awk '{ total += $1; count++ } END { print total/(1000*count) }')
    VALUES="$VALUES $AVG"
    LABELS="$LABELS $(( $n ))"
done

echo $VALUES
echo $LABELS
