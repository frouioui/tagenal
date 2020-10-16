
start_minikube			= minikube start --cpus=10 --memory=10000 --disk-size=30g

vitess_operator_folder		= ./vitess/examples/operator
vitess_common_folder		= ./vitess/examples/common

clone_vitess_repo = git clone git@github.com:vitessio/vitess.git


mysql_client_alias	=	mysql -h 127.0.0.1 -P 15306 -u user
vtctl_client_alias	=	vtctlclient -server=localhost:15999

show_vitess_tablets			=	echo "show vitess_tablets;" | $(mysql_client_alias) --table
show_user_table_describe	=	echo "describe user;" | $(mysql_client_alias) --table

init_SQLSchema_usersdb		= $(vtctl_client_alias) ApplySchema -sql='$(shell cat database/init_users.sql)' usersdb
init_VSchema_usersdb		= $(vtctl_client_alias) ApplyVSchema -vschema='$(shell cat database/vschema_database_users.json)' usersdb

# Init UsersDB sequences (SQL / VSchema)
init_SQLSchema_usersdb_seq	= $(vtctl_client_alias) ApplySchema -sql='$(shell cat database/init_usersdb_seq.sql)' usersdb
init_VSchema_usersdb_seq	= $(vtctl_client_alias) ApplyVSchema -vschema='$(shell cat database/vschema_database_users_seq.json)' usersdb

# Init UsersDB sharding rules (VSchema)
init_VSchema_usersdb_shard	= $(vtctl_client_alias) ApplyVSchema -vschema='$(shell cat database/vschema_database_users_shard.json)' usersdb

.PHONY: set_aliases

dashboard_minikube:
	minikube dashboard &

start_minikube:
	$(start_minikube)

init_database:
	$(init_SQLSchema_usersdb)
	$(init_VSchema_usersdb)
	@echo "Wait, now sharding ..."
	@sleep 4
	$(init_SQLSchema_usersdb_seq)
	$(init_VSchema_usersdb_seq)
	$(init_VSchema_usersdb_shard)

clone_vitess_github:
	$(clone_vitess_github)

show_vitess_tablets:
	$(show_vitess_tablets)

show_user_table_describe:
	$(show_user_table_describe)