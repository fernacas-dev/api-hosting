# !/bin/bash

echo "Remove node-exporter"
docker stop node-exporter && docker rm node-exporter
echo "Remove cadvisor"
docker stop cadvisor && docker rm cadvisor

echo "Start node-exporter"

docker run -p 9100:9100 --name=node-exporter --restart=always \
  -v "/:/host:ro,rslave" \
  -d prom/node-exporter

echo "Start cadvisor"

docker run -p 8084:8080 --name=cadvisor --restart=always \
  --device /dev/kmsg:/dev/kmsg \
  -v /:/rootfs:ro \
  -v /var/run:/var/run:ro \
  -v /sys:/sys:ro \
  -v /var/lib/docker/:/var/lib/docker:ro \
  -d google/cadvisor
