#!/usr/bin/env bash

# TODO: redirect the output of Vtctld into a log files rather than stdout
kubectl port-forward "$(kubectl get service --selector="planetscale.com/component=vtctld" -o name | head -n1)" 15000 15999 &

kubectl port-forward "$(kubectl get service --selector="planetscale.com/component=vtgate" -o name | head -n1)" 15306:3306 &

kubectl port-forward "$(kubectl get service --selector="planetscale.com/component=vttablet" -o name | head -n1)" 9104:9104 &