#!/bin/bash

C=0
while (( $(echo "$C < 3" |bc -l) )); do
watch -bctg -n 10 git pull
./nma.sh GitWatch WatchStopped "GitWatch Stopped" 2 &
C=`ls -1 | wc -l`
echo $C
done
./nma.sh GitWatch FilesDetected "I see $C files" 2 &

while (( $(echo "$C > 3" |bc -l) )); do
./nma.sh GitWatch FilesDetected "I see $C files" 2 &
sleep 5m
done
