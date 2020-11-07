#!/usr/bin/env bash

cd "${0%/*}"

if [ ! -d "../lib/grafana-dashboards" ]
then
    chmod +x ../lib/script/get-grafana-dashboards.sh
    ../lib/script/get-grafana-dashboards.sh
fi

sed 's/Prometheus/prometheus/g' ../lib/grafana-dashboards/dashboards/MySQL_Overview.json > ../lib/grafana-dashboards/dashboards/MySQL_Overview_mod.json

if [ ! -d "../lib/kube-prometheus" ]
then
    chmod +x ../lib/script/get-kube-prom.sh
    ../lib/script/get-kube-prom.sh
fi

cd ../lib/kube-prometheus
./build.sh ../../monitoring/monitoring.jsonnet