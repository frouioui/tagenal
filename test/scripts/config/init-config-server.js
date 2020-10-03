rs.initiate(
    {
        _id: "rs-config-server",
        configsvr: true,
        version: 1,
        members: [
            { _id: 0, host: "mongo-config-01:27017" },
            { _id: 1, host: "mongo-config-02:27017" },
            { _id: 2, host: "mongo-config-03:27017" },
        ]
    }
)