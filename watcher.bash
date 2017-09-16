#!/bin/bash

INDEX="{"index": {"_index": "meminfo", "_type": "log", "_id":0}}"

curl http://jaylin.ethosdistro.com/?json=yes -o data.json

echo $INDEX > watcher.json
cat data.json >> watcher.json

curl -u user:XC9ZEpdrjqQW -XPOST http://35.202.37.2/elasticsearch/watcherdata/log/_bulk --data-binary @watcher.json
