#!/bin/bash

curl -s http://jaylin.ethosdistro.com/?json=yes -o data.json

curl -s -u user:XC9ZEpdrjqQW -XPOST "http://35.202.37.2/elasticsearch/watcherdata-`date +%j`/log" -d @data.json > /dev/null


curl -s http://jaylin.ethosdistro.com/?json=yes -o data.json | \
jq --arg timestamp `date -u --iso-8601=seconds` '.timestamp = $timestamp' > data.json
curl -s -u user:XC9ZEpdrjqQW -XPOST "http://35.202.37.2/elasticsearch/watcherdata/log" -d @data.json > /dev/null
