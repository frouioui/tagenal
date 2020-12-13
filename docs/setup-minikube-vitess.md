# Setup Minikube / Vitess / Traefik

In this first section of the quick start, we will:
1. Start our kubernetes cluster
2. Run the vitess-operator in our kubernetes cluster
3. Setup a simple version of the Vitess cluster
4. Setup Traefik

## Start our kubernetes cluster

We are going to start minikube using the following command:

```
make start_minikube
```

> Minikube will use the driver `hyperkit`, which could be an issue for Linux and Windows distributions, as mentioned in [issue #11](https://github.com/frouioui/tagenal/issues/11).

Additionally, minikube-dashboard can be run using the following command, ideally in a second terminal:

```
make start_minikube_dashboard
```

## Setup Vitess kubernetes operator

We are now going to do a few things. We will create a `vitess` namespace in our kubernetes cluster. We create the custom resources (CRD) and other kubernetes objects required to run the operator, taken from the official [Vitess Operator GitHub repository](https://github.com/planetscale/vitess-operator). Finally, we will run the operator on kubernetes.

> You need to have cloned the vitess-operator repository in the `./lib` folder. You can simply run the command: `make clone_vitess_operator`.

The make-command we will use is:

```
make install_vitess_operator
```

After a moment, vitess-operator's pod should be up and running in kubernetes, which can be checked by running the following command.

```
kubectl get pods -n vitess
```

## Setup a simple vitess cluster

Now, we are going to setup a simple vitess cluster. We will create 3 keyspaces in vitess:

- config
- articles
- users

Each keyspace will be formed of 2 kubernetes pods. We will have 1 pod for the `master`, and 1 other pod for the `replica`. We are also going to create a `ConfigMap` and a `Secret`. Both are hosting config that will be used throughout this quick-start. The command we will now run is:

```
make init_kubernetes_unsharded_database
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
vitess-vttablet-zone1-0512591799-f4f6e2c2   3/3     Running   2          3m58s
vitess-vttablet-zone1-1651728703-a7361ca4   3/3     Running   2          3m58s
vitess-vttablet-zone1-1929229597-af5a9c25   3/3     Running   3          3m58s
vitess-vttablet-zone1-3555790305-a9553d09   3/3     Running   2          3m58s
vitess-vttablet-zone1-4109238961-30b1e68c   3/3     Running   3          3m58s
vitess-vttablet-zone1-4224402714-96c8f9a7   3/3     Running   3          3m58s
```

## Setup Traefik proxy

We will use Traefik Proxy as our kubernetes ingress controller. The following command will create all the necessary Kubernetes CRD, and setup traefik in our kubernetes cluster.

```
make setup_traefik
```

After a moment, the command:

```
kubectl get pods
```

Should output:

```
NAME      READY   UP-TO-DATE   AVAILABLE   AGE
traefik   1/1     1            1           5m29s
```

Now, we will create the ingress routes to access our Vitess cluster. We will create two HTTP routes, one for vttablet's metrics, and another one for vtctld. We will also create 2 TCP routes, one to access VTGate's mysql interface, and the other to use vtctlclient.

```
make setup_traefik_vitess
```

We now need to get minikube's ip address. We will modify the `hosts file` of our host. Allowing us to forward minikube's ip address to some custom domain name.

```
minikube ip
```

The expected output should be similar to:

```
192.168.64.10
```

We will now open `/etc/hosts` and add the following lines at the very end of the file:

```
192.168.64.10 tagenal
192.168.64.10 grafana.tagenal
192.168.64.10 prometheus.tagenal
192.168.64.10 alertmanager.tagenal
192.168.64.10 api.tagenal
192.168.64.10 jaeger.tagenal
```

> Replace the IP address by the one that `minikube ip` showed.

>For now, we can ignore the meaning of the last 5 lines. They will come in handy in the next sections of this quick start.

To conclude this step, we can now access the following URLs:

- http://tagenal:8080
- http://tagenal/vtctld/app/dashboard
- http://tagenal/vttablet/metrics

We can also connect to vitess cluster using a MySQL client with the following command:

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

## Conclusion

In this section we have:
- Started our kubernetes cluster
- Created a basic Vitess cluster
- Setup Traefik

We can verify the status of our Vitess cluster by running the command:

```
make show_vitess_tablets
```

This command will display all the vttablets running in our cluster. It should output something like:

```
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
| Cell  | Keyspace | Shard | TabletType | State   | Alias            | Hostname    | MasterTermStartTime  |
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
| zone1 | articles | -     | MASTER     | SERVING | zone1-0512591799 | 172.17.0.15 | 2020-12-04T12:38:37Z |
| zone1 | articles | -     | REPLICA    | SERVING | zone1-3555790305 | 172.17.0.14 |                      |
| zone1 | config   | -     | MASTER     | SERVING | zone1-1651728703 | 172.17.0.13 | 2020-12-04T12:38:54Z |
| zone1 | config   | -     | REPLICA    | SERVING | zone1-4224402714 | 172.17.0.12 |                      |
| zone1 | users    | -     | MASTER     | SERVING | zone1-1929229597 | 172.17.0.11 | 2020-12-04T12:39:34Z |
| zone1 | users    | -     | REPLICA    | SERVING | zone1-4109238961 | 172.17.0.16 |                      |
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
```

## Next step

The next step is to shard our newly created Vitess cluster. Which is detailed [in the next section](./setup-sharded-vitess-cluster.md).