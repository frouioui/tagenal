package db

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseManager struct {
	db *sql.DB
}

func NewDatabaseManager() (dbm *DatabaseManager, err error) {
	dbm = &DatabaseManager{}
	connStr := "user@tcp(traefik:3000)/users"

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	dbm.db = db
	dbm.db.SetConnMaxLifetime(time.Second * 1)
	return dbm, err
}

func (dbm *DatabaseManager) GetUserByID(ID uint64) (user User, err error) {
	qc := `WHERE _id=?`
	user, err = dbm.fetchUser(qc, strconv.Itoa(int(ID)))
	if err != nil {
		log.Println(err.Error())
		return user, err
	}
	return user, nil
}

func (dbm *DatabaseManager) GetUsersOfRegion(region string) (users []User, err error) {
	qc := `WHERE region=?`
	users, err = dbm.fetchUsers(qc, region)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (dbm *DatabaseManager) fetchUser(qc string, args ...interface{}) (user User, err error) {
	err = dbm.db.QueryRow(`SELECT * FROM user `+qc, args...).Scan(
		&user.ID, &user.Timestamp, &user.ID2, &user.UID, &user.Name, &user.Gender,
		&user.Email, &user.Phone, &user.Dept, &user.Grade, &user.Language,
		&user.Region, &user.Role, &user.PreferTags, &user.ObtainedCredits)
	return user, err
}

func (dbm *DatabaseManager) fetchUsers(qc string, args ...interface{}) (users []User, err error) {
	rows, err := dbm.db.Query(`SELECT * FROM user `+qc, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Timestamp, &user.ID2, &user.UID, &user.Name, &user.Gender,
			&user.Email, &user.Phone, &user.Dept, &user.Grade, &user.Language, &user.Region,
			&user.Role, &user.PreferTags, &user.ObtainedCredits)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
