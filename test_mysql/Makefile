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
SHARD_SWITCH_READ_RDONLY_USERS				= $(VTCTL_CLIENT) SwitchReads -tablet_type=rdonly $(V_KEYSPACE_USERS).user2user
SHARD_SWITCH_WRITE_USERS					= $(VTCTL_CLIENT) SwitchWrites $(V_KEYSPACE_USERS).user2user

# Region sharding commands
REGION_SHARD_INIT_CONFIG_USERS_VSCHEMA			= $(VTCTL_CLIENT) ApplyVSchema -vschema='$(shell cat $(DATABASE_FOLDER_PATH)/vschema/region/vschema_users_shard_region.json)' $(V_KEYSPACE_USERS)
REGION_SHARD_INIT_CONFIG_USER_LOOKUP_VINDEX		= $(VTCTL_CLIENT) CreateLookupVindex -tablet_types=REPLICA $(V_KEYSPACE_USERS) '$(shell cat $(DATABASE_FOLDER_PATH)/vschema/region/vschema_users_shard_lookup_vindex.json)'
REGION_SHARD_EXTERNALIZE_USER_LOOKUP_VINDEX		= $(VTCTL_CLIENT) ExternalizeVindex $(V_KEYSPACE_USERS).user_region_lookup

GET_USERS_VSCHEMA	= $(VTCTL_CLIENT) GetVSchema $(V_KEYSPACE_USERS)

list_vtctld:
	kubectl get pods --selector="planetscale.com/component=vtctld" -o custom-columns=":metadata.name"

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

init_users_increment_sequence:
	$(SHARD_INIT_CONFIG_USERS_SEQUENCES_SQL)
	$(SHARD_INIT_CONFIG_USERS_SEQUENCES_VSCHEMA)
	$(INIT_USERS_VSCHEMA)

init_sharded_database:
	$(SHARD_INIT_CONFIG_USERS_SEQUENCES_SQL)
	$(SHARD_INIT_CONFIG_USERS_SEQUENCES_VSCHEMA)
	$(SHARD_ALTER_USERS_TABLES_SQL)
	$(SHARD_INIT_USERS_VSCHEMA)
	kubectl apply -f kubernetes/init_cluster_vitess_sharded.yaml

init_region_sharding_users: copy_locations_json_to_k8s
	# kubectl cp ./database/locations/locations.json vitess-vttablet-zone1-3724453505-5a510085:/tmp/countries.json
	# kubectl cp ./database/locations/locations.json vitess-vttablet-zone1-3774659171-51411321:/tmp/countries.json
	# kubectl cp ./database/locations/locations.json vitess-vttablet-zone1-3890619061-d6d202c2:/tmp/countries.json
	# kubectl cp ./database/locations/locations.json vitess-vttablet-zone1-0808196278-480a1110:/tmp/countries.json
	kubectl cp ./database/locations/locations.json vitess-vttablet-zone1-1651728703-a7361ca4:/tmp/countries.json
	kubectl cp ./database/locations/locations.json vitess-vttablet-zone1-1929229597-af5a9c25:/tmp/countries.json
	kubectl cp ./database/locations/locations.json vitess-vttablet-zone1-4109238961-30b1e68c:/tmp/countries.json
	kubectl cp ./database/locations/locations.json vitess-vttablet-zone1-4224402714-96c8f9a7:/tmp/countries.json
	$(REGION_SHARD_INIT_CONFIG_USERS_VSCHEMA)
	$(REGION_SHARD_INIT_CONFIG_USER_LOOKUP_VINDEX)
	@echo Wait ...
	@sleep 5
	$(REGION_SHARD_EXTERNALIZE_USER_LOOKUP_VINDEX)

resharding_process_users:
	$(SHARD_INIT_RESHARD_USERS)
	$(SHARD_VERIFY_USERS_SHARDING_PROCESS)
	$(SHARD_SWITCH_READ_REPLICA_USERS)
	$(SHARD_SWITCH_READ_RDONLY_USERS)
	$(SHARD_SWITCH_WRITE_USERS)

reshard_users_db:
	$(SHARD_INIT_RESHARD_USERS)

copy_locations_json_to_k8s:
	kubectl cp ./database/locations/locations.json $(shell kubectl get pods --selector="planetscale.com/component=vtctld" -o custom-columns=":metadata.name"):/tmp/countries.json
	kubectl cp ./database/locations/locations.json $(shell kubectl get pods --selector="planetscale.com/component=vtgate" -o custom-columns=":metadata.name"):/tmp/countries.json

show_vttablets:
	kubectl get pods --selector="planetscale.com/component=vttablet" -o custom-columns=":metadata.name" 

show_vitess_tablets:
	echo "show vitess_tablets;" | $(MYSQL_CLIENT) --table

insert_few_user_row:
	$(MYSQL_CLIENT) < ./database/insert/insert_data_users.sql

show_user_table:
	@$(MYSQL_CLIENT) --table < ./database/select/select_user.sql
	@$(MYSQL_CLIENT) --table < ./database/select/select_user_shard_1.sql
	@$(MYSQL_CLIENT) --table < ./database/select/select_user_shard_2.sql

.PHONY: set_aliases
