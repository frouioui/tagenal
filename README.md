# Tagenal

[![Build Status](https://travis-ci.com/frouioui/tagenal.svg?token=XhmJBhJBxshbY6hsWepE&branch=master)](https://travis-ci.com/frouioui/tagenal)

## Description of the project
<img align="right" width="100" height="100" src="./assets/img/Tsinghua_University_Logo.png">

This project is part of the **Distributed Database Systems** class of the Advanced Computer Science master degree at **Tsinghua University**.

<br>

## List of main features

- Bulk load of the User, Article, and Read tables
- Query users, articles, and users' readings
- Insert new data in the Be-Read table
- Query top 5 daily/weekly/monthly articles, with their details
- Efficient execution of the data insert, update, and queries
- Monitoring of the whole distributed system

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

