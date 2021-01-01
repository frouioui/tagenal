# Tagenal

[![Build Status](https://travis-ci.com/frouioui/tagenal.svg?token=XhmJBhJBxshbY6hsWepE&branch=master)](https://travis-ci.com/frouioui/tagenal) [![Go Report Card](https://goreportcard.com/badge/github.com/frouioui/tagenal)](https://goreportcard.com/report/github.com/frouioui/tagenal) [![Maintainability](https://api.codeclimate.com/v1/badges/55fb8c66d617c9425aff/maintainability)](https://codeclimate.com/github/frouioui/tagenal/maintainability)

## Description

Tagenal is a playground with a set of tools that enable experimentation in a cloud-native application.

Tagenal uses:

- Relational Database Sharding with Vitess
- Container Orchestration with Kubernetes
- Complete Runtime Tracing and Observation with Jaeger
- Application State Monitoring with Grafana, Alertmanager and Promotheus
- Caching with Redis Cluster
- APIs and Front-End Application

## Requirements
To run locally, tagenal needs:

- Have at least 10Gb of available RAM on the host
- Have Kubernetes / Minikube installed
- Have Golang version 1.15.x installed
- Install [yq](https://github.com/mikefarah/yq) a YAML processor
- Install [jsonnet-bundler](https://github.com/jsonnet-bundler/jsonnet-bundler) (jb), allowing us to deal with jsonnet files
- Install `jsonnet` using your system's packet manager
- Install `vtctlclient`, the following command `go get vitess.io/vitess/go/cmd/vtctlclient` can be used
- Install `mysql` and `mysql-client` using your system's packet manager
- Install `gojsontoyaml`, the following command `go get github.com/brancz/gojsontoyaml` can be used
- Run the shell scripts that are located in `./lib/*.sh`. These scripts will download the necessary basic libraries and repositories

## Quick Start

In this quick start we will cover the following items:

1. [Setup the Kubernetes cluster](./docs/setup-minikube.md)
1. [Setup Jaeger](./docs/setup-jaeger.md)
1. [Setup Traefik Proxy](./docs/setup-traefik.md)
1. [Setup the Vitess cluster](./docs/setup-vitess.md)
1. [Setup the Redis cluster](./docs/setup-redis.md)
1. [Setup monitoring with Grafana, Prometheus and Alertmanager](./docs/setup-monitoring.md)
1. [Setup the APIs and frontend](./docs/setup-api-frontend.md)

Once the quick start is over, we will have a fully setup application using distributed database systems.

## Generate sample data

The documentation and script to generate the sample data can be found [here](./scrips/gen/).

## Architecture

<img src="./docs/Tagenal k8s.png">
