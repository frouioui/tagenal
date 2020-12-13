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

----
## Quick Start

In this quick start we will cover the following items:

- Setup our Kubernetes cluster
- Create a sharded Vitess cluster
- Add Traefik Proxy to our cluster
- Enable tracing with Jaeger
- Setup some APIs
- Setup a Frontend application
- Setup some monitoring with Grafana, Prometheus and Alertmanager

Once the quick start is over, we will have a fully setup application using distributed database systems.

### Requirements
Before we start, there are some requirements:

- Have at least 10Gb of available RAM on the host
- Install [yq](https://github.com/mikefarah/yq) a YAML processor
- Install [jsonnet-bundler](https://github.com/jsonnet-bundler/jsonnet-bundler) (jb), allowing us to deal with jsonnet files
- Install jsonnet using you system's packet manager
- Install vtctlclient, the following command `go get vitess.io/vitess/go/cmd/vtctlclient` can be used
- Install `mysql` and `mysql-client` using your system's packet manager
- Have Kubernetes / Minikube installed
- Have Golang version 1.15.x installed
- Run the shell scripts that are located in `./lib/*.sh`. These scripts will download the basic libraries and repositories that we need

### Steps
The quick start is architectured as followed:
1. [Setup of a simple cluster](./docs/setup-minikube-vitess.md)
2. [Setup our sharded Vitess cluster](./docs/setup-sharded-vitess-cluster.md)
3. [Setup monitoring](./docs/setup-monitoring.md)
4. [Setup Jaeger](./docs/setup-jaeger.md)
5. [Start the APIs](./docs/setup-apis.md)
6. [Start the frontend](./docs/setup-frontend.md)

---

## Data models

Tagenal is composed of 2 main tables for now. They are:

### user table
```sql
CREATE TABLE user (
  _id INT NOT NULL auto_increment,
  timestamp CHAR(14) DEFAULT NULL,
  id CHAR(5) DEFAULT NULL,
  uid CHAR(5) DEFAULT NULL,
  name CHAR(9) DEFAULT NULL,
  gender CHAR(7) DEFAULT NULL,
  email CHAR(10) DEFAULT NULL,
  phone CHAR(10) DEFAULT NULL,
  dept CHAR(9) DEFAULT NULL,
  grade CHAR(7) DEFAULT NULL,
  language CHAR(3) DEFAULT NULL,
  region VARBINARY(256),
  role CHAR(6) DEFAULT NULL,
  preferTags CHAR(7) DEFAULT NULL,
  obtainedCredits CHAR(3) DEFAULT NULL,
  PRIMARY KEY(_id)
);
```

### article table
```sql
CREATE TABLE article (
  _id INT NOT NULL auto_increment,
  timestamp CHAR(14) DEFAULT NULL,
  id CHAR(7) DEFAULT NULL,
  aid CHAR(7) DEFAULT NULL,
  title CHAR(15) DEFAULT NULL,
  category VARBINARY(256) DEFAULT NULL,
  abstract CHAR(30) DEFAULT NULL,
  articleTags CHAR(14) DEFAULT NULL,
  authors CHAR(40) DEFAULT NULL,
  language CHAR(3) DEFAULT NULL,
  text TEXT(500) DEFAULT NULL,
  image CHAR(255) DEFAULT NULL,
  video CHAR(255) DEFAULT NULL,
  PRIMARY KEY(_id)
);
```
