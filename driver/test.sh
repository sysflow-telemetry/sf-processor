#!/bin/bash

for ((i=1;i<=100;i++)); 
do 
    echo "run $i"
    ./sfprocessor -log=error -config=./pipeline.graphlet.local.json ../resources/traces/../../../sf-graphlet/pynb/data/demo/1592328169.sampa.sf &> /tmp/out$i
    RES=$(tail -n 1 /tmp/out$i)
    if [ $RES = "false" ]; then
        cat /tmp/out$i
    fi
done

