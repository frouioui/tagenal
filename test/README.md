# Test for MongoDB cluster with Docker-Compose

## The cluster

The cluster is, as shown on the official documentation, composed of 3 main parts.
- Router
- Config servers
- Data nodes

The router part contains only one node in this example. The config servers is a replica set of 3 nodes.

The data nodes are composed of:

- 2 shards
- 3 nodes in each shard

## How to run

Must have installed docker-compose before hand. To run:

`docker-compose up`

It will show a bunch of logs. If it is the first time starting the cluster, then it must be initialized with the Makefile.

The Makefile has several commands, here are the most important ones to do after starting the cluster with docker-compose:

- `make init_all` this command will init all the shards, router, and config servers. This command run multiple shell commands, but the last shell command, which will setup the router, is likely to fail: in case the shards and config servers did not have enough time to initialize. To check if this command worked, use the following one.

- `make router_status`, this command should output the state of the whole cluster. If the command fails then the router hasn't been initialized correctly. To initiliaze it correctly, use the following command.

- `make init_router`, this will simply initialize the router only.

- `make router_setup_shard_tags`, this command will setup the zone tags for the shards, it will define the geographical areas of the different shards. This command will also create our test database called `tagenal`, and a collection `users`. This collection will be set up to automatically split the users depending on their region: users with region "Beijing" will be sotred in the shard #1, and users from "Hong Kong" in the shard #2.

After running these commands, our test cluster should be set up.

## How to test

Installing `MongoDB Compass`, will allow us to connect to the cluster thanks to a GUI. Once it is installed, the connection string for the cluster is: `mongodb://localhost:21117`.

From there, we should be able to see the `tagenal` database on the sidebar, and its `users` collection. On the `users` collection we can now import a JSON file of users.

This JSON file can be generated with the JSON data generator script, located in: `./scripts/gen/genTable_mongoDB10GB_json.py` (from the root of this repository).

After importing the users, we can check on the `Explain Plan` tab of the collection, click `EXPLAIN`, and see that data is coming from the 2 shards.

