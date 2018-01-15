#!/bin/bash

MINER01=195
MINER02=90

if [[ `curl 'http://jaylin.ethosdistro.com/?json=yes' | jq '.rigs | .. | objects | ."miner_hashes"' | grep "\ 00.00\ "` ]]; then
#00.00 not found
echo "Cards are up"
else
echo "00.00 found"
#/opt/minermonitor/nma.sh HashWatch DeadCard "00.00 was found in hash numbers" 2
fi

RESULT1=`curl 'http://jaylin.ethosdistro.com/?json=yes' | jq '.rigs."501d77"."hash"'`
RESULT2=`curl 'http://jaylin.ethosdistro.com/?json=yes' | jq '.rigs."501b79"."hash"'`

if (( $(echo "$MINER01 > $RESULT1" |bc -l) )); then
/opt/minermonitor/nma.sh HashWatch LowHash "Hashing is for miner01 is $RESULT1" 2
else
echo "miner01 is $RESULT1"
fi

if (( $(echo "$MINER02 > $RESULT2" |bc -l) )); then
/opt/minermonitor/nma.sh HashWatch LowHash "Hashing is for miner02 is $RESULT2" 2
else
echo "miner02 is $RESULT2"
fi
