# Quick Start - 4. Setup Vitess

In this fourth section we are going to setup the Vitess cluster.

## Setup Vitess kubernetes operator

We are now going to do a few things. We will create a `vitess` namespace in our kubernetes cluster. We create the CRDs and other kubernetes objects required in order to run the operator, taken from the official [Vitess Operator GitHub repository](https://github.com/planetscale/vitess-operator). Finally, we will run the operator on kubernetes.

> You need to have cloned the vitess-operator repository in the `./lib` folder. You can simply run the command: `make clone_vitess_operator`.

The make command we will use is:

```
make setup_vitess_operator_kubernetes
```

After a moment, vitess-operator's pod should be up and running in kubernetes, which can be checked by running the following command.

```
kubectl get pods -n vitess
```

## Setup Vitess in kubernetes cluster

Now, we are going to setup a simple vitess cluster. We will create three keyspaces in our Vitess cluster. The three namespaces are:

- `config`
- `articles`
- `users`

Each keyspace will be formed of 2 kubernetes pods. We will have 1 pod for the `master`, and 1 other pod for the `replica` vttablet. We are also going to create a `ConfigMap` and a `Secret`. Both are hosting config that will be used throughout this quick start. The command we will run is:

```
make setup_vitess_kubernetes
```

After a moment, two new deployments should show up. We have the `vtgate` and `vtctld` deployments:

```
kubectl get deployment -n vitess
```
Should output:

```
NAME                           READY   UP-TO-DATE   AVAILABLE   AGE
vitess-operator                1/1     1            1           6m25s
vitess-zone1-vtctld-484cbf8c   1/1     1            1           2m39s
vitess-zone1-vtgate-1c6a88c4   1/1     1            1           2m39s
```

And we should also be able to see all our new `vttablet`:

```
kubectl get pods --selector="planetscale.com/component=vttablet" -n vitess
```

Should output:

```
NAME                                        READY   STATUS    RESTARTS   AGE
vitess-vttablet-zone1-0136547469-ddf1c39f   3/3     Running   1          5m
vitess-vttablet-zone1-0808196278-480a1110   3/3     Running   2          7m
vitess-vttablet-zone1-0837820985-44599818   3/3     Running   1          5m
vitess-vttablet-zone1-1651728703-a7361ca4   3/3     Running   2          5m
vitess-vttablet-zone1-3133170080-c4ca8d9c   3/3     Running   1          7m
vitess-vttablet-zone1-3249758605-17aa0869   3/3     Running   2          5m
vitess-vttablet-zone1-3724453505-5a510085   3/3     Running   2          5m
vitess-vttablet-zone1-3774659171-51411321   3/3     Running   1          7m
vitess-vttablet-zone1-3890619061-d6d202c2   3/3     Running   1          5m
vitess-vttablet-zone1-4224402714-96c8f9a7   3/3     Running   2          5m
```

Now, we will create the ingress routes to access our Vitess cluster. We will create two HTTP routes, one for vttablet's metrics, and another one for vtctld. We will also create 2 TCP routes, one to access VTGate's mysql interface, and the other to use vtctlclient.

```
make setup_traefik_vitess
```

We can access the two new endpoints to the following URLs:

- http://vtctld.tagenal/app/dashboard
- http://tagenal/vttablet/metrics

We can also connect to our Vitess cluster by using a MySQL client with the following command:

```
mysql -h tagenal -P 3000 -u user
```

We can use vtctlclient with this command:

```
vtctlclient -server=tagenal:8000
```

> We can also setup aliases:
> 
> ```
> alias mysql="mysql -h tagenal -P 3000 -u user"
> alias vtctlclient="vtctlclient -server=tagenal:8000"
> ```
> These aliases allow us to directly use `mysql` and `vtctlclient` commands.

## Setup the MySQL tables and Vindexes

We are now going to do a few things. We will create the SQL tables, create the incrementation sequences that will allow us to auto increment our tables, and finally we will upload the Vindexes to Vitess. The Vindexes are defining the behaviors our the tables in Vitess. We are going to execute the following command:

```
make init_vitess_all
```

## Setup the VReplication streams

Now that we have setup our tables and Vindexes, we are going to create some VReplication stream. VReplication streams are being used to replicate data across the shards of Vitess.

We will create a stream from the shard `-80` of the articles keyspace to shard `80-` of the same keyspace. This stream will replicate all the articles that have a `science` category into the shard `80-`. We will also create a second stream for the be_read table. This second stream will populate the `be_read` tables thanks to the `user_read` records.

We run the commands:

```
make init_vreplication_articles
make init_vreplication_article_be_read
```

## Setup Kubernetes CronJobs

We are now going to create some kubernetes CronJobs. We need to add three jobs. The jobs created by the CronJobs object will come populate the `popularity` table, given a certain cron schedule, the `popularity` table will be updated with new records given the `user_read` table. One CronJobs will handle the `daily` records, another one will compute the `weekly` records, and the last one will compute the `monthly` ones.

To setup the CronJobs we use the following command:

```
make init_cronjobs_popularity
```

## Insert some data

We are now going to insert some data into our newly created vitess cluster.

```
make insert_few_user_row
make insert_few_article_row
```

## Conclusion

In this section we have setup our Vitess cluster. We can verify the status of our Vitess cluster by running the command:

```
make show_vitess_tablets
```

This command will display all the vttablets running in our cluster. It should output something like:

```
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
| Cell  | Keyspace | Shard | TabletType | State   | Alias            | Hostname    | MasterTermStartTime  |
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
| zone1 | articles | -80   | MASTER     | SERVING | zone1-3249758605 | 172.17.0.7  | 2020-12-18T16:46:41Z |
| zone1 | articles | -80   | REPLICA    | SERVING | zone1-0837820985 | 172.17.0.18 |                      |
| zone1 | articles | 80-   | MASTER     | SERVING | zone1-3133170080 | 172.17.0.10 | 2020-12-18T16:46:33Z |
| zone1 | articles | 80-   | REPLICA    | SERVING | zone1-0136547469 | 172.17.0.16 |                      |
| zone1 | config   | -     | MASTER     | SERVING | zone1-1651728703 | 172.17.0.13 | 2020-12-18T16:46:46Z |
| zone1 | config   | -     | REPLICA    | SERVING | zone1-4224402714 | 172.17.0.11 |                      |
| zone1 | users    | -80   | MASTER     | SERVING | zone1-3774659171 | 172.17.0.9  | 2020-12-18T16:46:58Z |
| zone1 | users    | -80   | REPLICA    | SERVING | zone1-3724453505 | 172.17.0.5  |                      |
| zone1 | users    | 80-   | MASTER     | SERVING | zone1-3890619061 | 172.17.0.12 | 2020-12-18T16:46:43Z |
| zone1 | users    | 80-   | REPLICA    | SERVING | zone1-0808196278 | 172.17.0.8  |                      |
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
```

Thanks to the ingress routes we have created with Traefik Proxy earlier, we can access Vitess Vtctld web UI at http://vtctld.tagenal/app/dashboard.

## Next

Now, we will create a cache layer with Redis, which is explained [here](./setup-monitoring.md)