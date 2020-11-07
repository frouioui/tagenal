#!/usr/bin/env bash

cd ..
git clone --depth 1 git@github.com/prometheus-operator/kube-prometheus.github

cd kube-prometheus
jb install github.com/prometheus-operator/kube-prometheus/jsonnet/kube-prometheus@release-0.4
