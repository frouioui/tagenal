# Setup Jaeger

This section will cover how to setup the `observability` namespace of our kubernetes cluster.

## Setup Jaeger

We use the following command for a few things. It will create the whole `observability` namespace and the required CRDs. It will also add Jaeger configuration to the Vitess cluster and to Traefik.

```
make setup_jaeger
```

Tagenal is mostly experiment-oriented and not production-oriented. We then use the `AllInOne` configuration of Jaeger, which will greatly ease the deployment.

The command also created the ingress route to access Jaeger dashboard. Which can be access at this URL: http://jaeger.tagenal. On this interface we can see the trace of our services.

## Next step

The next step is to run the APIs in our kubernetes cluster. Which is detailed [in the next section](./setup-apis.md).