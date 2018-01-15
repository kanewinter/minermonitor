#!/bin/bash

C=0
while (( $(echo "$C < 2" |bc -l) )); do
watch -bctg -n 10 git pull
./nma.sh GitWatch WatchStopped "GitWatch Stopped" 2
C=`ls -1 wagerr | wc -l`
done
./nma.sh GitWatch FilesDetected "I see $C files" 2
