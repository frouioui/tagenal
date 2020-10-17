##
## Tagenal
## Tsinghua Univeristy
##

V_KEYSPACE_USERS	= users
V_KEYSPACE_CONFIG	= config

DATABASE_FOLDER_PATH	= ./database

VITESS_OPERATOR_PATH	= ./vitess/examples/operator
VITESS_COMMON_PATH		= ./vitess/examples/common

# Aliases
MYSQL_CLIENT	=	mysql -h 127.0.0.1 -P 15306 -u user
VTCTL_CLIENT	=	vtctlclient -server=localhost:15999

# Unsharded commands
INIT_USERS_TABLES	= $(VTCTL_CLIENT) ApplySchema -sql="$(shell cat $(DATABASE_FOLDER_PATH)/init/init_users.sql)" $(V_KEYSPACE_USERS)
INIT_USERS_VSCHEMA	= $(VTCTL_CLIENT) ApplyVSchema -vschema='$(shell cat $(DATABASE_FOLDER_PATH)/vschema/vschema_users_initial.json)' $(V_KEYSPACE_USERS)

# Sharded commands
SHARD_INIT_CONFIG_USERS_SEQUENCES_SQL		= $(VTCTL_CLIENT) ApplySchema -sql="$(shell cat $(DATABASE_FOLDER_PATH)/init/init_users_seq.sql)" $(V_KEYSPACE_CONFIG)
SHARD_INIT_CONFIG_USERS_SEQUENCES_VSCHEMA	= $(VTCTL_CLIENT) ApplyVSchema -vschema='$(shell cat $(DATABASE_FOLDER_PATH)/vschema/vschema_config_seq.json)' $(V_KEYSPACE_CONFIG)
SHARD_ALTER_USERS_TABLES_SQL				= $(VTCTL_CLIENT) ApplySchema -sql="$(shell cat $(DATABASE_FOLDER_PATH)/init/alter_users_tables_sharding.sql)" $(V_KEYSPACE_USERS)
SHARD_INIT_USERS_VSCHEMA					= $(VTCTL_CLIENT) ApplyVSchema -vschema='$(shell cat $(DATABASE_FOLDER_PATH)/vschema/vschema_users_shard.json)' $(V_KEYSPACE_USERS)
SHARD_INIT_RESHARD_USERS					= $(VTCTL_CLIENT) Reshard $(V_KEYSPACE_USERS).user2user '-' '-80,80-'
SHARD_VERIFY_USERS_SHARDING_PROCESS			= $(VTCTL_CLIENT) VDiff $(V_KEYSPACE_USERS).user2user
SHARD_SWITCH_READ_REPLICA_USERS				= $(VTCTL_CLIENT) SwitchReads -tablet_type=replica $(V_KEYSPACE_USERS).user2user
SHARD_SWITCH_WRITE_USERS					= $(VTCTL_CLIENT) SwitchWrites $(V_KEYSPACE_USERS).user2user


start_minikube:
	minikube start --cpus=12 --memory=14000 --disk-size=80g

start_minikube_dashboard:
	minikube dashboard

clone_vitess_github:
	git clone git@github.com:vitessio/vitess.git

install_vitess_operator:
	kubectl apply -f $(VITESS_OPERATOR_PATH)/operator.yaml

init_kubernetes_unsharded_database:
	kubectl apply -f kubernetes/init_cluster_vitess.yaml

port_forwarding_vitess:
	./script/port_forwarding.sh

init_unsharded_database:
	$(INIT_USERS_TABLES)
	$(INIT_USERS_VSCHEMA)

init_sharded_database:
	$(SHARD_INIT_CONFIG_USERS_SEQUENCES_SQL)
	$(SHARD_INIT_CONFIG_USERS_SEQUENCES_VSCHEMA)
	$(SHARD_ALTER_USERS_TABLES_SQL)
	$(SHARD_INIT_USERS_VSCHEMA)
	kubectl apply -f kubernetes/init_cluster_vitess_sharded.yaml

resharding_process_users:
	$(SHARD_INIT_RESHARD_USERS)
	$(SHARD_VERIFY_USERS_SHARDING_PROCESS)
	$(SHARD_SWITCH_READ_REPLICA_USERS)
	$(SHARD_SWITCH_WRITE_USERS)

reshard_users_db:
	$(SHARD_INIT_RESHARD_USERS)

show_vitess_tablets:
	echo "show vitess_tablets;" | $(MYSQL_CLIENT) --table

insert_few_user_row:
	$(MYSQL_CLIENT) < ./database/insert/insert_data_users.sql

show_user_table:
	@$(MYSQL_CLIENT) --table < ./database/select/select_user.sql
	@$(MYSQL_CLIENT) --table < ./database/select/select_user_shard_1.sql
	@$(MYSQL_CLIENT) --table < ./database/select/select_user_shard_2.sql

.PHONY: set_aliases