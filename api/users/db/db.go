package db

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"vitess.io/vitess/go/vt/vitessdriver"
)

// DatabaseManager contains the connection pool to Vitess
// MySQL cluster.
// DatabaseManager's methods enable using the models with
// the rest of the cluster properly.
type DatabaseManager struct {
	db *sql.DB
}

func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var parentCtx opentracing.SpanContext
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan != nil {
		parentCtx = parentSpan.Context()
	}

	tracer := opentracing.GlobalTracer()

	span := tracer.StartSpan(
		method,
		opentracing.ChildOf(parentCtx),
		opentracing.Tag{Key: string(ext.Component), Value: "gRPC"},
		ext.SpanKindRPCClient,
	)
	defer span.Finish()

	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	} else {
		md = md.Copy()
	}

	mdw := MetaDataWriter{md}
	err := tracer.Inject(span.Context(), opentracing.TextMap, mdw)
	if err != nil {
		span.LogFields(otlog.String("inject-error", err.Error()))
	}

	newCtx := metadata.NewOutgoingContext(ctx, md)
	err = invoker(newCtx, method, req, reply, cc, opts...)
	if err != nil {
		span.LogFields(otlog.String("call-error", err.Error()))
	}
	return err
}

// NewDatabaseManager will create a new connection pool
// to Vitess MySQL cluster using the `users` keyspace.
//
// TODO: add some configuration in the call of this function
func NewDatabaseManager() (dbm *DatabaseManager, err error) {
	dbm = &DatabaseManager{}
	dbm.db, err = vitessdriver.OpenWithConfiguration(
		vitessdriver.Configuration{
			Address:         "vitess-zone1-vtgate-srv.vitess:15999",
			Target:          "users@master",
			GRPCDialOptions: []grpc.DialOption{grpc.WithUnaryInterceptor(UnaryClientInterceptor)},
		},
	)
	if err != nil {
		return nil, err
	}
	dbm.db.SetConnMaxLifetime(time.Second * 5)
	return dbm, err
}

// GetUserByID will returns a single User corresponding
// to the given ID.
func (dbm *DatabaseManager) GetUserByID(ctx context.Context, vtspanctx string, ID uint64) (user User, err error) {
	qc := `WHERE id=?`
	user, err = dbm.fetchUser(ctx, vtspanctx, qc, strconv.Itoa(int(ID)))
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
func (dbm *DatabaseManager) GetUsersOfRegion(ctx context.Context, vtspanctx, region string) (users []User, err error) {
	qc := `WHERE region=?`
	users, err = dbm.fetchUsers(ctx, vtspanctx, qc, region)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (dbm *DatabaseManager) fetchUser(ctx context.Context, vtspanctx, qc string, args ...interface{}) (user User, err error) {
	psql := vtspanctx + `SELECT * FROM user ` + qc
	err = dbm.db.QueryRowContext(ctx, psql, args...).Scan(
		&user.ID, &user.Timestamp, &user.UID, &user.Name, &user.Gender,
		&user.Email, &user.Phone, &user.Dept, &user.Grade, &user.Language,
		&user.Region, &user.Role, &user.PreferTags, &user.ObtainedCredits)
	return user, err
}

func (dbm *DatabaseManager) fetchUsers(ctx context.Context, vtspanctx, qc string, args ...interface{}) (users []User, err error) {
	psql := vtspanctx + `SELECT * FROM user ` + qc
	rows, err := dbm.db.QueryContext(ctx, psql, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Timestamp, &user.UID, &user.Name, &user.Gender,
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
	sql := `INSERT INTO user (timestamp,uid,name,gender,email,phone,dept,grade,language,region,role,preferTags,obtainedCredits) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)`

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

	res, err := query.Exec(user.Timestamp, user.UID, user.Name, user.Gender,
		user.Email, user.Phone, user.Dept, user.Grade, user.Language, user.Region,
		user.Role, user.PreferTags, user.ObtainedCredits)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
