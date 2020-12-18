# Quick Start - 7. Setup the APIs and frontend

In this final section, we are going to setup two APIs and a frontend application. We will deploy them to the `default` namespace of our kubernetes cluster.

We will cover how to build, push, and run the frontend application and APIs. The frontend code can be found [here](./api/users/README.md), and the APIs can be found here:

- `users` [[see](./api/users/)]
- `articles` [[see](./api/articles/)]

## Frontend

### Build and push

After modifying the frontend code, and creating a new version. We can build and push the new version onto a docker registry. We can modify the name of the docker image and the repository in the frontend's Makefile.

```
make build_push_frontend
```

### Run the frontend on kubernetes

To run the frontend on our kubernetes cluster, we use the following command:

```
make run_frontend_k8s
```

> To use another image than the default one, we need to change the kubernetes manifests in: `./kubernetes/frontend/frontend.yaml`, and specify the new image name/repository.

We can access the frontend at this URL: http://tagenal


### Stop the frontend

To stop the frontend, use the following command:

```
make stop_frontend_k8s
```


## APIs

### Build and Push docker images

After modifying the codebase. A new version of the docker image can be built and pushed to a public docker repository. We do so using the following command:

```
make build_push_apis
```

### Run the APIs on kubernetes

To run the APIs on our kubernetes cluster we use the following command. This will create the two APIs in the `default` namespace of our kubernetes cluster.

```
make run_apis_k8s
```

> To use another image than the default ones, we can change the kubernetes manifests in: `./kubernetes/api/**/*_api_server.yaml`, and specify the proper image names.

Now that our APIs are up and running, we can access them using the ingress routes that were automatically defined in the above command. The URLs are:

- http://api.tagenal/users
- http://api.tagenal/articles

### Stop the APIs

To stop the APIs we run the following command:

```
make stop_apis_k8s
```