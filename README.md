# Tagenal

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

## Generation of the database

All the necessary scripts for the databases' data generation can be found in `./scripts/gen/`.

## Data models

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
- Install vtctlclient, the following command `go get vitess.io/vitess/go/cmd/vtctlclient` can be used
- Install `mysql` and `mysql-client` using your system's packet manager
- Have Kubernetes / Minikube installed
- Have Golang version 1.15.x installed
- Run the shell scripts that are located in `./lib/*.sh`. These scripts will download the basic libraries and repositories that we need

### Sections
The quick start is architectured as followed:
- [Setup of a simple cluster](./docs/setup-minikube-vitess.md)
- Setup our sharded Vitess cluster


### 5. Create MySQL tables and Vitess VSchemas

We will now move on to the creation of the first tables, and the VSchemas.

We are going to create the tables in the two keyspaces: users, and articles. And at the same time we will give Vitess an empty VSchema for each table.

`make init_unsharded_database`

### 6. Initialize the incrementation sequences in the config keyspace

We will now create incrementation sequences. They will be used to increment the `PRIMARY KEY` of the sharded tables, in keyspace users and articles. The incrementation sequences are stored in the config keyspace.

In addition to creating the sequences, we will modify the VSchema of the users and articles keyspaces' tables, to take into account the new incrementation sequences.

`make init_config_increment_sequence`


### 7. Initialize sharded database in kubernetes

We will now create new pods to prepare for the sharding. This new kubernetes configuration of the vitess cluster will create new vttablets.

In the two keyspaces, users and articles, we are going to have the following configuration:
- shard `-`
- shard `-80`
- shard `80-`

All shards will keep the same replication topology for now.

`make init_sharded_database`

This command will take a little while to fully finish. The status can be checked by running:

`kubectl get pods`

Additionally, the `make show_vitess_tablets` command should show only tablets with a running state `SERVING`.

### 8. Insert a first few rows

Once the previous step is fully completed, simply add a few rows:

`make insert_few_article_row`

`make insert_few_user_row`

### 9. Setup the region-based sharding

We are now going to setup the region-based sharding on the keyspaces users and articles.

`make init_region_sharding_users`

`make init_region_sharding_articles`

### 10. Start the resharding process of the cluster

Now we are complete the resharding process of our keyspaces.

`make resharding_process_users`

`make resharding_process_articles`

As an example, the expected output for the users keyspace is:

```
vtctlclient -server=localhost:15999 Reshard users.user2user '-' '-80,80-'
vtctlclient -server=localhost:15999 VDiff users.user2user
Summary for user: {ProcessedRows:5 MatchingRows:5 MismatchedRows:0 ExtraRowsSource:0 ExtraRowsTarget:0}
E1101 20:23:50.599068   52103 main.go:64] E1101 19:23:50.595298 vdiff.go:780] Draining extra row(s) found on the source starting with: [INT32(1) VARBINARY("\x01\x16k@\xb4J\xbaK\xd6")]
Summary for user_lookup: {ProcessedRows:5 MatchingRows:0 MismatchedRows:0 ExtraRowsSource:5 ExtraRowsTarget:0}
vtctlclient -server=localhost:15999 SwitchReads -tablet_type=replica users.user2user
vtctlclient -server=localhost:15999 SwitchReads -tablet_type=rdonly users.user2user
vtctlclient -server=localhost:15999 SwitchWrites users.user2user
```

For the articles keyspace the output should be similar.

However, an issue have been found where the `VDiff` command fails when trying to replicate empty rows. This issue will cause the master vttablet of the shard `-` to fail. In this case, simply delete the pods, wait for another one to start up. Once the whole keyspace looks good, all vttablets are in `SERVING` states, you shall continue the procedure, using the following commands:

```
vtctlclient -server=tagenal:8000 SwitchReads -tablet_type=replica users.user2user
vtctlclient -server=tagenal:8000 SwitchReads -tablet_type=rdonly users.user2user
vtctlclient -server=tagenal:8000 SwitchWrites users.user2user
```

*Note: replace `user` and `users` by whichever table / keyspace is failing.*

### 11. Shutdown the old shards and add replicas

We are going to shutdown the shard `-` from keyspace users and articles.

`make final_vitess_cluster`

This command will also increase the count of replica to 2 instead of 1.

### 12. Start the VReplication of the article table from shard -80 to shard 80-

We are now going to create a VRep stream in order to duplicate the articles with the science category, from shard `-80` to shard `80-`. To do so, use the following command:

`make init_vreplication_articles`

The expected output is:

```
vtctlclient -server=localhost:15999 VReplicationExec zone1-136547469 'insert into _vt.vreplication (db_name, source, pos, max_tps, max_replication_lag, tablet_types, time_updated, transaction_timestamp, state) values('"'"'articles'"'"', '"'"'keyspace:\"articles\" shard:\"-80\" filter:<rules:<match:\"article\" filter:\"select * from article where category=\'"'"'science\'"'"'\" > > '"'"', '"'"''"'"', 9999, 9999, '"'"'master'"'"', 0, 0, '"'"'Running'"'"')' 
+
+
```

### 13. Test

We are now going to display our two tables `user` and `article`.

Let's start first with the `user` table. We will display all the rows, then only the rows in shard `-80` (users located in `Beijing`), and finally the rows in shard `80-` (users located in `Hong Kong`):

`make show_user_table`

Expected output:

```
Users
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+-----------+-------+------------+-----------------+
| _id | timestamp     | id   | uid  | name  | gender | email  | phone  | dept   | grade  | language | region    | role  | preferTags | obtainedCredits |
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+-----------+-------+------------+-----------------+
|   1 | 1506328859000 | u0   | 0    | user0 | male   | email0 | phone0 | dept13 | grade1 | zh       | Beijing   | role2 | tags24     | 42              |
|   2 | 1506328859001 | u1   | 1    | user1 | female | email1 | phone1 | dept5  | grade1 | en       | Beijing   | role2 | tags7      | 22              |
|   3 | 1506328859002 | u2   | 2    | user2 | male   | email2 | phone2 | dept4  | grade4 | en       | Beijing   | role2 | tags46     | 62              |
|   4 | 1506328859003 | u3   | 3    | user3 | female | email3 | phone3 | dept15 | grade4 | zh       | Beijing   | role1 | tags0      | 2               |
|   5 | 1506328859004 | u4   | 4    | user4 | male   | email4 | phone4 | dept15 | grade4 | en       | Hong Kong | role2 | tags18     | 63              |
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+-----------+-------+------------+-----------------+
Users -80
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+---------+-------+------------+-----------------+
| _id | timestamp     | id   | uid  | name  | gender | email  | phone  | dept   | grade  | language | region  | role  | preferTags | obtainedCredits |
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+---------+-------+------------+-----------------+
|   1 | 1506328859000 | u0   | 0    | user0 | male   | email0 | phone0 | dept13 | grade1 | zh       | Beijing | role2 | tags24     | 42              |
|   2 | 1506328859001 | u1   | 1    | user1 | female | email1 | phone1 | dept5  | grade1 | en       | Beijing | role2 | tags7      | 22              |
|   3 | 1506328859002 | u2   | 2    | user2 | male   | email2 | phone2 | dept4  | grade4 | en       | Beijing | role2 | tags46     | 62              |
|   4 | 1506328859003 | u3   | 3    | user3 | female | email3 | phone3 | dept15 | grade4 | zh       | Beijing | role1 | tags0      | 2               |
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+---------+-------+------------+-----------------+
Users 80-
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+-----------+-------+------------+-----------------+
| _id | timestamp     | id   | uid  | name  | gender | email  | phone  | dept   | grade  | language | region    | role  | preferTags | obtainedCredits |
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+-----------+-------+------------+-----------------+
|   5 | 1506328859004 | u4   | 4    | user4 | male   | email4 | phone4 | dept15 | grade4 | en       | Hong Kong | role2 | tags18     | 63              |
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+-----------+-------+------------+-----------------+
```

With the same motivation as the previous command, we will run the command displaying the same type of information, for the `article` table. This time, the rows are fragmented using the `category` column, the `science` category will be stored in shard `-80`. The `science` and `technology` categories will be kept in shard `80-`.

`make show_article_table`

And the expected output is:

```
Articles
+-----+---------------+------+------+--------+------------+-----------------------+-------------+------------+----------+-------------+-----------------------------------------------------------------------------+-------+
| _id | timestamp     | id   | aid  | title  | category   | abstract              | articleTags | authors    | language | text        | image                                                                       | video |
+-----+---------------+------+------+--------+------------+-----------------------+-------------+------------+----------+-------------+-----------------------------------------------------------------------------+-------+
|   3 | 1506000000002 | a2   | 2    | title2 | science    | abstract of article 2 | tags40      | author384  | zh       | text_a2.txt | image_a2_0.jpg,image_a2_1.jpg,                                              |       |
|   5 | 1506000000004 | a4   | 4    | title4 | science    | abstract of article 4 | tags5       | author613  | zh       | text_a4.txt | image_a4_0.jpg,                                                             |       |
|   1 | 1506000000000 | a0   | 0    | title0 | technology | abstract of article 0 | tags46      | author1059 | en       | text_a0.txt | image_a0_0.jpg,image_a0_1.jpg,image_a0_2.jpg,                               |       |
|   2 | 1506000000001 | a1   | 1    | title1 | technology | abstract of article 1 | tags48      | author1950 | en       | text_a1.txt | image_a1_0.jpg,image_a1_1.jpg,image_a1_2.jpg,                               |       |
|   3 | 1506000000002 | a2   | 2    | title2 | science    | abstract of article 2 | tags40      | author384  | zh       | text_a2.txt | image_a2_0.jpg,image_a2_1.jpg,                                              |       |
|   4 | 1506000000003 | a3   | 3    | title3 | technology | abstract of article 3 | tags9       | author1741 | en       | text_a3.txt | image_a3_0.jpg,                                                             |       |
|   5 | 1506000000004 | a4   | 4    | title4 | science    | abstract of article 4 | tags5       | author613  | zh       | text_a4.txt | image_a4_0.jpg,                                                             |       |
|   6 | 1506000000005 | a5   | 5    | title5 | technology | abstract of article 5 | tags43      | author1131 | en       | text_a5.txt | image_a5_0.jpg,image_a5_1.jpg,image_a5_2.jpg,image_a5_3.jpg,                |       |
|   7 | 1506000000006 | a6   | 6    | title6 | technology | abstract of article 6 | tags19      | author916  | en       | text_a6.txt | image_a6_0.jpg,image_a6_1.jpg,image_a6_2.jpg,image_a6_3.jpg,image_a6_4.jpg, |       |
+-----+---------------+------+------+--------+------------+-----------------------+-------------+------------+----------+-------------+-----------------------------------------------------------------------------+-------+
Articles -80
+-----+---------------+------+------+--------+----------+-----------------------+-------------+-----------+----------+-------------+--------------------------------+-------+
| _id | timestamp     | id   | aid  | title  | category | abstract              | articleTags | authors   | language | text        | image                          | video |
+-----+---------------+------+------+--------+----------+-----------------------+-------------+-----------+----------+-------------+--------------------------------+-------+
|   3 | 1506000000002 | a2   | 2    | title2 | science  | abstract of article 2 | tags40      | author384 | zh       | text_a2.txt | image_a2_0.jpg,image_a2_1.jpg, |       |
|   5 | 1506000000004 | a4   | 4    | title4 | science  | abstract of article 4 | tags5       | author613 | zh       | text_a4.txt | image_a4_0.jpg,                |       |
+-----+---------------+------+------+--------+----------+-----------------------+-------------+-----------+----------+-------------+--------------------------------+-------+
Articles 80-
+-----+---------------+------+------+--------+------------+-----------------------+-------------+------------+----------+-------------+-----------------------------------------------------------------------------+-------+
| _id | timestamp     | id   | aid  | title  | category   | abstract              | articleTags | authors    | language | text        | image                                                                       | video |
+-----+---------------+------+------+--------+------------+-----------------------+-------------+------------+----------+-------------+-----------------------------------------------------------------------------+-------+
|   1 | 1506000000000 | a0   | 0    | title0 | technology | abstract of article 0 | tags46      | author1059 | en       | text_a0.txt | image_a0_0.jpg,image_a0_1.jpg,image_a0_2.jpg,                               |       |
|   2 | 1506000000001 | a1   | 1    | title1 | technology | abstract of article 1 | tags48      | author1950 | en       | text_a1.txt | image_a1_0.jpg,image_a1_1.jpg,image_a1_2.jpg,                               |       |
|   3 | 1506000000002 | a2   | 2    | title2 | science    | abstract of article 2 | tags40      | author384  | zh       | text_a2.txt | image_a2_0.jpg,image_a2_1.jpg,                                              |       |
|   4 | 1506000000003 | a3   | 3    | title3 | technology | abstract of article 3 | tags9       | author1741 | en       | text_a3.txt | image_a3_0.jpg,                                                             |       |
|   5 | 1506000000004 | a4   | 4    | title4 | science    | abstract of article 4 | tags5       | author613  | zh       | text_a4.txt | image_a4_0.jpg,                                                             |       |
|   6 | 1506000000005 | a5   | 5    | title5 | technology | abstract of article 5 | tags43      | author1131 | en       | text_a5.txt | image_a5_0.jpg,image_a5_1.jpg,image_a5_2.jpg,image_a5_3.jpg,                |       |
|   7 | 1506000000006 | a6   | 6    | title6 | technology | abstract of article 6 | tags19      | author916  | en       | text_a6.txt | image_a6_0.jpg,image_a6_1.jpg,image_a6_2.jpg,image_a6_3.jpg,image_a6_4.jpg, |       |
+-----+---------------+------+------+--------+------------+-----------------------+-------------+------------+----------+-------------+-----------------------------------------------------------------------------+-------+
```

## Setup monitoring

The following set of commands will allow us to setup a fully monitored cluster using Grafana, Prometheus, and Alertmanager.

### 1. Run kube-prometheus

We are going to use the library kube-prometheus (https://github.com/prometheus-operator/kube-prometheus). This libary gives us a good interface to the prometheus Kubernetes operator, and allow us to quickly add and modify Grafana dashboards using jsonnet.

To build the operator's manifest, and create all the Kubernetes CRD in our cluster use the following command:

`make run_monitoring`

It will also instantiate the whole `monitoring` keyspace inside Kubernetes, with it, the pods and deployment of Prometheus, Grafana, Alertmanager.

### 2. Setup ingress routes

To access our newly created services we are going to create 3 ingress routes with Traefik.

`make setup_traefik_monitoring`

Now we can access our services:

- http://grafana.tagenal
- http://prometheus.tagenal
- http://alertmanager.tagenal

> Note: this step requires you to run [Setup Traefik proxy](###4.-setup-traefik-proxy)

We can login on http://grafana.tagenal with the following credentials:
- User: admin
- Password: admin

## Setup the APIs services

In this section we will elaborate on how to build, run, access tagenal's APIs services.

There are currently two APIs:

- `users` [[doc](./api/users/README.md)][[folder](./api/users/)]
- `articles` [[doc](./api/articles/README.md)][[folder](./api/articles/)]

### Build and push a new docker image

After modifying the codebase, new version of the docker image can be build and pushed to a public docker repository.

`make build_push_apis`

This command will require to change the docker-image name in each of the service's Makefile.

### Run the APIs on kubernetes

To run the APIs on kubernetes:

`make run_apis_k8s`

If you want to use another image than the default ones, this can be done by changing the kubernetes manifests in: `./kubernetes/api/**/*_api_server.yaml`.

### Stop the APIs

`make stop_apis_k8s`

## Setup the frontend

This section will cover how to build, push, and run the frontend application of tagenal.

The frontend's code can also be found [here](./api/users/README.md).

### Build and push the frontend image

Run the following command to build and push the frontend docker image into a docker registery:

`make build_push_frontend`

This command will require to change the docker image name in the frontend Makefile.

### Run the frontend on kubernetes

To run the frontend on kubernetes:

`make run_frontend_k8s`

### Stop the frontend

`make stop_frontend_k8s`