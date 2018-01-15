#!/bin/bash

<<<<<<< HEAD
MINER01=195
MINER02=90

if [[ `curl 'http://jaylin.ethosdistro.com/?json=yes' | jq '.rigs | .. | objects | ."miner_hashes"' | grep "\ 00.00\ "` ]]; then
#00.00 not found
echo "Cards are up"
else
echo "00.00 found"
#/opt/minermonitor/nma.sh HashWatch DeadCard "00.00 was found in hash numbers" 2
fi
=======
MINER01=5500
MINER02=3100

while true
do
>>>>>>> 7aa086c1d9ecb269d05ffde21e921f1e89f4e0ec

RESULT1=`curl 'http://jaylin.ethosdistro.com/?json=yes' | jq '.rigs."501d77"."hash"'`
RESULT2=`curl 'http://jaylin.ethosdistro.com/?json=yes' | jq '.rigs."501b79"."hash"'`

if (( $(echo "$MINER01 > $RESULT1" |bc -l) )); then
<<<<<<< HEAD
/opt/minermonitor/nma.sh HashWatch LowHash "Hashing is for miner01 is $RESULT1" 2
=======
/minermonitor/nma.sh HashWatch LowHash "Hashing is for miner01 is $RESULT1" 2
>>>>>>> 7aa086c1d9ecb269d05ffde21e921f1e89f4e0ec
else
echo "miner01 is $RESULT1"
fi

if (( $(echo "$MINER02 > $RESULT2" |bc -l) )); then
<<<<<<< HEAD
/opt/minermonitor/nma.sh HashWatch LowHash "Hashing is for miner02 is $RESULT2" 2
=======
/minermonitor/nma.sh HashWatch LowHash "Hashing is for miner02 is $RESULT2" 2
>>>>>>> 7aa086c1d9ecb269d05ffde21e921f1e89f4e0ec
else
echo "miner02 is $RESULT2"
fi

echo "`date` Sleeping..."
sleep 1m

done
