sh.removeShardTag("rs-shard-01", "BEI");
sh.removeShardTag("rs-shard-02", "HKG");

sh.addShardTag("rs-shard-01", "BEI");
sh.addShardTag("rs-shard-02", "HKG");

sh.disableBalancing("tagenal.users");

sh.enableSharding("tagenal");

sh.shardCollection("tagenal.users", { region: 1, uid: 1 });
sh.addTagRange(
    "tagenal.users",
    { "region": "Beijing", "uid": MinKey },
    { "region": "Beijing", "uid": MaxKey },
    "BEI"
);
sh.addTagRange(
    "tagenal.users",
    { "region": "Hong Kong", "uid": MinKey },
    { "region": "Hong Kong", "uid": MaxKey },
    "HKG"
);

sh.enableBalancing("tagenal.users");
sh.startBalancer();