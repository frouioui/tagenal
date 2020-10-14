
start_minikube			= minikube start --cpus=10 --memory=10000 --disk-size=30g

vitess_operator_folder		= ./vitess/examples/operator
vitess_common_folder		= ./vitess/examples/common

clone_vitess_repo = git clone git@github.com:vitessio/vitess.git


mysql_client_alias	=	mysql -h 127.0.0.1 -P 15306 -u user
vtctl_client_alias	=	vtctlclient -server=localhost:15999


show_vitess_tablets			=	echo "show vitess_tablets;" | $(mysql_client_alias) --table
show_user_table_describe	=	echo "describe user;" | $(mysql_client_alias) --table

init_Schema_users_table	= $(vtctl_client_alias) ApplySchema -sql="$(cat database/init_users.sql)" usersdb
init_VSchema_usersdb	= $(vtctl_client_alias) ApplyVSchema -vschema="$(cat database/vschema_database_users.json)" usersdb

.PHONY: set_aliases

start_minikube:
	$(start_minikube)

clone_vitess_github:
	$(clone_vitess_github)

show_vitess_tablets:
	$(show_vitess_tablets)

show_user_table_describe:
	$(show_user_table_describe)