# Setup the frontend

This section will cover how to build, push, and run the frontend application of tagenal. The frontend code can be found [here](./api/users/README.md).

## Build and push the frontend docker image

After modifying the frontend code, and creating a new version. We can build and push the new version onto a docker registry. We can modify the name of the docker image and the repository in the frontend's Makefile.

```
make build_push_frontend
```

## Run the frontend on kubernetes

To run the frontend on our kubernetes cluster, we use the following command:

```
make run_frontend_k8s
```

> To use another image than the default one, we need to change the kubernetes manifests in: `./kubernetes/frontend/frontend.yaml`, and specify the new image name/repository.

We can access the frontend at this URL: http://tagenal


## Stop the frontend

To stop the frontend, use the following command:

```
make stop_frontend_k8s
```