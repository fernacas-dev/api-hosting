# !/bin/bash

echo "Remove prometheus"
docker stop prometheus && docker rm prometheus
echo "Remove grafana"
docker stop grafana && docker rm grafana

echo "Start prometheus"

MY_PATH=/home/fernando/MONITORINGDATA
ACTUAL=$PWD

docker run \
 -p 9090:9090 \
 --name=prometheus \
 --restart=always \
 --network=go-network \
 -v $ACTUAL/prometheus/config.yml:/etc/prometheus/prometheus.yml \
 -v $MY_PATH/prometheus_db/var/lib:/var/lib/prometheus \
 -v $MY_PATH/prometheus_db/prometheus:/prometheus \
 -v $MY_PATH/prometheus_db/alert.rules:/etc/prometheus/alert.rules \
 -d prom/prometheus \
 --config.file=/etc/prometheus/prometheus.yml \
 --web.route-prefix=/ \
 --storage.tsdb.retention.time=200h \
 --web.enable-lifecycle

#echo "Start node-exporter"
#docker run \
# -p 9100:9100 \
# --network=go-network \
# --name=node-exporter \
# --restart=always \
# -v "/:/host:ro,rslave" \
# -d registry.gitlab-vcl.xutil.net:5000/devops/node-exporter

#echo "Start cadvisor"
#docker run \
# -p 8085:8080 \
# --name=cadvisor \
# --network=go-network \
# --restart=always \
# --device /dev/kmsg:/dev/kmsg \
# -v /:/rootfs:ro \
# -v /var/run:/var/run:ro \
# -v /sys:/sys:ro \
# -v /var/lib/docker/:/var/lib/docker:ro \
# -d registry.gitlab-vcl.xutil.net:5000/devops/cadvisor

echo "Start grafana"
docker run \
 -p 3000:3000 \
 -e GF_SECURITY_ADMIN_PASSWORD=Xetid2020* \
 --name=grafana  \
 --network=go-network \
 --restart=always \
 -v $MY_PATH/grafana_db:/var/lib/grafana \
 -v /home/fernando/metrics/logs:/var/log/nginx \
 -d grafana/grafana

