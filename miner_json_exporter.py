from prometheus_client import start_http_server, Metric, REGISTRY
from prometheus_client.core import GaugeMetricFamily, REGISTRY
import json
import requests
import sys
import time
import urllib2
import logging, sys
logging.basicConfig(stream=sys.stderr, level=logging.DEBUG)

class JsonCollector(object):
  def __init__(self, endpoint):
    self._endpoint = endpoint
  def collect(self):
    metric = GaugeMetricFamily(
        'miner_hashes',
        'list of hashes',
        labels=["jobname"])

    result = json.load(urllib2.urlopen(
        "http://jaylin.ethosdistro.com/?json=yes="
        + "rigs[name,miner_hashes[string]]"))

    for rig in result['rigs']:
      name = rig['name']
      # If there's a null result, we want to export a zero.
      status = rig['miner_hashes'] or {}
      metric.add_metric([name], status)

    yield metric

    # Metrics with labels for the documents loaded
    metric = Metric('per_hash-gpu', 'avg hashing')

    result = json.load(urllib2.urlopen(
        "http://jaylin.ethosdistro.com/?json=yes="
        + "rigs[name,per_hash-gpu[long]]"))

    for rig in result['rigs']:
      name = rig['name']
      # If there's a null result, we want to export a zero.
      status = rig['per_hash-gpu'] or {}
      metric.add_metric([name], status)

    yield metric




if __name__ == '__main__':
  # Usage: json_exporter.py port endpoint
logging.info('first switch', sys.argv[1])
  start_http_server(int(sys.argv[1]))
logging.info('first switch', sys.argv[2])
  REGISTRY.register(JsonCollector(sys.argv[2]))


  while True: time.sleep(1)
