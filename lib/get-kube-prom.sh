#!/usr/bin/env bash

git clone --depth 1 git@github.com:prometheus-operator/kube-prometheus.git

cd kube-prometheus
jb install github.com/prometheus-operator/kube-prometheus/jsonnet/kube-prometheus@release-0.4
