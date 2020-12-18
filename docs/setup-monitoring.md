# Quick Start - 5. Setup monitoring

The following section will allow us to setup the monitoring tools for our kubernetes cluster. We will be using Grafana, Prometheus, and Alertmanager. The goal of this section is to setup the `monitoring` namespace.

## Setup kube-prometheus

We are going to use the library [kube-prometheus](https://github.com/prometheus-operator/kube-prometheus). This library gives us a good interface to the prometheus Kubernetes operator, allowing us to quickly add and modify Grafana dashboards using jsonnet.

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

## Next

We are now reaching the end of this quick start, to finish, we will complete our application by setting up the APIs and a frontend application, this step is detailed [here](./setup-api-frontend.md).