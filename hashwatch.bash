#!/bin/bash

MINER01=0
MINER02=1300

RESULT1=`curl -s localhost:7979/metrics | grep -v '^#' | grep miner01_hash | awk '{print $2}'`
RESULT2=`curl -s localhost:7979/metrics | grep -v '^#' | grep miner02_hash | awk '{print $2}'`

if (( $(echo "$MINER01 > $RESULT1" |bc -l) )); then
/opt/minermonitor/nma.sh HashWatch LowHash "Hashing is for miner01 is $RESULT1" 2
else
echo "$RESULT1"
fi

if (( $(echo "$MINER02 > $RESULT2" |bc -l) )); then
/opt/minermonitor/nma.sh HashWatch LowHash "Hashing is for miner02 is $RESULT2" 2
else
echo "$RESULT2"
fi
