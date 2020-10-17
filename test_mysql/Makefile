##
## Tagenal
## Tsinghua Univeristy
##

start_minikube			= minikube start --cpus=12 --memory=14000 --disk-size=80g

vitess_operator_folder		= ./vitess/examples/operator
vitess_common_folder		= ./vitess/examples/common

clone_vitess_repo = git clone git@github.com:vitessio/vitess.git


## Aliases
mysql_client_alias	=	mysql -h 127.0.0.1 -P 15306 -u user
vtctl_client_alias	=	vtctlclient -server=localhost:15999


## SQL Commands
show_vitess_tablets			=	echo "show vitess_tablets;" | $(mysql_client_alias) --table
show_user_table_describe	=	echo "use users; describe user;" | $(mysql_client_alias) --table
select_user_table			=	echo "use users; select * from user;" | $(mysql_client_alias) --table


## Init of UsersDB
init_SQLSchema_usersdb		= $(vtctl_client_alias) ApplySchema -sql='$(shell cat database/init_users.sql)' users
init_VSchema_usersdb		= $(vtctl_client_alias) ApplyVSchema -vschema='$(shell cat database/vschema_database_users.json)' users


## Init sequences and sharding of UsersDB
init_SQLSchema_usersdb_seq	= $(vtctl_client_alias) ApplySchema -sql='$(shell cat database/alter_users_sharding.sql)' usersdb
init_VSchema_usersdb_shard	= $(vtctl_client_alias) ApplyVSchema -vschema='$(shell cat database/vschema_database_users_shard.json)' usersdb


start_resharding_process_usersdb	= $(vtctl_client_alias) Reshard usersdb.users2users "-" "-80,80-"


.PHONY: set_aliases

dashboard_minikube:
	minikube dashboard &

start_minikube:
	$(start_minikube)

install_operator:
	kubectl apply -f $(vitess_operator_folder)/operator.yaml

init_database:
	$(init_SQLSchema_usersdb)
	$(init_VSchema_usersdb)
	@echo "Wait, now sharding ..."
	@sleep 4
	$(init_SQLSchema_usersdb_seq)
	$(init_VSchema_usersdb_seq)
	$(init_VSchema_usersdb_shard)

reshard_users_db:
	$(start_resharding_process_usersdb)

clone_vitess_github:
	$(clone_vitess_github)

show_vitess_tablets:
	$(show_vitess_tablets)

show_user_table:
	$(select_user_table)
	@echo "\n"
	$(show_user_table_describe)