#!/bin/bash

MINER01=5500
MINER02=3100

RESULT1=`curl 'http://jaylin.ethosdistro.com/?json=yes' | jq '.rigs."501d77"."hash"'`
RESULT2=`curl 'http://jaylin.ethosdistro.com/?json=yes' | jq '.rigs."501b79"."hash"'`

if (( $(echo "$MINER01 > $RESULT1" |bc -l) )); then
/minermonitor/nma.sh HashWatch LowHash "Hashing is for miner01 is $RESULT1" 2
else
echo "$RESULT1"
fi

if (( $(echo "$MINER02 > $RESULT2" |bc -l) )); then
/minermonitor/nma.sh HashWatch LowHash "Hashing is for miner02 is $RESULT2" 2
else
echo "$RESULT2"
fi

echo "Sleeping..."
sleep 1m