#!/usr/bin/env bash

cd "${0%/*}"

if [ ! -d "../lib/kube-prometheus" ]
then
    chmod +x ../lib/script/get-kube-prom.sh
    ../lib/script/get-kube-prom.sh
fi

cd ../lib/kube-prometheus
./build.sh ../../monitoring/monitoring.jsonnet