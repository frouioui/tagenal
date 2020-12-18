# Quick Start - 2. Setup Jaeger

This section will cover how to setup our first kubernetes namespace: `observability`. This namespace will host our observability stack, for now only composed of Jaeger.

## Setup Jaeger

The following command will do a few things for us. It will create the `observability` namespace, and the CRDs that Jaeger needs to successfully run on kubernetes. Then, it will 

```
make setup_jaeger
```

Tagenal is mostly experiment-oriented and not production-oriented. We then use the `AllInOne` configuration of Jaeger, which will greatly ease the deployment.

By using minikube dashboard or `kubectl` we can observe the successful deployment of Jaeger.

## Next

For now, we can't access Jaeger UI, we will cover that in the next section thanks to Traefik Proxy. Next step is [here](./setup-traefik.md).

<!-- The next step is to run the APIs in our kubernetes cluster. Which is detailed [in the next section](./setup-apis.md). -->