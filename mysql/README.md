# Tagenal MySQL Cluster

Tagenal's MySQL cluster relies on Vitess and Kubernetes.

## Setup

### **1.** Start minikube

`make start_minikube`

### **2.** Start kubernetes dashboard (optional)

This command should be run in second terminal.

`make start_minikube_dashboard`

### **3.** Install vitess operator in kubernetes

`make install_vitess_operator`

Expected output:

```
kubectl apply -f ./vitess/examples/operator/operator.yaml
customresourcedefinition.apiextensions.k8s.io/etcdlockservers.planetscale.com created
customresourcedefinition.apiextensions.k8s.io/vitessbackups.planetscale.com created
customresourcedefinition.apiextensions.k8s.io/vitessbackupstorages.planetscale.com created
customresourcedefinition.apiextensions.k8s.io/vitesscells.planetscale.com created
customresourcedefinition.apiextensions.k8s.io/vitessclusters.planetscale.com created
customresourcedefinition.apiextensions.k8s.io/vitesskeyspaces.planetscale.com created
customresourcedefinition.apiextensions.k8s.io/vitessshards.planetscale.com created
serviceaccount/vitess-operator created
role.rbac.authorization.k8s.io/vitess-operator created
rolebinding.rbac.authorization.k8s.io/vitess-operator created
priorityclass.scheduling.k8s.io/vitess created
priorityclass.scheduling.k8s.io/vitess-operator-control-plane created
deployment.apps/vitess-operator created
```

### **4.** Initialize unsharded database in kubernetes

`make init_kubernetes_unsharded_database`

Expected output:

```
kubectl apply -f kubernetes/vitess_cluster_secret.yaml
secret/vitess-cluster-secret created
kubectl apply -f kubernetes/vitess_cluster_config.yaml
configmap/vitess-cluster-config-sharding created
kubectl apply -f kubernetes/init_cluster_vitess.yaml
vitesscluster.planetscale.com/vitess created
```

### **5.** Initialize unsharded database in kubernetes

This command should be run in a third terminal.

`make port_forwarding_vitess`

Expected output:

```
./script/port_forwarding.sh
Forwarding from 127.0.0.1:15000 -> 15000
Forwarding from [::1]:15000 -> 15000
Forwarding from 127.0.0.1:15999 -> 15999
Forwarding from [::1]:15999 -> 15999
Forwarding from 127.0.0.1:15306 -> 3306
Forwarding from [::1]:15306 -> 3306
```

After this command, Vitess cluster dashboard (vtctld) should be accessible at http://localhost:15000.


### **6.** Initialize MySQL user table and Vitess VSchema

`make init_unsharded_database`

Expected output:

```
vtctlclient -server=localhost:15999 ApplySchema -sql="DROP TABLE IF EXISTS user;  CREATE TABLE user (   _id INT NOT NULL auto_increment,   timestamp CHAR(14) DEFAULT NULL,   id CHAR(5) DEFAULT NULL,   uid CHAR(5) DEFAULT NULL,   name CHAR(9) DEFAULT NULL,   gender CHAR(7) DEFAULT NULL,   email CHAR(10) DEFAULT NULL,   phone CHAR(10) DEFAULT NULL,   dept CHAR(9) DEFAULT NULL,   grade CHAR(7) DEFAULT NULL,   language CHAR(3) DEFAULT NULL,   region VARBINARY(256),   role CHAR(6) DEFAULT NULL,   preferTags CHAR(7) DEFAULT NULL,   obtainedCredits CHAR(3) DEFAULT NULL,   PRIMARY KEY(_id) ) ENGINE=InnoDB DEFAULT CHARSET=utf8;" users
vtctlclient -server=localhost:15999 ApplyVSchema -vschema='{     "tables": {         "user": {}     } }' users
New VSchema object:
{
  "tables": {
    "user": {

    }
  }
}
If this is not what you expected, check the input data (as JSON parsing will skip unexpected fields).
```

### **7.** Initialize the incrementation sequence of the user table in the config keyspace

`make init_users_increment_sequence`

Expected output:

```
vtctlclient -server=localhost:15999 ApplySchema -sql="CREATE TABLE users_seq(id INT, next_id BIGINT, cache BIGINT, PRIMARY KEY(id)) comment 'vitess_sequence'; INSERT INTO users_seq(id, next_id, cache) VALUES(0, 1, 100);" config
vtctlclient -server=localhost:15999 ApplyVSchema -vschema='{     "tables": {         "users_seq": {             "type": "sequence"         }     } }' config
New VSchema object:
{
  "tables": {
    "users_seq": {
      "type": "sequence"
    }
  }
}
If this is not what you expected, check the input data (as JSON parsing will skip unexpected fields).
vtctlclient -server=localhost:15999 ApplyVSchema -vschema='{     "tables": {         "user": {}     } }' users
New VSchema object:
{
  "tables": {
    "user": {

    }
  }
}
If this is not what you expected, check the input data (as JSON parsing will skip unexpected fields).
vtctlclient -server=localhost:15999 ApplySchema -sql="ALTER TABLE user change _id _id INT NOT NULL;" users
vtctlclient -server=localhost:15999 ApplyVSchema -vschema='{     "sharded": true,     "vindexes": {         "hash": {             "type": "hash"         }     },     "tables": {         "user": {             "column_vindexes": [                 {                     "column": "_id",                     "name": "hash"                 }             ],             "auto_increment": {                 "column": "_id",                 "sequence": "users_seq"             }         }     } }' users
New VSchema object:
{
  "sharded": true,
  "vindexes": {
    "hash": {
      "type": "hash"
    }
  },
  "tables": {
    "user": {
      "columnVindexes": [
        {
          "column": "_id",
          "name": "hash"
        }
      ],
      "autoIncrement": {
        "column": "_id",
        "sequence": "users_seq"
      }
    }
  }
}
If this is not what you expected, check the input data (as JSON parsing will skip unexpected fields).
```

### **8.** Initialize sharded database in kubernetes

`make init_sharded_database`

Expected output:

```
Wait ...
kubectl apply -f kubernetes/init_cluster_vitess_sharded.yaml
vitesscluster.planetscale.com/vitess configured
```

### **9.** Kill previous port forwarding processes

In the third terminal, where the `make port_forwarding_vitess` was run, we will run the following command:

`ps`

And expect at least these two lines in the output, (the content of the lines can slightly differ):

```
38950 ttys000    0:00.52 kubectl port-forward service/vitess-vtctld-047eb1a6 15000 15999
38951 ttys000    0:00.42 kubectl port-forward service/vitess-vtgate-e05d80f4 15306:3306
```

We will take the 2 foremost left numbers, in this case `38950` and `38951`, and run the following command:

`kill 38950 38951`

Finally, we can run the `make port_forwarding_vitess` again.

### **10.** Insert the first few rows of the user table

`make insert_few_user_row`

Expected output:

```
mysql -h 127.0.0.1 -P 15306 -u user < ./database/insert/insert_data_users.sql
```

### **11.** Setup the region-based sharding

`make init_region_sharding_users`

Expected output:

```
vtctlclient -server=localhost:15999 ApplyVSchema -vschema='{   "sharded": true,   "vindexes": {     "region_vdx": {       "type": "region_json",       "params": {         "region_map": "/tmp/countries.json",         "region_bytes": "1"       }     }   },   "tables": {     "user": {       "column_vindexes": [         {           "columns": [             "_id",             "region"           ],           "name": "region_vdx"         }       ],       "auto_increment": {         "column": "_id",         "sequence": "users_seq"       }     }   } }' users
New VSchema object:
{
  "sharded": true,
  "vindexes": {
    "region_vdx": {
      "type": "region_json",
      "params": {
        "region_bytes": "1",
        "region_map": "/tmp/countries.json"
      }
    }
  },
  "tables": {
    "user": {
      "columnVindexes": [
        {
          "name": "region_vdx",
          "columns": [
            "_id",
            "region"
          ]
        }
      ],
      "autoIncrement": {
        "column": "_id",
        "sequence": "users_seq"
      }
    }
  }
}
If this is not what you expected, check the input data (as JSON parsing will skip unexpected fields).
vtctlclient -server=localhost:15999 CreateLookupVindex -tablet_types=REPLICA users '{   "sharded": true,   "vindexes": {     "user_region_lookup": {       "type": "consistent_lookup_unique",       "params": {         "table": "users.user_lookup",         "from": "_id",         "to": "keyspace_id"       },       "owner": "user"     }   },   "tables": {     "user": {       "column_vindexes": [         {           "column": "_id",           "name": "user_region_lookup"         }       ]     }   } }'
Wait ...
vtctlclient -server=localhost:15999 ExternalizeVindex users.user_region_lookup
```

### **12.** Start the resharding process of the cluster

`make resharding_process_users`

Expected output:

```
vtctlclient -server=localhost:15999 Reshard users.user2user '-' '-80,80-'
vtctlclient -server=localhost:15999 VDiff users.user2user
Summary for user: {ProcessedRows:5 MatchingRows:5 MismatchedRows:0 ExtraRowsSource:0 ExtraRowsTarget:0}
E1101 20:23:50.599068   52103 main.go:64] E1101 19:23:50.595298 vdiff.go:780] Draining extra row(s) found on the source starting with: [INT32(1) VARBINARY("\x01\x16k@\xb4J\xbaK\xd6")]
Summary for user_lookup: {ProcessedRows:5 MatchingRows:0 MismatchedRows:0 ExtraRowsSource:5 ExtraRowsTarget:0}
vtctlclient -server=localhost:15999 SwitchReads -tablet_type=replica users.user2user
vtctlclient -server=localhost:15999 SwitchReads -tablet_type=rdonly users.user2user
vtctlclient -server=localhost:15999 SwitchWrites users.user2user
```

### **13.** Test, display the user table

`make show_user_table`

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