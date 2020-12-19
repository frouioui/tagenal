# Quick Start - 3. Setup Traefik

In this third section we will cover how to add an ingress controller to our kubernetes cluster. We will use Traefik Proxy.

## Setup Traefik Proxy

The following command will create all the necessary CRDs that Traefik needs, and setup Traefik in our kubernetes cluster. It will also create an ingress route pointing to Jaeger web UI, which we previously created.

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

We now need to get minikube's ip address. We will modify the `hosts file` of our host. Allowing us to forward minikube's ip address to some custom domain name, in our case the domain `tagenal` and its subdomains.

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
192.168.64.10 vtctld.tagenal
```

> Replace the IP address by the one that `minikube ip` showed.

> We already add all the domains so we won't have to modify the file again. Most of them will come in handy later during the quick start.

To conclude this step, we can now access the following URLs:

- http://tagenal:8080
- http://jaeger.tagenal

The former is Traefik web UI, and the latter is Jaeger web UI.

## Conclusion

We have added Traefik to our kubernetes cluster, we added a route to Jaeger's web UI.

## Next

Now it is time to setup the Vitess cluster, which is described [here](./setup-vitess.md).