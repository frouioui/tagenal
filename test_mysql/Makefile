
start_minikube			= minikube start --cpus=10 --memory=10000 --disk-size=30g

vitess_operator_folder		= ./vitess/examples/operator
vitess_common_folder		= ./vitess/examples/common

clone_vitess_repo = git clone git@github.com:vitessio/vitess.git

set_alias_mysql_client = alias mysql="mysql -h 127.0.0.1 -P 15306 -u user"
set_alias_vtctl_client = alias vtctlclient="vtctlclient -server=localhost:15999"

select_commerce_data	=	mysql --table < $(vitess_operator_folder)/select_commerce_data.sql

show_vitess_tablets		=	echo "show vitess_tablets;" | mysql --table

.PHONY: set_aliases

start_minikube:
	$(start_minikube)

clone_vitess_github:
	$(clone_vitess_github)

set_aliases:
	$(set_alias_mysql_client)
	$(set_alias_vtctl_client)

show_vitess_tablets:
	$(show_vitess_tablets)

select_commerce_data:
	$(select_commerce_data)