
V_KEYSPACE_USERS	= users
V_KEYSPACE_ARTICLES	= articles
V_KEYSPACE_CONFIG	= config

DATABASE_FOLDER_PATH	= ./database
KUBE_PROMETHEUS_PATH	= ./lib/kube-prometheus
VITESS_OPERATOR_PATH	= ./lib/vitess-operator/deploy

# Aliases
MYSQL_CLIENT	=	mysql -h tagenal -P 3000 -u user
VTCTL_CLIENT	=	vtctlclient -server=tagenal:8000

# Shards
## Users
SHARD_ALTER_USERS_TABLES_SQL				= $(VTCTL_CLIENT) ApplySchema -sql="$(shell cat $(DATABASE_FOLDER_PATH)/users/init/alter_users_auto_increment.sql)" $(V_KEYSPACE_USERS)
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

GET_USERS_VSCHEMA	= $(VTCTL_CLIENT) GetVSchema $(V_KEYSPACE_USERS)

list_vtctld:
	kubectl get pods --selector="planetscale.com/component=vtctld" -o custom-columns=":metadata.name"

start_minikube:
	minikube start --driver=hyperkit --mount --mount-string ./data:/mount_data --kubernetes-version=v1.19.2 --cpus=12 --memory=11500 --disk-size=80g --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook --extra-config=scheduler.address=0.0.0.0 --extra-config=controller-manager.address=0.0.0.0
	minikube addons disable metrics-server
	@sleep 15
	echo "$(shell pwd)/data -alldirs -mapall="$(shell id -u)":"$(shell id -g)" $(shell minikube ip)" | sudo tee -a /etc/exports && sudo nfsd restart

start_minikube_dashboard:
	minikube dashboard

clone_vitess_operator:
	./lib/get-vitess-operator.sh

setup_vitess_operator_kubernetes:
	kubectl apply -f kubernetes/vitess/vitess_namespace.yaml
	kubectl config set-context $(shell kubectl config current-context) --namespace=vitess
	kubectl apply -f $(VITESS_OPERATOR_PATH)/crds/
	kubectl apply -f $(VITESS_OPERATOR_PATH)/priority.yaml
	kubectl apply -f $(VITESS_OPERATOR_PATH)/role.yaml
	kubectl apply -f $(VITESS_OPERATOR_PATH)/role_binding.yaml
	kubectl apply -f $(VITESS_OPERATOR_PATH)/service_account.yaml
	kubectl apply -f kubernetes/vitess/vitess_operator.yaml
	kubectl config set-context $(shell kubectl config current-context) --namespace=default

setup_vitess_kubernetes:
	kubectl apply -f kubernetes/vitess/vitess_cluster_secret.yaml
	kubectl apply -f kubernetes/vitess/vitess_cluster_config.yaml
	kubectl apply -f kubernetes/vitess/vitess_cluster.yaml	
	kubectl apply -f kubernetes/vitess/vitess_vtgate_service.yaml

init_mysql_tables:
	$(VTCTL_CLIENT) ApplySchema -sql="$(shell cat $(DATABASE_FOLDER_PATH)/users/init/init_users.sql)" $(V_KEYSPACE_USERS)
	$(VTCTL_CLIENT) ApplySchema -sql="$(shell cat $(DATABASE_FOLDER_PATH)/articles/init/init_articles.sql)" $(V_KEYSPACE_ARTICLES)

init_increment_sequences:
	$(VTCTL_CLIENT) ApplySchema -sql="$(shell cat $(DATABASE_FOLDER_PATH)/config/init/init_increment_seq.sql)" $(V_KEYSPACE_CONFIG)
	$(VTCTL_CLIENT) ApplyVSchema -vschema='$(shell cat $(DATABASE_FOLDER_PATH)/config/vschema/vschema_config_seq.json)' $(V_KEYSPACE_CONFIG)

init_region_sharding_users:
	$(VTCTL_CLIENT) ApplyVSchema -vschema='$(shell cat $(DATABASE_FOLDER_PATH)/users/vschema/vschema_users_shard_region.json)' $(V_KEYSPACE_USERS)

init_region_sharding_articles:
	$(VTCTL_CLIENT) ApplyVSchema -vschema='$(shell cat $(DATABASE_FOLDER_PATH)/articles/vschema/vschema_articles_shard_region.json)' $(V_KEYSPACE_ARTICLES)

init_vitess_all:
	$(VTCTL_CLIENT) ApplySchema -sql="$(shell cat $(DATABASE_FOLDER_PATH)/users/init/init_users.sql)" $(V_KEYSPACE_USERS)
	$(VTCTL_CLIENT) ApplySchema -sql="$(shell cat $(DATABASE_FOLDER_PATH)/articles/init/init_articles.sql)" $(V_KEYSPACE_ARTICLES)
	$(VTCTL_CLIENT) ApplySchema -sql="$(shell cat $(DATABASE_FOLDER_PATH)/config/init/init_increment_seq.sql)" $(V_KEYSPACE_CONFIG)
	$(VTCTL_CLIENT) ApplyVSchema -vschema='$(shell cat $(DATABASE_FOLDER_PATH)/config/vschema/vschema_config_seq.json)' $(V_KEYSPACE_CONFIG)
	$(VTCTL_CLIENT) ApplyVSchema -vschema='$(shell cat $(DATABASE_FOLDER_PATH)/users/vschema/vschema_users_shard_region.json)' $(V_KEYSPACE_USERS)
	$(VTCTL_CLIENT) ApplyVSchema -vschema='$(shell cat $(DATABASE_FOLDER_PATH)/articles/vschema/vschema_articles_shard_region.json)' $(V_KEYSPACE_ARTICLES)

init_vreplication_articles:
	$(VTCTL_CLIENT) $(shell go run scripts/vreplgen.go articles_science '$(shell $(VTCTL_CLIENT) GetShard articles/80-)') 
	$(VTCTL_CLIENT) $(shell go run scripts/vreplgen.go popularity_science '$(shell $(VTCTL_CLIENT) GetShard articles/80-)') 

init_vreplication_article_be_read:
	$(VTCTL_CLIENT) $(shell go run scripts/vreplgen.go be_read_articles '$(shell $(VTCTL_CLIENT) GetShard articles/-80)' -80)
	$(VTCTL_CLIENT) $(shell go run scripts/vreplgen.go be_read_articles '$(shell $(VTCTL_CLIENT) GetShard articles/-80)' 80-)
	$(VTCTL_CLIENT) $(shell go run scripts/vreplgen.go be_read_articles '$(shell $(VTCTL_CLIENT) GetShard articles/80-)' -80)
	$(VTCTL_CLIENT) $(shell go run scripts/vreplgen.go be_read_articles '$(shell $(VTCTL_CLIENT) GetShard articles/80-)' 80-)

init_cronjobs_popularity:
	kubectl apply -f kubernetes/jobs/job_popularity.yaml

build_monitoring_manifests: $(shell chmod +x ./monitoring/build.sh)
	./monitoring/build.sh

run_monitoring: build_monitoring_manifests
	kubectl create -f $(KUBE_PROMETHEUS_PATH)/manifests/setup
	until kubectl get servicemonitors --all-namespaces ; do date; sleep 1; echo ""; done
	kubectl create -f $(KUBE_PROMETHEUS_PATH)/manifests/

setup_jaeger: $(shell chmod +x ./scripts/jaeger.sh)
	./scripts/jaeger.sh
	kubectl create -n observability -f ./kubernetes/jaeger/operator.yaml
	kubectl create -n observability -f ./kubernetes/jaeger/jaeger.yaml

setup_traefik:
	kubectl create -f kubernetes/traefik/traefik_crd.yaml
	kubectl create -f kubernetes/traefik/traefik_rbac.yaml
	@echo Wait 5s
	@sleep 5
	kubectl create -f kubernetes/traefik/traefik_deployment_jaeger.yaml
	kubectl create -n observability -f ./kubernetes/jaeger/jaeger_ui_ingress_route.yaml

setup_traefik_vitess: $(shell chmod +x ./kubernetes/traefik/vitess/build.sh)
	./kubernetes/traefik/vitess/build.sh
	kubectl create -f kubernetes/traefik/vitess/

setup_traefik_monitoring:
	kubectl create -f kubernetes/traefik/monitoring/

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

show_user_read_table:
	@$(MYSQL_CLIENT) --table < ./database/users/select/select_user_read.sql
	@$(MYSQL_CLIENT) --table < ./database/users/select/select_user_read_shard_1.sql
	@$(MYSQL_CLIENT) --table < ./database/users/select/select_user_read_shard_2.sql

show_article_table:
	@$(MYSQL_CLIENT) --table < ./database/articles/select/select_article.sql
	@$(MYSQL_CLIENT) --table < ./database/articles/select/select_article_shard_1.sql
	@$(MYSQL_CLIENT) --table < ./database/articles/select/select_article_shard_2.sql

show_article_keyspace:
	@$(MYSQL_CLIENT) --table < ./database/articles/select/select_article_all.sql

show_article_beread_table:
	@$(MYSQL_CLIENT) --table < ./database/articles/select/select_beread.sql
	@$(MYSQL_CLIENT) --table < ./database/articles/select/select_beread_shard_1.sql
	@$(MYSQL_CLIENT) --table < ./database/articles/select/select_beread_shard_2.sql


setup_redis: $(shell chmod +x ./kubernetes/redis/setup.sh)
	kubectl apply -f kubernetes/redis/redis.yaml

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
