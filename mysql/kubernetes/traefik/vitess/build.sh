#!/usr/bin/env bash

yq w -i vitess_vtctld_ingress_route.yaml "spec.routes[0].services[0].name" $(kubectl get service --selector="planetscale.com/component=vtctld" -o custom-columns=":metadata.name" | head -n2)