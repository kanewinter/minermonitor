#!/bin/bash

MINER01=8000
MINER02=3200

if [[ `curl 'http://jaylin.ethosdistro.com/?json=yes' | jq '.rigs | .. | objects | ."miner_hashes"' | grep "\ 00.00\ "` ]]; then
#00.00 not found
echo "Cards are up"
else
/opt/minermonitor/nma.sh HashWatch DeadCard "00.00 was found in hash numbers" 2
fi

RESULT1=`curl 'http://jaylin.ethosdistro.com/?json=yes' | jq '.rigs."4fe11f"."hash"'`
RESULT2=`curl 'http://jaylin.ethosdistro.com/?json=yes' | jq '.rigs."501b79"."hash"'`

if (( $(echo "$MINER01 > $RESULT1" |bc -l) )); then
./nma.sh HashWatch LowHash "Hashing is for miner01 is $RESULT1" 2
else
echo "$RESULT1"
fi

if (( $(echo "$MINER02 > $RESULT2" |bc -l) )); then
./nma.sh HashWatch LowHash "Hashing is for miner02 is $RESULT2" 2
else
echo "$RESULT2"
fi
