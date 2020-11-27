package db

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	"vitess.io/vitess/go/vt/vitessdriver"
)

// DatabaseManager contains the connection pool to Vitess
// MySQL cluster.
// DatabaseManager's methods enable using the models with
// the rest of the cluster properly.
type DatabaseManager struct {
	db *sql.DB
}

// NewDatabaseManager will create a new connection pool
// to Vitess MySQL cluster using the `users` keyspace.
//
// TODO: add some configuration in the call of this function
func NewDatabaseManager() (dbm *DatabaseManager, err error) {
	dbm = &DatabaseManager{}
	dbm.db, err = vitessdriver.Open("traefik:9112", "users@master")
	if err != nil {
		return nil, err
	}
	dbm.db.SetConnMaxLifetime(time.Second * 5)
	return dbm, err
}

// GetUserByID will returns a single User corresponding
// to the given ID.
func (dbm *DatabaseManager) GetUserByID(ID uint64) (user User, err error) {
	qc := `WHERE _id LIKE ?`
	user, err = dbm.fetchUser(qc, strconv.Itoa(int(ID)))
	if err != nil {
		log.Println(err.Error())
		return user, err
	}
	return user, nil
}

// GetUsersOfRegion fetches all the users matching the given
// region filter. It returns an array of Users, and an error
// if there is any. No other filter than region is being used
// thus, the function might be highly computational intensive
// depending on the data stored in Vitess.
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

// InsertUser inserts the given User into Vitess MySQL cluster.
// The new auto-generated ID will be returned, in addition to an
// error if any.
func (dbm *DatabaseManager) InsertUser(user User) (newID int, err error) {
	sql := `INSERT INTO user (timestamp,id,uid,name,gender,email,phone,dept,grade,language,region,role,preferTags,obtainedCredits) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	id, err := dbm.insertUser(sql, user)
	if err != nil {
		return 0, err
	}
	newID = int(id)
	return newID, nil
}

func (dbm *DatabaseManager) insertUser(sql string, user User) (newID int64, err error) {
	query, err := dbm.db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer query.Close()

	res, err := query.Exec(user.Timestamp, user.ID2, user.UID, user.Name, user.Gender,
		user.Email, user.Phone, user.Dept, user.Grade, user.Language, user.Region,
		user.Role, user.PreferTags, user.ObtainedCredits)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
