package main

import (
	"bytes"
	"encoding/json"
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

func main() {

	if len(os.Args) != 2 {
		log.Fatal("1 argument required, enter shard's info JSON")
	}

	shardInfoStr := os.Args[1]
	shardInfo := shardInfo{}
	err := json.Unmarshal([]byte(shardInfoStr), &shardInfo)
	if err != nil {
		log.Fatal(err)
	}

	tabletID := fmt.Sprintf("%s-%d", shardInfo.MasterAlias.Cell, shardInfo.MasterAlias.UID)
	vtctl := "vtctlclient"
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

	val := sqltypes.NewVarBinary(fmt.Sprintf("%v", bls))
	var sqlEscaped bytes.Buffer
	val.EncodeSQL(&sqlEscaped)
	query := fmt.Sprintf("insert into _vt.vreplication "+
		"(db_name, source, pos, max_tps, max_replication_lag, tablet_types, time_updated, transaction_timestamp, state) values"+
		"('%s', %s, '', 9999, 9999, 'master', 0, 0, 'Running')", dbName, sqlEscaped.String())

	fmt.Printf("%s VReplicationExec %s '%s'\n", vtctl, tabletID, strings.Replace(query, "'", "'\"'\"'", -1))
}
