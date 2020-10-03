rs.initiate(
    {
        _id: "rs-shard-02",
        version: 1,
        members: [
            { _id: 0, host: "mongo-02-a:27017" },
            { _id: 1, host: "mongo-02-b:27017" },
            { _id: 2, host: "mongo-02-c:27017" },
        ]
    }
)