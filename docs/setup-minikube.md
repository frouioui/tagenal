# Quick Start - 1. Setup Kubernetes

In this first section of the quick start, we will setup our kubernetes cluster by using minikube.

> You need to have installed minikube, and respect the hardware requirements (with at least 10Gb of RAM).

## Start our kubernetes cluster

We are going to start minikube using the following command:

```
make start_minikube
```

> Minikube will use the driver `hyperkit`, which could be an issue for Linux and Windows distributions, as mentioned in [issue #11](https://github.com/frouioui/tagenal/issues/11).

## Start minikube dashboard (optional)

We can also start the provided minikube-dashboard. We can start it using the following command, ideally in a second terminal:

```
minikube dashboard
```

## Conclusion

After realizing this step, we have a ready-to-go kubernetes cluster. We can either check the status of our cluster using minikube dashboard, or by simply using the command line with `kubectl`.

## Next

In the next section we will setup Jaeger. Detailed [here](./setup-jaeger.md).