#!/bin/sh

kubectl port-forward "$(kubectl get service --selector="planetscale.com/component=vtctld" -o name | head -n1)" 15000 15999 &
kubectl port-forward "$(kubectl get service --selector="planetscale.com/component=vtgate" -o name | head -n1)" 15306:3306 &