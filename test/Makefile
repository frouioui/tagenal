ROUTER_01	=	mongo-router-01

CONFIG_SERVER_01	=	mongo-config-01
CONFIG_SERVER_02	=	mongo-config-02

SHARD_01	= 	mongo-01-a
SHARD_02	= 	mongo-02-a

DCP_EXEC	=	docker-compose exec

PATH_SCRIPT	=	./scripts
PATH_DATA	=	./data

cmd_init_config_server	=	$(DCP_EXEC) $(CONFIG_SERVER_01) sh -c "mongo < /scripts/config/init-config-server.js"

cmd_init_shards			=	$(DCP_EXEC) $(SHARD_01) sh -c "mongo < /scripts/shard/init-shard-01.js" ; $(DCP_EXEC) $(SHARD_02) sh -c "mongo < /scripts/shard/init-shard-02.js"

cmd_init_router			=	$(DCP_EXEC) $(ROUTER_01) sh -c "mongo < /scripts/router/init-router.js"

init_shards:
	$(cmd_init_shards)

init_router:
	$(cmd_init_router)

init_all:
	$(cmd_init_config_server)
	$(cmd_init_shards)
	@echo "wait"
	sleep 10
	$(cmd_init_router)

router_status:
	$(DCP_EXEC) $(ROUTER_01) sh -c "mongo < /scripts/router/status.js"

router_setup_shard_tags:
	$(DCP_EXEC) $(ROUTER_01) sh -c "mongo < /scripts/router/init-shard-tags.js"

flush_db_data:
	rm -Rf $(PATH_DATA)/shard01/shard01a/* $(PATH_DATA)/shard01/shard01b/* $(PATH_DATA)/shard01/shard01c/*
	rm -Rf $(PATH_DATA)/shard02/shard02a/* $(PATH_DATA)/shard02/shard02b/* $(PATH_DATA)/shard02/shard02c/*

.PHONY: router_status