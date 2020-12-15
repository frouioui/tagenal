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

func replicationArticlesToBeRead() (*binlogdatapb.BinlogSource, string) {
	shard := os.Args[3]
	if shard == "" {
		panic(errors.New("missing source shard info"))
	}
	dbName := "articles"

	filter := &binlogdatapb.Filter{
		Rules: []*binlogdatapb.Rule{{
			Match:  "be_read",
			Filter: "SELECT id as id, timestamp, id AS aid, 0 as reads_nb, CONCAT('ruid_', id) as read_uid_list, 0 as comments_nb, CONCAT('cuid_', id) as comment_uid_list, 0 as agrees_nb, CONCAT('auid_', id) as agree_uid_list, 0 as shares_nb, CONCAT('suid_', id) as share_uid_list  FROM article",
		}},
	}

	bls := &binlogdatapb.BinlogSource{
		Keyspace: "articles",
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

func main() {
	// vtctl := "vtctlclient -server=tagenal:8000"

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
	}

	val := sqltypes.NewVarBinary(fmt.Sprintf("%v", bls))
	var sqlEscaped bytes.Buffer
	val.EncodeSQL(&sqlEscaped)
	query := fmt.Sprintf("insert into _vt.vreplication "+
		"(db_name, source, pos, max_tps, max_replication_lag, tablet_types, time_updated, transaction_timestamp, state) values"+
		"('%s', %s, '', 9999, 9999, 'master', 0, 0, 'Running')", dbName, sqlEscaped.String())

	fmt.Printf("VReplicationExec %s '%s'\n", tabletDestID, strings.Replace(query, "'", "'\"'\"'", -1))
}
