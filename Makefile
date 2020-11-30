V_KEYSPACE_USERS	= users
V_KEYSPACE_ARTICLES	= articles
V_KEYSPACE_CONFIG	= config

DATABASE_FOLDER_PATH	= ./database

KUBE_PROMETHEUS_PATH	= ./lib/kube-prometheus

VITESS_OPERATOR_PATH	= ./lib/vitess/examples/operator
VITESS_COMMON_PATH		= ./lib/vitess/examples/common

# Aliases
MYSQL_CLIENT	=	mysql -h tagenal -P 3000 -u user
VTCTL_CLIENT	=	vtctlclient -server=tagenal:8000

# Unsharded commands
INIT_USERS_TABLES	= $(VTCTL_CLIENT) ApplySchema -sql="$(shell cat $(DATABASE_FOLDER_PATH)/users/init/init_users.sql)" $(V_KEYSPACE_USERS)
INIT_USERS_VSCHEMA	= $(VTCTL_CLIENT) ApplyVSchema -vschema='$(shell cat $(DATABASE_FOLDER_PATH)/users/vschema/vschema_users_initial.json)' $(V_KEYSPACE_USERS)

INIT_ARTICLES_TABLES	= $(VTCTL_CLIENT) ApplySchema -sql="$(shell cat $(DATABASE_FOLDER_PATH)/articles/init/init_articles.sql)" $(V_KEYSPACE_ARTICLES)
INIT_ARTICLES_VSCHEMA	= $(VTCTL_CLIENT) ApplyVSchema -vschema='$(shell cat $(DATABASE_FOLDER_PATH)/articles/vschema/vschema_articles_initial.json)' $(V_KEYSPACE_ARTICLES)

# Sharded commands
## Config
SHARD_INIT_CONFIG_SEQUENCES_SQL		= $(VTCTL_CLIENT) ApplySchema -sql="$(shell cat $(DATABASE_FOLDER_PATH)/config/init/init_increment_seq.sql)" $(V_KEYSPACE_CONFIG)
SHARD_INIT_CONFIG_SEQUENCES_VSCHEMA	= $(VTCTL_CLIENT) ApplyVSchema -vschema='$(shell cat $(DATABASE_FOLDER_PATH)/config/vschema/vschema_config_seq.json)' $(V_KEYSPACE_CONFIG)

## Users
SHARD_ALTER_USERS_TABLES_SQL				= $(VTCTL_CLIENT) ApplySchema -sql="$(shell cat $(DATABASE_FOLDER_PATH)/users/init/alter_user_auto_increment.sql)" $(V_KEYSPACE_USERS)
SHARD_INIT_USERS_VSCHEMA					= $(VTCTL_CLIENT) ApplyVSchema -vschema='$(shell cat $(DATABASE_FOLDER_PATH)/users/vschema/vschema_users_shard.json)' $(V_KEYSPACE_USERS)
SHARD_INIT_RESHARD_USERS					= $(VTCTL_CLIENT) Reshard $(V_KEYSPACE_USERS).user2user '-' '-80,80-'
SHARD_VERIFY_USERS_SHARDING_PROCESS			= $(VTCTL_CLIENT) VDiff $(V_KEYSPACE_USERS).user2user
SHARD_SWITCH_READ_REPLICA_USERS				= $(VTCTL_CLIENT) SwitchReads -tablet_type=replica $(V_KEYSPACE_USERS).user2user
SHARD_SWITCH_READ_RDONLY_USERS				= $(VTCTL_CLIENT) SwitchReads -tablet_type=rdonly $(V_KEYSPACE_USERS).user2user
SHARD_SWITCH_WRITE_USERS					= $(VTCTL_CLIENT) SwitchWrites $(V_KEYSPACE_USERS).user2user

## Articles
SHARD_ALTER_ARTICLES_TABLES_SQL				= $(VTCTL_CLIENT) ApplySchema -sql="$(shell cat $(DATABASE_FOLDER_PATH)/articles/init/alter_articles_auto_increment.sql)" $(V_KEYSPACE_ARTICLES)
SHARD_INIT_ARTICLES_VSCHEMA					= $(VTCTL_CLIENT) ApplyVSchema -vschema='$(shell cat $(DATABASE_FOLDER_PATH)/articles/vschema/vschema_articles_shard.json)' $(V_KEYSPACE_ARTICLES)
SHARD_INIT_RESHARD_ARTICLES					= $(VTCTL_CLIENT) Reshard $(V_KEYSPACE_ARTICLES).article2article '-' '-80,80-'
SHARD_VERIFY_ARTICLES_SHARDING_PROCESS		= $(VTCTL_CLIENT) VDiff $(V_KEYSPACE_ARTICLES).article2article
SHARD_SWITCH_READ_REPLICA_ARTICLES			= $(VTCTL_CLIENT) SwitchReads -tablet_type=replica $(V_KEYSPACE_ARTICLES).article2article
SHARD_SWITCH_READ_RDONLY_ARTICLES			= $(VTCTL_CLIENT) SwitchReads -tablet_type=rdonly $(V_KEYSPACE_ARTICLES).article2article
SHARD_SWITCH_WRITE_ARTICLES					= $(VTCTL_CLIENT) SwitchWrites $(V_KEYSPACE_ARTICLES).article2article
SHARD_REPLICATION_CATEGORY_ARTICLE			= $(shell go run scripts/vreplgen.go '$(shell $(VTCTL_CLIENT) GetShard articles/80-)') 

# Region sharding commands
## Users
REGION_SHARD_INIT_CONFIG_USERS_VSCHEMA			= $(VTCTL_CLIENT) ApplyVSchema -vschema='$(shell cat $(DATABASE_FOLDER_PATH)/users/vschema/region/vschema_users_shard_region.json)' $(V_KEYSPACE_USERS)
REGION_SHARD_INIT_CONFIG_USER_LOOKUP_VINDEX		= $(VTCTL_CLIENT) CreateLookupVindex -tablet_types=REPLICA $(V_KEYSPACE_USERS) '$(shell cat $(DATABASE_FOLDER_PATH)/users/vschema/region/vschema_users_shard_lookup_vindex.json)'
REGION_SHARD_EXTERNALIZE_USER_LOOKUP_VINDEX		= $(VTCTL_CLIENT) ExternalizeVindex $(V_KEYSPACE_USERS).user_region_lookup

## Articles
REGION_SHARD_INIT_CONFIG_ARTICLES_VSCHEMA			= $(VTCTL_CLIENT) ApplyVSchema -vschema='$(shell cat $(DATABASE_FOLDER_PATH)/articles/vschema/region/vschema_articles_shard_region.json)' $(V_KEYSPACE_ARTICLES)
REGION_SHARD_INIT_CONFIG_ARTICLE_LOOKUP_VINDEX		= $(VTCTL_CLIENT) CreateLookupVindex -tablet_types=REPLICA $(V_KEYSPACE_ARTICLES) '$(shell cat $(DATABASE_FOLDER_PATH)/articles/vschema/region/vschema_articles_shard_lookup_vindex.json)'
REGION_SHARD_EXTERNALIZE_ARTICLE_LOOKUP_VINDEX		= $(VTCTL_CLIENT) ExternalizeVindex $(V_KEYSPACE_ARTICLES).article_region_lookup

GET_USERS_VSCHEMA	= $(VTCTL_CLIENT) GetVSchema $(V_KEYSPACE_USERS)

list_vtctld:
	kubectl get pods --selector="planetscale.com/component=vtctld" -o custom-columns=":metadata.name"

start_minikube:
	minikube start --driver=hyperkit --kubernetes-version=v1.19.2 --cpus=10 --memory=10000 --disk-size=80g --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook --extra-config=scheduler.address=0.0.0.0 --extra-config=controller-manager.address=0.0.0.0
	minikube addons disable metrics-server

start_minikube_dashboard:
	minikube dashboard

clone_vitess_github:
	./lib/script/get-vitess.sh

install_vitess_operator:
	kubectl apply -f kubernetes/vitess_namespace.yaml
	kubectl config set-context $(shell kubectl config current-context) --namespace=vitess
	kubectl apply -f $(VITESS_OPERATOR_PATH)/operator.yaml
	kubectl config set-context $(shell kubectl config current-context) --namespace=default

init_kubernetes_unsharded_database:
	kubectl apply -f kubernetes/vitess_cluster_secret.yaml
	kubectl apply -f kubernetes/vitess_cluster_config.yaml
	kubectl apply -f kubernetes/init_cluster_vitess.yaml

port_forwarding_vitess:
	./script/port_forwarding.sh

init_unsharded_database:
	$(INIT_USERS_TABLES)
	$(INIT_USERS_VSCHEMA)
	$(INIT_ARTICLES_TABLES)
	$(INIT_ARTICLES_VSCHEMA)

init_config_increment_sequence:
	$(SHARD_INIT_CONFIG_SEQUENCES_SQL)
	$(SHARD_INIT_CONFIG_SEQUENCES_VSCHEMA)

	$(SHARD_ALTER_USERS_TABLES_SQL)
	$(SHARD_INIT_USERS_VSCHEMA)
	
	$(SHARD_ALTER_ARTICLES_TABLES_SQL)
	$(SHARD_INIT_ARTICLES_VSCHEMA)

init_sharded_database:
	kubectl apply -f kubernetes/init_cluster_vitess_sharded.yaml

init_region_sharding_users:
	$(REGION_SHARD_INIT_CONFIG_USERS_VSCHEMA)
	$(REGION_SHARD_INIT_CONFIG_USER_LOOKUP_VINDEX)
	@echo Wait ...
	@sleep 5
	$(REGION_SHARD_EXTERNALIZE_USER_LOOKUP_VINDEX)

init_region_sharding_articles:
	$(REGION_SHARD_INIT_CONFIG_ARTICLES_VSCHEMA)
	$(REGION_SHARD_INIT_CONFIG_ARTICLE_LOOKUP_VINDEX)
	@echo Wait ...
	@sleep 5
	$(REGION_SHARD_EXTERNALIZE_ARTICLE_LOOKUP_VINDEX)

init_region_sharding:
	$(init_region_sharding_users)
	$(init_region_sharding_articles)

resharding_process_users:
	$(SHARD_INIT_RESHARD_USERS)
	$(SHARD_VERIFY_USERS_SHARDING_PROCESS)
	$(SHARD_SWITCH_READ_REPLICA_USERS)
	$(SHARD_SWITCH_READ_RDONLY_USERS)
	$(SHARD_SWITCH_WRITE_USERS)

resharding_process_articles:
	$(SHARD_INIT_RESHARD_ARTICLES)
	$(SHARD_VERIFY_ARTICLES_SHARDING_PROCESS)
	$(SHARD_SWITCH_READ_REPLICA_ARTICLES)
	$(SHARD_SWITCH_READ_RDONLY_ARTICLES)
	$(SHARD_SWITCH_WRITE_ARTICLES)

resharding_process:
	$(resharding_process_users)
	$(resharding_process_articles)

init_vreplication_articles:
	$(SHARD_REPLICATION_CATEGORY_ARTICLE)

final_vitess_cluster:
	kubectl apply -f kubernetes/init_cluster_vitess_sharded_final.yaml

build_monitoring_manifests: $(shell chmod +x ./monitoring/build.sh)
	./monitoring/build.sh

run_monitoring: build_monitoring_manifests
	kubectl create -f $(KUBE_PROMETHEUS_PATH)/manifests/setup
	until kubectl get servicemonitors --all-namespaces ; do date; sleep 1; echo ""; done
	kubectl create -f $(KUBE_PROMETHEUS_PATH)/manifests/

setup_traefik:
	kubectl create -f kubernetes/traefik/traefik_crd.yaml
	kubectl create -f kubernetes/traefik/traefik_rbac.yaml
	@echo Wait 5s
	@sleep 5
	kubectl create -f kubernetes/traefik/traefik_deployment.yaml

setup_traefik_vitess: $(shell chmod +x ./kubernetes/traefik/vitess/build.sh)
	./kubernetes/traefik/vitess/build.sh
	kubectl create -f kubernetes/traefik/vitess/

setup_traefik_monitoring:
	kubectl create -f kubernetes/traefik/monitoring/

copy_locations_json_to_k8s:
	kubectl cp ./database/locations/locations.json $(shell kubectl get pods --selector="planetscale.com/component=vtctld" -o custom-columns=":metadata.name"):/tmp/countries.json
	kubectl cp ./database/locations/locations.json $(shell kubectl get pods --selector="planetscale.com/component=vtgate" -o custom-columns=":metadata.name"):/tmp/countries.json

show_vttablets:
	kubectl get pods --namespace=vitess --selector="planetscale.com/component=vttablet" -o custom-columns=":metadata.name" 

show_vitess_tablets:
	echo "show vitess_tablets;" | $(MYSQL_CLIENT) --table

insert_few_user_row:
	$(MYSQL_CLIENT) < ./database/users/insert/insert_data_users.sql

insert_few_article_row:
	$(MYSQL_CLIENT) < ./database/articles/insert/insert_data_article.sql

show_user_table:
	@$(MYSQL_CLIENT) --table < ./database/users/select/select_user.sql
	@$(MYSQL_CLIENT) --table < ./database/users/select/select_user_shard_1.sql
	@$(MYSQL_CLIENT) --table < ./database/users/select/select_user_shard_2.sql

show_article_table:
	@$(MYSQL_CLIENT) --table < ./database/articles/select/select_article.sql
	@$(MYSQL_CLIENT) --table < ./database/articles/select/select_article_shard_1.sql
	@$(MYSQL_CLIENT) --table < ./database/articles/select/select_article_shard_2.sql

build_push_apis:
	make -C ./api/users/
	make -C ./api/articles/

build_apis:
	make -C ./api/users/ protobuild dockerbuild 
	make -C ./api/articles/ protobuild dockerbuild

build_push_frontend:
	make -C ./frontend/

build_frontend:
	make -C ./frontend/ dockerbuild 

run_apis_k8s:
	kubectl apply -f ./kubernetes/api/users/
	kubectl apply -f ./kubernetes/api/articles/

run_frontend_k8s:
	kubectl apply -f ./kubernetes/frontend/

stop_apis_k8s:
	kubectl delete -f ./kubernetes/api/users/
	kubectl delete -f ./kubernetes/api/articles/

stop_frontend_k8s:
	kubectl delete -f ./kubernetes/frontend/

.PHONY: set_aliases
