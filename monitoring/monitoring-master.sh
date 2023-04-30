# !/bin/bash

#docker stop postgres-exporter
#docker rm stop postgres-exporter

#docker stop keycloak-postgres-exporter
#docker rm stop keycloak-postgres-exporter

#docker run --name=postgres-exporter \
# --restart=always -d \
#  -p 9187:9187 \
#  -e DATA_SOURCE_NAME="postgresql://root:Xetid2019*@192.168.31.24:26257/postgres?sslmode=disable" \
#  registry.gitlab-vcl.xutil.net:5000/devops/postgres-exporter

#docker run --name=keycloak-postgres-exporter \
# --restart=always -d \
#  -p 9188:9187 \
#  -e DATA_SOURCE_NAME="postgresql://root:Xetid2019*@10.12.170.220:5432/postgres?sslmode=disable" \
#  registry.gitlab-vcl.xutil.net:5000/devops/postgres-exporter

echo "Remove prometheus"
docker stop prometheus && docker rm prometheus
#echo "Remove node-exporter"
#docker stop node-exporter && docker rm node-exporter
#echo "Remove cadvisor"
#docker stop cadvisor && docker rm cadvisor
echo "Remove grafana"
docker stop grafana && docker rm grafana

MY_PATH=/home/cloud/MONITORINGDATA
ACTUAL=/home/cloud/golang_projects/api-hosting/monitoring

echo "Start prometheus"

docker run -p 9090:9090 --name=prometheus --restart=always \
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
#docker run -p 9100:9100 --name=node-exporter --restart=always \
# -v "/:/host:ro,rslave" \
# -d prom/node-exporter

#echo "Start cadvisor"

#docker run -p 8085:8080 --name=cadvisor --restart=always \
# --device /dev/kmsg:/dev/kmsg \
# -v /:/rootfs:ro \
# -v /var/run:/var/run:ro \
# -v /sys:/sys:ro \
# -v /var/lib/docker/:/var/lib/docker:ro \
# -d registry.gitlab-vcl.xutil.net:5000/devops/cadvisor

echo "Start grafana"

docker run -p 3000:3000 -e GF_SECURITY_ADMIN_PASSWORD=DontTouchMyGrafana2021* --name=grafana --restart=always \
  -v $MY_PATH/grafana_db:/var/lib/grafana \
  -d grafana/grafana
