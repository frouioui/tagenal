package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"vitess.io/vitess/go/sqltypes"
	binlogdatapb "vitess.io/vitess/go/vt/proto/binlogdata"
)

type aliasInfo struct {
	Cell string `json:"cell"`
	UID  int    `json:"uid"`
}

type shardInfo struct {
	MasterAlias aliasInfo `json:"master_alias"`
}

func replicationReadStats() (*binlogdatapb.BinlogSource, string) {
	shard := os.Args[3]
	if shard == "" {
		panic(errors.New("missing source shard info"))
	}
	dbName := "users"

	filter := &binlogdatapb.Filter{
		Rules: []*binlogdatapb.Rule{{
			Match:  "read_stats",
			Filter: "SELECT aid AS id, 0 AS timestamp, aid, SUM(read_or_not) AS reads_nb, SUM(comment_or_not) AS comments_nb, SUM(agree_or_not) AS agrees_nb, SUM(share_or_not) AS shares_nb FROM user_read GROUP BY aid;",
		}},
	}

	bls := &binlogdatapb.BinlogSource{
		Keyspace: "users",
		Shard:    shard,
		Filter:   filter,
		OnDdl:    binlogdatapb.OnDDLAction_IGNORE,
	}
	return bls, dbName
}

func replicationArticlesToBeRead() (*binlogdatapb.BinlogSource, string) {
	shard := os.Args[3]
	if shard == "" {
		panic(errors.New("missing source shard info"))
	}
	dbName := "users"

	filter := &binlogdatapb.Filter{
		Rules: []*binlogdatapb.Rule{
			{
				Match:  "be_read",
				Filter: "SELECT aid AS id, 0 AS timestamp, aid, SUM(read_or_not) AS reads_nb, CONCAT('ruid_', id) as read_uid_list, SUM(comment_or_not) AS comments_nb, CONCAT('cuid_', id) as comment_uid_list, SUM(agree_or_not) AS agrees_nb, CONCAT('auid_', id) as agree_uid_list, SUM(share_or_not) AS shares_nb, CONCAT('suid_', id) as share_uid_list FROM user_read GROUP BY aid;",
			},
		},
	}

	bls := &binlogdatapb.BinlogSource{
		Keyspace: "users",
		Shard:    shard,
		Filter:   filter,
		OnDdl:    binlogdatapb.OnDDLAction_IGNORE,
	}
	return bls, dbName
}

func replicationArticlesScience() (*binlogdatapb.BinlogSource, string) {
	dbName := "articles"

	filter := &binlogdatapb.Filter{
		Rules: []*binlogdatapb.Rule{{
			Match:  "article",
			Filter: "select * from article where category='science'",
		}},
	}

	bls := &binlogdatapb.BinlogSource{
		Keyspace: "articles",
		Shard:    "-80",
		Filter:   filter,
		OnDdl:    binlogdatapb.OnDDLAction_IGNORE,
	}
	return bls, dbName
}

func replicationPopularityScience() (*binlogdatapb.BinlogSource, string) {
	dbName := "articles"

	filter := &binlogdatapb.Filter{
		Rules: []*binlogdatapb.Rule{{
			Match:  "popularity",
			Filter: "select * from popularity",
		}},
	}

	bls := &binlogdatapb.BinlogSource{
		Keyspace: "articles",
		Shard:    "-80",
		Filter:   filter,
		OnDdl:    binlogdatapb.OnDDLAction_IGNORE,
	}
	return bls, dbName
}

func main() {
	shardInfoStr := os.Args[2]
	shardInfo := shardInfo{}
	err := json.Unmarshal([]byte(shardInfoStr), &shardInfo)
	if err != nil {
		log.Fatal(err)
	}
	tabletDestID := fmt.Sprintf("%s-%d", shardInfo.MasterAlias.Cell, shardInfo.MasterAlias.UID)

	bls := &binlogdatapb.BinlogSource{}
	dbName := ""
	if os.Args[1] == "articles_science" {
		bls, dbName = replicationArticlesScience()
	} else if os.Args[1] == "be_read_articles" {
		bls, dbName = replicationArticlesToBeRead()
	} else if os.Args[1] == "read_stats" {
		bls, dbName = replicationReadStats()
	} else if os.Args[1] == "popularity_science" {
		bls, dbName = replicationPopularityScience()
	}

	val := sqltypes.NewVarBinary(fmt.Sprintf("%v", bls))
	var sqlEscaped bytes.Buffer
	val.EncodeSQL(&sqlEscaped)
	query := fmt.Sprintf("insert into _vt.vreplication "+
		"(db_name, source, pos, max_tps, max_replication_lag, tablet_types, time_updated, transaction_timestamp, state) values"+
		"('%s', %s, '', 9999, 9999, 'master', 0, 0, 'Running')", dbName, sqlEscaped.String())

	fmt.Printf("VReplicationExec %s '%s'\n", tabletDestID, strings.Replace(query, "'", "'\"'\"'", -1))
}
