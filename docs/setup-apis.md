# Setup the APIs

In this section we will elaborate on how to build, run, access tagenal's APIs services. There are currently two APIs:

- `users` [[see](./api/users/)]
- `articles` [[see](./api/articles/)]

## Build and Push docker images

After modifying the codebase. A new version of the docker image can be build and pushed to a public docker repository. We do so using the following command:

```
make build_push_apis
```

## Run the APIs on kubernetes

To run the APIs on our kubernetes cluster we use the following command. It will create the two APIs in the `default` namespace of our kubernetes cluster.

```
make run_apis_k8s
```

> To use another image than the default ones, we can change the kubernetes manifests in: `./kubernetes/api/**/*_api_server.yaml`, and specify the proper image names.

Now that our APIs are up and running, we can access them using the ingress routes that were automatically defined in the above command. The URLs are:

- http://api.tagenal/users
- http://api.tagenal/articles

## Stop the APIs

To stop the APIs we run the following command:

```
make stop_apis_k8s
```