package main

import (
	"bytes"
	"fmt"
	"strings"

	"vitess.io/vitess/go/sqltypes"
	binlogdatapb "vitess.io/vitess/go/vt/proto/binlogdata"
)

func main() {
	vtctl := "vtctlclient"
	tabletID := "zone1-0136547469"
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
	fmt.Println(bls)
	fmt.Println(val, "\n")
	var sqlEscaped bytes.Buffer
	val.EncodeSQL(&sqlEscaped)
	query := fmt.Sprintf("insert into _vt.vreplication "+
		"(db_name, source, pos, max_tps, max_replication_lag, tablet_types, time_updated, transaction_timestamp, state) values"+
		"('%s', %s, '', 9999, 9999, 'master', 0, 0, 'Running')", dbName, sqlEscaped.String())

	fmt.Printf("%s VReplicationExec %s '%s'\n", vtctl, tabletID, strings.Replace(query, "'", "'\"'\"'", -1))
}
