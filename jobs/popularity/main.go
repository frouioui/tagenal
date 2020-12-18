package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"vitess.io/vitess/go/vt/vitessdriver"
)

const (
	defaultGranularity  = "daily"
	defaultLimitArticle = 5
)

type UserReadCount struct {
	AID   int
	Reads int
}

type Popularity struct {
	ID          int
	Timestamp   *time.Time
	MainAID     int
	Temporality string
	ListAID     string
}

type DatabaseManager struct {
	dbArticles *sql.DB
	dbUsers    *sql.DB
}

type Params struct {
	Granularity string
	Limit       int
	VtGate      string
}

func getNbDaysFromGranularity(granularity string) int {
	if granularity == "daily" {
		return 1
	} else if granularity == "weekly" {
		return 7
	} else if granularity == "monthly" {
		return 30
	}
	return 0
}

func (dbm *DatabaseManager) insetPopularity(popularity Popularity) error {
	sql := `INSERT INTO popularity (timestamp,temporality,main_aid,aid_list) VALUES (?,?,?,?)`
	query, err := dbm.dbArticles.Prepare(sql)
	if err != nil {
		return err
	}
	defer query.Close()

	_, err = query.Exec(time.Now().Unix(), popularity.Temporality, popularity.MainAID, popularity.ListAID)
	if err != nil {
		return err
	}
	return nil
}

func (dbm *DatabaseManager) getCategoryFromArticle(aid int) (category string, err error) {
	err = dbm.dbArticles.QueryRow(`SELECT category FROM article WHERE id=?`, aid).Scan(&category)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return category, nil
}

func (dbm *DatabaseManager) getUserRead(limit int, granularity string) (urds []UserReadCount, err error) {
	sql := fmt.Sprintf(`SELECT aid, SUM(read_or_not) AS reads FROM user_read WHERE FROM_UNIXTIME(timestamp) >= (NOW() - INTERVAL %d DAY) AND (FROM_UNIXTIME(timestamp) <= NOW()) GROUP BY aid ORDER BY reads DESC LIMIT %d`, getNbDaysFromGranularity(granularity), limit)
	rows, err := dbm.dbUsers.Query(sql)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		var urd UserReadCount
		err = rows.Scan(&urd.AID, &urd.Reads)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		urds = append(urds, urd)
	}
	return urds, nil
}

func createVitessDriverConfig(vtgate, target string) vitessdriver.Configuration {
	return vitessdriver.Configuration{
		Address: vtgate,
		Target:  target,
	}
}

func createDatabaseManager(vtgate string) (DatabaseManager, error) {
	var err error
	dbm := DatabaseManager{}
	dbm.dbArticles, err = vitessdriver.OpenWithConfiguration(createVitessDriverConfig(vtgate, "articles@master"))
	if err != nil {
		return DatabaseManager{}, err
	}
	dbm.dbUsers, err = vitessdriver.OpenWithConfiguration(createVitessDriverConfig(vtgate, "users@master"))
	if err != nil {
		return DatabaseManager{}, err
	}
	return dbm, nil
}

func getParams() Params {
	p := Params{}
	flag.IntVar(&p.Limit, "limit", defaultLimitArticle, "limit <int> to limit fetched articles")
	flag.StringVar(&p.Granularity, "gran", defaultGranularity, "gran <string{daily|weekly|monthly}> to specify the granularity")
	flag.StringVar(&p.VtGate, "vtgate", "", "vtgate <string> to specify vtgate address")
	flag.Parse()
	return p
}

func main() {
	p := getParams()

	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("job has started")

	dbm, err := createDatabaseManager(p.VtGate)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dbm.dbArticles.Close()
	defer dbm.dbUsers.Close()

	log.Println("connected to db")
	userReadsCount, err := dbm.getUserRead(p.Limit, p.Granularity)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("fetched users")

	newPopularity := Popularity{
		MainAID:     userReadsCount[0].AID,
		Temporality: p.Granularity,
	}
	for i, ur := range userReadsCount {
		cat, err := dbm.getCategoryFromArticle(ur.AID)
		if err != nil {
			log.Fatal(err.Error())
		}
		if cat == "technology" {
			newPopularity.MainAID = ur.AID
		}
		if i == 0 {
			newPopularity.ListAID = fmt.Sprintf("%d", ur.AID)
		} else {
			newPopularity.ListAID = fmt.Sprintf("%s,%d", newPopularity.ListAID, ur.AID)
		}
	}
	log.Println("details done")
	err = dbm.insetPopularity(newPopularity)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("job has finished")
}
