# Setup a sharded vitess cluster

This section will cover how to shard our Vitess cluster.

## Create MySQL tables and setup Vitess configuration

We now move onto the creation of the first MySQL tables and VSchemas. We are going to create the tables in the two keyspaces: users, and articles. And at the same time we will give Vitess an empty VSchema for each table.

```
make init_unsharded_database
```

We are now going to other tables in the `config` keyspace we defined earlier. These tables will be used as incrementation sequences for Vitess. Since we are going to shard our two other keyspaces, they cannot handle `auto_increment` MySQL column. Therefore we need other tables in a non-sharded keyspace to handle the `PRIMARY KEY` columns auto incrementation. We are also going to modify the VSchema of the `users` and `articles` keyspaces. So they take into account the newly created incrementation sequences.

```
make init_config_increment_sequence
```

## Create new pods for sharding

We will now create new pods in our kubernetes cluster to prepare for the sharding of our vitess keyspaces. In the keyspaces `users` and `articles`, we are going to have 3 shards for now:

- A shard `-`
- A shard `-80`
- A shard `80-`

These 3 shards, will each contain 2 vttablets, a master, and a replica.

```
make init_sharded_database
```

This command might take a little while to fully finish. We can check our new pods with the following command:

```
kubectl get pods --selector="planetscale.com/component=vttablet" -n vitess
```

It should output:

```
NAME                                        READY   STATUS    RESTARTS   AGE
vitess-vttablet-zone1-0136547469-ddf1c39f   3/3     Running   1          2m18s
vitess-vttablet-zone1-0512591799-f4f6e2c2   3/3     Running   0          107s
vitess-vttablet-zone1-0808196278-480a1110   3/3     Running   1          2m18s
vitess-vttablet-zone1-0837820985-44599818   3/3     Running   1          2m18s
vitess-vttablet-zone1-1651728703-a7361ca4   3/3     Running   0          108s
vitess-vttablet-zone1-1929229597-af5a9c25   3/3     Running   0          110s
vitess-vttablet-zone1-3133170080-c4ca8d9c   3/3     Running   1          2m18s
vitess-vttablet-zone1-3249758605-17aa0869   3/3     Running   1          2m19s
vitess-vttablet-zone1-3555790305-a9553d09   3/3     Running   0          33s
vitess-vttablet-zone1-3724453505-5a510085   3/3     Running   1          2m19s
vitess-vttablet-zone1-3774659171-51411321   3/3     Running   1          2m18s
vitess-vttablet-zone1-3890619061-d6d202c2   3/3     Running   1          2m18s
vitess-vttablet-zone1-4109238961-30b1e68c   3/3     Running   0          30s
vitess-vttablet-zone1-4224402714-96c8f9a7   3/3     Running   0          35s
```

Additionally, we can use the `make show_vitess_tablets` to see our vttablets. The command should show only tablets with a running state `SERVING`.

```
make show_vitess_tablets
```

Should now ouputs:

```
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
| Cell  | Keyspace | Shard | TabletType | State   | Alias            | Hostname    | MasterTermStartTime  |
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
| zone1 | articles | -     | MASTER     | SERVING | zone1-0512591799 | 172.17.0.11 | 2020-12-04T13:37:27Z |
| zone1 | articles | -     | REPLICA    | SERVING | zone1-3555790305 | 172.17.0.12 |                      |
| zone1 | articles | -80   | MASTER     | SERVING | zone1-3249758605 | 172.17.0.18 | 2020-12-04T13:36:38Z |
| zone1 | articles | -80   | REPLICA    | SERVING | zone1-0837820985 | 172.17.0.22 |                      |
| zone1 | articles | 80-   | MASTER     | SERVING | zone1-0136547469 | 172.17.0.26 | 2020-12-04T13:36:39Z |
| zone1 | articles | 80-   | REPLICA    | SERVING | zone1-3133170080 | 172.17.0.23 |                      |
| zone1 | config   | -     | MASTER     | SERVING | zone1-1651728703 | 172.17.0.13 | 2020-12-04T13:37:30Z |
| zone1 | config   | -     | REPLICA    | SERVING | zone1-4224402714 | 172.17.0.6  |                      |
| zone1 | users    | -     | MASTER     | SERVING | zone1-1929229597 | 172.17.0.10 | 2020-12-04T13:37:34Z |
| zone1 | users    | -     | REPLICA    | SERVING | zone1-4109238961 | 172.17.0.14 |                      |
| zone1 | users    | -80   | MASTER     | SERVING | zone1-3774659171 | 172.17.0.20 | 2020-12-04T13:36:27Z |
| zone1 | users    | -80   | REPLICA    | SERVING | zone1-3724453505 | 172.17.0.21 |                      |
| zone1 | users    | 80-   | MASTER     | SERVING | zone1-0808196278 | 172.17.0.24 | 2020-12-04T13:36:40Z |
| zone1 | users    | 80-   | REPLICA    | SERVING | zone1-3890619061 | 172.17.0.25 |                      |
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
```

## Setup region-based sharding

We are now going to setup region-based sharding on the keyspaces `users` and `articles`. We use the following commands:

```
make init_region_sharding_users
make init_region_sharding_articles
```

We can now complete the resharding process of our keyspaces. For that, we will use the following commands:

```
make resharding_process_users
make resharding_process_articles
```

The results of the above commands can be verified by running

```
make show_vitess_tablets
```

The expected output is:

```
+-------+----------+-------+------------+-------------+------------------+-------------+----------------------+
| Cell  | Keyspace | Shard | TabletType | State       | Alias            | Hostname    | MasterTermStartTime  |
+-------+----------+-------+------------+-------------+------------------+-------------+----------------------+
| zone1 | articles | -     | MASTER     | NOT_SERVING | zone1-0512591799 | 172.17.0.11 | 2020-12-04T13:37:27Z |
| zone1 | articles | -     | REPLICA    | SERVING     | zone1-3555790305 | 172.17.0.12 |                      |
| zone1 | articles | -80   | MASTER     | SERVING     | zone1-3249758605 | 172.17.0.18 | 2020-12-04T13:36:38Z |
| zone1 | articles | -80   | REPLICA    | SERVING     | zone1-0837820985 | 172.17.0.22 |                      |
| zone1 | articles | 80-   | MASTER     | SERVING     | zone1-0136547469 | 172.17.0.26 | 2020-12-04T13:36:39Z |
| zone1 | articles | 80-   | REPLICA    | SERVING     | zone1-3133170080 | 172.17.0.23 |                      |
| zone1 | config   | -     | MASTER     | SERVING     | zone1-1651728703 | 172.17.0.13 | 2020-12-04T13:37:30Z |
| zone1 | config   | -     | REPLICA    | SERVING     | zone1-4224402714 | 172.17.0.6  |                      |
| zone1 | users    | -     | MASTER     | NOT_SERVING | zone1-1929229597 | 172.17.0.10 | 2020-12-04T13:37:34Z |
| zone1 | users    | -     | REPLICA    | SERVING     | zone1-4109238961 | 172.17.0.14 |                      |
| zone1 | users    | -80   | MASTER     | SERVING     | zone1-3774659171 | 172.17.0.20 | 2020-12-04T13:36:27Z |
| zone1 | users    | -80   | REPLICA    | SERVING     | zone1-3724453505 | 172.17.0.21 |                      |
| zone1 | users    | 80-   | MASTER     | SERVING     | zone1-0808196278 | 172.17.0.24 | 2020-12-04T13:36:40Z |
| zone1 | users    | 80-   | REPLICA    | SERVING     | zone1-3890619061 | 172.17.0.25 |                      |
+-------+----------+-------+------------+-------------+------------------+-------------+----------------------+
```

Here the masters vttablets of shard `-` of the `users` and `articles` keyspaces are in a `NOT_SERVING` state which is totally expected. We are not using the shard `-` anymore. We are now going to remove these unused shard. And we will also give the final setup to our vitess cluster. We use the following command:

```
make final_vitess_cluster
```

This operation will take a while to create the new pods. We can verify the status of our operation with the commands:

```
kubectl get pods --selector="planetscale.com/component=vttablet" -n vitess

make show_vitess_tablets
```

The expected outputs are:

```
NAME                                        READY   STATUS    RESTARTS   AGE
vitess-vttablet-zone1-0136547469-ddf1c39f   3/3     Running   0          2m2s
vitess-vttablet-zone1-0808196278-480a1110   3/3     Running   0          110s
vitess-vttablet-zone1-0837820985-44599818   3/3     Running   0          2m2s
vitess-vttablet-zone1-1651728703-a7361ca4   3/3     Running   0          112s
vitess-vttablet-zone1-3133170080-c4ca8d9c   3/3     Running   0          52s
vitess-vttablet-zone1-3249758605-17aa0869   3/3     Running   0          52s
vitess-vttablet-zone1-3724453505-5a510085   3/3     Running   0          2m2s
vitess-vttablet-zone1-3774659171-51411321   3/3     Running   0          52s
vitess-vttablet-zone1-3890619061-d6d202c2   3/3     Running   0          41s
vitess-vttablet-zone1-4224402714-96c8f9a7   3/3     Running   0          50s
```
And 
```
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
| Cell  | Keyspace | Shard | TabletType | State   | Alias            | Hostname    | MasterTermStartTime  |
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
| zone1 | articles | -80   | MASTER     | SERVING | zone1-0837820985 | 172.17.0.11 | 2020-12-04T14:10:59Z |
| zone1 | articles | -80   | REPLICA    | SERVING | zone1-3249758605 | 172.17.0.17 |                      |
| zone1 | articles | 80-   | MASTER     | SERVING | zone1-0136547469 | 172.17.0.12 | 2020-12-04T14:10:54Z |
| zone1 | articles | 80-   | REPLICA    | SERVING | zone1-3133170080 | 172.17.0.18 |                      |
| zone1 | config   | -     | MASTER     | SERVING | zone1-1651728703 | 172.17.0.13 | 2020-12-04T14:11:03Z |
| zone1 | config   | -     | REPLICA    | SERVING | zone1-4224402714 | 172.17.0.19 |                      |
| zone1 | users    | -80   | MASTER     | SERVING | zone1-3724453505 | 172.17.0.10 | 2020-12-04T14:10:52Z |
| zone1 | users    | -80   | REPLICA    | SERVING | zone1-3774659171 | 172.17.0.6  |                      |
| zone1 | users    | 80-   | MASTER     | SERVING | zone1-0808196278 | 172.17.0.14 | 2020-12-04T14:11:09Z |
| zone1 | users    | 80-   | REPLICA    | SERVING | zone1-3890619061 | 172.17.0.20 |                      |
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
```

Once this is done, we can move one and create a VReplication stream between. This stream duplicate the articles with the science category, from shard `-80` to shard `80-`. To do so, use the following command:

```
make init_vreplication_articles
```

> This replication is an initial requirement from the project.

## Conclusion

The sharding is finished. We can now test our sharded cluster. We will go ahead and insert a few rows in multiple tables using the commands:

```
make insert_few_user_row
make insert_few_article_row
```

Then, we can get a glance at our tables. Let's start first with the `user` table. We will display all the rows, then only the rows in shard `-80` (users located in `Beijing`), and finally the rows in shard `80-` (users located in `Hong Kong`):

```
make show_user_table
```

Expected output:

```
Users
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+-----------+-------+------------+-----------------+
| _id | timestamp     | id   | uid  | name  | gender | email  | phone  | dept   | grade  | language | region    | role  | preferTags | obtainedCredits |
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+-----------+-------+------------+-----------------+
|   1 | 1506328859000 | u0   | 0    | user0 | male   | email0 | phone0 | dept13 | grade1 | zh       | Beijing   | role2 | tags24     | 42              |
|   2 | 1506328859001 | u1   | 1    | user1 | female | email1 | phone1 | dept5  | grade1 | en       | Beijing   | role2 | tags7      | 22              |
|   3 | 1506328859002 | u2   | 2    | user2 | male   | email2 | phone2 | dept4  | grade4 | en       | Beijing   | role2 | tags46     | 62              |
|   4 | 1506328859003 | u3   | 3    | user3 | female | email3 | phone3 | dept15 | grade4 | zh       | Beijing   | role1 | tags0      | 2               |
|   5 | 1506328859004 | u4   | 4    | user4 | male   | email4 | phone4 | dept15 | grade4 | en       | Hong Kong | role2 | tags18     | 63              |
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+-----------+-------+------------+-----------------+
Users -80
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+---------+-------+------------+-----------------+
| _id | timestamp     | id   | uid  | name  | gender | email  | phone  | dept   | grade  | language | region  | role  | preferTags | obtainedCredits |
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+---------+-------+------------+-----------------+
|   1 | 1506328859000 | u0   | 0    | user0 | male   | email0 | phone0 | dept13 | grade1 | zh       | Beijing | role2 | tags24     | 42              |
|   2 | 1506328859001 | u1   | 1    | user1 | female | email1 | phone1 | dept5  | grade1 | en       | Beijing | role2 | tags7      | 22              |
|   3 | 1506328859002 | u2   | 2    | user2 | male   | email2 | phone2 | dept4  | grade4 | en       | Beijing | role2 | tags46     | 62              |
|   4 | 1506328859003 | u3   | 3    | user3 | female | email3 | phone3 | dept15 | grade4 | zh       | Beijing | role1 | tags0      | 2               |
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+---------+-------+------------+-----------------+
Users 80-
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+-----------+-------+------------+-----------------+
| _id | timestamp     | id   | uid  | name  | gender | email  | phone  | dept   | grade  | language | region    | role  | preferTags | obtainedCredits |
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+-----------+-------+------------+-----------------+
|   5 | 1506328859004 | u4   | 4    | user4 | male   | email4 | phone4 | dept15 | grade4 | en       | Hong Kong | role2 | tags18     | 63              |
+-----+---------------+------+------+-------+--------+--------+--------+--------+--------+----------+-----------+-------+------------+-----------------+
```

With the same motivation as the previous command, we will run the command displaying the same type of information, for the `article` table. This time, the rows are fragmented using the `category` column, the `science` category will be stored in shard `-80`. The `science` and `technology` categories will be kept in shard `80-`.

`make show_article_table`

And the expected output is:

```
Articles
+-----+---------------+------+------+--------+------------+-----------------------+-------------+------------+----------+-------------+-----------------------------------------------------------------------------+-------+
| _id | timestamp     | id   | aid  | title  | category   | abstract              | articleTags | authors    | language | text        | image                                                                       | video |
+-----+---------------+------+------+--------+------------+-----------------------+-------------+------------+----------+-------------+-----------------------------------------------------------------------------+-------+
|   3 | 1506000000002 | a2   | 2    | title2 | science    | abstract of article 2 | tags40      | author384  | zh       | text_a2.txt | image_a2_0.jpg,image_a2_1.jpg,                                              |       |
|   5 | 1506000000004 | a4   | 4    | title4 | science    | abstract of article 4 | tags5       | author613  | zh       | text_a4.txt | image_a4_0.jpg,                                                             |       |
|   1 | 1506000000000 | a0   | 0    | title0 | technology | abstract of article 0 | tags46      | author1059 | en       | text_a0.txt | image_a0_0.jpg,image_a0_1.jpg,image_a0_2.jpg,                               |       |
|   2 | 1506000000001 | a1   | 1    | title1 | technology | abstract of article 1 | tags48      | author1950 | en       | text_a1.txt | image_a1_0.jpg,image_a1_1.jpg,image_a1_2.jpg,                               |       |
|   3 | 1506000000002 | a2   | 2    | title2 | science    | abstract of article 2 | tags40      | author384  | zh       | text_a2.txt | image_a2_0.jpg,image_a2_1.jpg,                                              |       |
|   4 | 1506000000003 | a3   | 3    | title3 | technology | abstract of article 3 | tags9       | author1741 | en       | text_a3.txt | image_a3_0.jpg,                                                             |       |
|   5 | 1506000000004 | a4   | 4    | title4 | science    | abstract of article 4 | tags5       | author613  | zh       | text_a4.txt | image_a4_0.jpg,                                                             |       |
|   6 | 1506000000005 | a5   | 5    | title5 | technology | abstract of article 5 | tags43      | author1131 | en       | text_a5.txt | image_a5_0.jpg,image_a5_1.jpg,image_a5_2.jpg,image_a5_3.jpg,                |       |
|   7 | 1506000000006 | a6   | 6    | title6 | technology | abstract of article 6 | tags19      | author916  | en       | text_a6.txt | image_a6_0.jpg,image_a6_1.jpg,image_a6_2.jpg,image_a6_3.jpg,image_a6_4.jpg, |       |
+-----+---------------+------+------+--------+------------+-----------------------+-------------+------------+----------+-------------+-----------------------------------------------------------------------------+-------+
Articles -80
+-----+---------------+------+------+--------+----------+-----------------------+-------------+-----------+----------+-------------+--------------------------------+-------+
| _id | timestamp     | id   | aid  | title  | category | abstract              | articleTags | authors   | language | text        | image                          | video |
+-----+---------------+------+------+--------+----------+-----------------------+-------------+-----------+----------+-------------+--------------------------------+-------+
|   3 | 1506000000002 | a2   | 2    | title2 | science  | abstract of article 2 | tags40      | author384 | zh       | text_a2.txt | image_a2_0.jpg,image_a2_1.jpg, |       |
|   5 | 1506000000004 | a4   | 4    | title4 | science  | abstract of article 4 | tags5       | author613 | zh       | text_a4.txt | image_a4_0.jpg,                |       |
+-----+---------------+------+------+--------+----------+-----------------------+-------------+-----------+----------+-------------+--------------------------------+-------+
Articles 80-
+-----+---------------+------+------+--------+------------+-----------------------+-------------+------------+----------+-------------+-----------------------------------------------------------------------------+-------+
| _id | timestamp     | id   | aid  | title  | category   | abstract              | articleTags | authors    | language | text        | image                                                                       | video |
+-----+---------------+------+------+--------+------------+-----------------------+-------------+------------+----------+-------------+-----------------------------------------------------------------------------+-------+
|   1 | 1506000000000 | a0   | 0    | title0 | technology | abstract of article 0 | tags46      | author1059 | en       | text_a0.txt | image_a0_0.jpg,image_a0_1.jpg,image_a0_2.jpg,                               |       |
|   2 | 1506000000001 | a1   | 1    | title1 | technology | abstract of article 1 | tags48      | author1950 | en       | text_a1.txt | image_a1_0.jpg,image_a1_1.jpg,image_a1_2.jpg,                               |       |
|   3 | 1506000000002 | a2   | 2    | title2 | science    | abstract of article 2 | tags40      | author384  | zh       | text_a2.txt | image_a2_0.jpg,image_a2_1.jpg,                                              |       |
|   4 | 1506000000003 | a3   | 3    | title3 | technology | abstract of article 3 | tags9       | author1741 | en       | text_a3.txt | image_a3_0.jpg,                                                             |       |
|   5 | 1506000000004 | a4   | 4    | title4 | science    | abstract of article 4 | tags5       | author613  | zh       | text_a4.txt | image_a4_0.jpg,                                                             |       |
|   6 | 1506000000005 | a5   | 5    | title5 | technology | abstract of article 5 | tags43      | author1131 | en       | text_a5.txt | image_a5_0.jpg,image_a5_1.jpg,image_a5_2.jpg,image_a5_3.jpg,                |       |
|   7 | 1506000000006 | a6   | 6    | title6 | technology | abstract of article 6 | tags19      | author916  | en       | text_a6.txt | image_a6_0.jpg,image_a6_1.jpg,image_a6_2.jpg,image_a6_3.jpg,image_a6_4.jpg, |       |
+-----+---------------+------+------+--------+------------+-----------------------+-------------+------------+----------+-------------+-----------------------------------------------------------------------------+-------+
```

## Next step

The next step is to shard our newly created Vitess cluster. Which is detailled [in the next section](./setup-sharded-vitess-cluster.md).