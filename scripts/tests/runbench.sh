#!/bin/bash

./benchmark-rules.sh 90 ./pipeline.falco.bench.json 52 ../../../datasets/k8s/wcm_drill_3_5 falco_rule_index 
./benchmark-rules.sh 90 ./pipeline.sigma.bench.json 136 ../../../datasets/k8s/wcm_drill_3_5 sigma_rule_index
./benchmark.sh 120 ./pipeline.falco.bench.json ../../../datasets/k8s/wcm_drill_3_5 falco_ruleset 
./benchmark.sh 120 ./pipeline.sigma.bench.json ../../../datasets/k8s/wcm_drill_3_5 sigma_ruleset