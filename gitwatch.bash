#!/bin/bash

wait watch -bcg -n 10 git pull

ls -l | grep -v ^l | wc -l

C=`ls -1 /opt/wagerr/ | wc -l`

if ( C > 3 ); then
./nma.sh GitWatch FilesDetected "I see $C files" 2
else
./nma.sh GitWatch WatchStopped "GitWatch Stopped" 2
fi
