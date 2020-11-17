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

## How to run

Tagenal uses Kubernetes and Vitess as main components, it is therefore necessary to install a few things before getting started.

### Prerequisite

- At least 12Gb of available RAM on the host
- Install yq a YAML processor https://github.com/mikefarah/yq
- Install jsonnet-bundler (jb) allowing us to deal with jsonnet files https://github.com/jsonnet-bundler/jsonnet-bundler
- Install vtctlclient `go get vitess.io/vitess/go/cmd/vtctlclient`
- Install `mysql` and `mysql-client` using your system's packet manager
- Have kubernetes
- Have golang version 1.15.x
- Run the scripts in `./lib/*.sh` to download a few repositories

### 1. Start minikube

To start minikube run the following command:

`make start_minikube`

Minikube will use the driver `hyperkit`, which could be an issue for Linux and Windows distribution as mentioned in issue #11 (https://github.com/frouioui/tagenal/issues/11).

Additionally Kubernetes dashboard can be run using the following command, ideally in a second terminal:

`make start_minikube_dashboard`

### 2. Setup Vitess kubernetes operator

This command will create vitess-operator Kubernetes CRD in Minikube and start the vitess-operator:

`make install_vitess_operator`

The last line of the output should be:

```
deployment.apps/vitess-operator created
```

After a moment, vitess-operator's pod should be up and running in kubernetes.

### 3. Initialize unsharded database

We will now do a simple first initialization of the vitess clusterr. We will load the needed configuration into kubernetes and create three keyspaces:

- config
- articles
- users

For now we will have a simple 1 Master / 1 Replica configuration for each keyspace.

The command is:

`make init_kubernetes_unsharded_database`

And, the expected output is:

```
kubectl apply -f kubernetes/vitess_cluster_secret.yaml
secret/vitess-cluster-secret created
kubectl apply -f kubernetes/vitess_cluster_config.yaml
configmap/vitess-cluster-config-sharding created
kubectl apply -f kubernetes/init_cluster_vitess.yaml
vitesscluster.planetscale.com/vitess created
```

### 4. Setup Traefik proxy

We use Traefik as a kubernetes ingress controller. Traefik proxy will enable us to commnunicate directly with our services.

The following command will create all the necessary Kubernetes CRD, and create a traefik deployment in addition to the Admin Web UI of traefik:

`make setup_traefik`

After a while, the tail of the output should look like:

```
deployment.apps/traefik created
service/traefik created
```

Now, we will create the ingress routes to access vitess' services. We will create two HTTP routes, one for vttablet, and another one for vtctld. We will also create two TCP routes, one to access VTGate's mysql interface, and the other to use vtctlclient.

`make setup_traefik_vitess`

Once the latter command is completed, we are going to fetch minikube's ip:

`minikube ip`

And expect an output similar to:

`192.168.64.10`

We will now open `/etc/hosts` and add the following lines at the very end of the file:

```
192.168.64.10 tagenal
192.168.64.10 grafana.tagenal
192.168.64.10 prometheus.tagenal
192.168.64.10 alertmanager.tagenal
```

*You replace the IP address by the one `minikube ip` showed you.*

For now, we can ignore the last 3 lines we added in the file, they will come in handy later on.

To conclude this step, we can now access the following URLs:

- http://tagenal:8080
- http://tagenal/vtctld/app/dashboard
- http://tagenal/vttablet/metrics

We can also connect to vitess cluster using a MySQL client with the following credentials:

`mysql -h tagenal -P 3000 -u user`

And finally, we can use vtctlclient with the following configuration:

`vtctlclient -server=tagenal:8000`

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