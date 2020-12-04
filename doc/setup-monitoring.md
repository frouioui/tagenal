# Setup monitoring

The following section will allow us to setup a fully monitored cluster using Grafana, Prometheus, and Alertmanager.

## Setup kube-prometheus

We are going to use the library [kube-prometheus](https://github.com/prometheus-operator/kube-prometheus). This libary gives us a good interface to the prometheus Kubernetes operator, and allow us to quickly add and modify Grafana dashboards using jsonnet.

To build the operator's manifest, and create all the Kubernetes CRD in our cluster use the following command:

```
make run_monitoring
```

It will also instantiate the whole `monitoring` keyspace inside Kubernetes, with it, the pods and deployment of Prometheus, Grafana, Alertmanager.

## Setup ingress routes

To access our newly created services, we create 3 ingress routes with the Traefik Proxy we defined earlier.

```
make setup_traefik_monitoring
```

Now we can access our services on these URLs:

- http://grafana.tagenal
- http://prometheus.tagenal
- http://alertmanager.tagenal

We can login on http://grafana.tagenal with the following credentials:
- User: *admin*
- Password: *admin*

## Next step

The next step is to setup Jaeger. Which is detailled [in the next section](./setup-jaeger.md).