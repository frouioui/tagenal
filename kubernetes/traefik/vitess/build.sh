#!/usr/bin/env bash

cd "${0%/*}"

yq w -i vitess_vtctld_ingress_route.yaml "spec.routes[0].services[0].name" $(kubectl get service --namespace=vitess --selector="planetscale.com/component=vtctld" -o custom-columns=":metadata.name" | head -n2)
yq w -i vitess_vtctld_client_ingress_route.yaml "spec.routes[0].services[0].name" $(kubectl get service --namespace=vitess --selector="planetscale.com/component=vtctld" -o custom-columns=":metadata.name" | head -n2)
yq w -i vitess_vtgate_ingress_route.yaml "spec.routes[0].services[0].name" $(kubectl get service --namespace=vitess -l planetscale.com/component=vtgate,planetscale.com/cell=zone1  -o custom-columns=":metadata.name" | head -n2)
yq w -i vitess_vtgate_grpc_ingress_route.yaml "spec.routes[0].services[0].name" $(kubectl get service --namespace=vitess -l planetscale.com/component=vtgate,planetscale.com/cell=zone1  -o custom-columns=":metadata.name" | head -n2)
yq w -i vitess_vttablet_ingress_route.yaml "spec.routes[0].services[0].name" $(kubectl get service --namespace=vitess --selector="planetscale.com/component=vttablet" -o custom-columns=":metadata.name" | head -n2)