#!/bin/bash

curl -s http://jaylin.ethosdistro.com/?json=yes -o data.json

curl -s -u user:XC9ZEpdrjqQW -XPOST "http://35.202.37.2/elasticsearch/watcherdata/log" -d @data.json > /dev/null

