# Quick Start - 5. Setup Redis

In this section we will create a Redis cluster, which we will use as a caching layer.

## Setup Redis cluster

The following command will setup a Redis cluster:

```
make setup_redis
```

Once the whole StatefulSet is up and running in kubernetes, we can proceed and execute the following command:

```
./kubernetes/redis/setup.sh
```

This command will tell the redis cluster that we want to split the cluster in three masters and three slaves.

## Next

Now, let's move on to the next step, seting up the monitoring in our cluster, which is explained [here](./setup-monitoring.md)