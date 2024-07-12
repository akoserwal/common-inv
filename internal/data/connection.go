package data

import (
	"common-inv/internal/conf"
	"database/sql"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	mocket "github.com/selvatico/go-mocket"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectionFactory struct {
	Config *conf.Data
	DB     *gorm.DB
}

var gormConfig *gorm.Config = &gorm.Config{
	PrepareStmt:       true,
	AllowGlobalUpdate: false, // change it to true to allow updates without the WHERE clause
	QueryFields:       true,
}

// NewConnectionFactory will initialize a singleton ConnectionFactory as needed and return the same instance.
// Go includes database connection pooling in the platform. Gorm uses the same and provides a method to
// clone a connection via New(), which is safe for use by concurrent Goroutines.
func NewConnectionFactory(config *conf.Data) (*ConnectionFactory, func()) {
	var db *gorm.DB
	var err error
	// refer to https://gorm.io/docs/gorm_config.html

	if config.Database.Dialect == "postgres" {
		add := ConnectionString(config)
		fmt.Println(add)
		db, err = gorm.Open(postgres.Open(add), gormConfig)
	} else {
		// TODO what other dialects do we support?
		panic(fmt.Sprintf("unsupported DB dialect: %s", config.Database.Dialect))
	}
	if err != nil {
		panic(fmt.Sprintf(
			"failed to connect to %s database %s with connection string: %s\nError: %s",
			config.Database.Dialect,
			config.Database.Name,
			LogSafeConnectionString(config),
			err.Error(),
		))
	}
	sqlDB, sqlDBErr := db.DB()
	if sqlDBErr != nil {
		panic(fmt.Errorf("unexpected connection error: %s", sqlDBErr))
	}

	sqlDB.SetMaxOpenConns(int(config.Database.MaxConnections))
	dbFactory := &ConnectionFactory{Config: config, DB: db}
	cleanup := func() {
		if err := dbFactory.close(); err != nil {
			log.Fatalf("Unable to close db connection: %s", err.Error())
		}
	}
	return dbFactory, cleanup
}

// NewMockConnectionFactory should only be used for defining mock database drivers
// This uses mocket under the hood, use the global mocket.Catcher to change how the database should respond to SQL
// queries
func NewMockConnectionFactory(dbConfig *conf.Data) *ConnectionFactory {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	sqlDB, err := sql.Open(mocket.DriverName, "connection_string")
	if err != nil {
		panic(err)
	}
	mocketDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}))
	if err != nil {
		panic(err)
	}
	connectionFactory := &ConnectionFactory{dbConfig, mocketDB}
	return connectionFactory
}

// New returns a new database connection
func (f *ConnectionFactory) New() *gorm.DB {
	if f.Config.Database.Debug {
		return f.DB.Debug()
	}
	return f.DB
}

// Checks to ensure a connection is present
func (f *ConnectionFactory) CheckConnection() error {
	return f.DB.Exec("SELECT 1").Error
}

// close will close the connection to the database.
// THIS MUST **NOT** BE CALLED UNTIL THE SERVER/PROCESS IS EXITING!!
// This should only ever be called once for the entire duration of the application and only at the end.
func (f *ConnectionFactory) close() error {
	sqlDB, sqlDBErr := f.DB.DB()
	if sqlDBErr != nil {
		return sqlDBErr
	}
	return sqlDB.Close()
}

// By default do no roll back transaction.
// only perform rollback if explicitly set by db.db.MarkForRollback(ctx, err)
const defaultRollbackPolicy = false

// TxFactory represents an sql transaction
type txFactory struct {
	resolved          bool
	rollbackFlag      bool
	tx                *sql.Tx
	txid              int64
	postCommitActions []func()
	db                *sql.DB
}

// newTransaction constructs a new Transaction object.
func (c *ConnectionFactory) newTransaction() (*txFactory, error) {
	sqlDB, sqlDBErr := c.DB.DB()
	if sqlDBErr != nil {
		return nil, sqlDBErr
	}
	f := &txFactory{
		db: sqlDB,
	}
	return f, f.begin()
}

func (f *txFactory) begin() error {
	tx, err := f.db.Begin()
	if err != nil {
		return err
	}

	var txid int64 = 0

	// current transaction ID set by postgres.  these are *not* distinct across time
	// and do get reset after postgres performs "vacuuming" to reclaim used IDs.
	row := tx.QueryRow("select txid_current()")
	if row != nil {
		err := row.Scan(&txid)
		if err != nil {
			return err
		}
	}

	f.tx = tx
	f.txid = txid
	f.resolved = false
	f.rollbackFlag = defaultRollbackPolicy
	return nil
}

// markedForRollback returns true if a transaction is flagged for rollback and false otherwise.
func (tx *txFactory) markedForRollback() bool {
	return tx.rollbackFlag
}

func ConnectionString(c *conf.Data) string {
	if c.Database.Sslmode != "disable" {
		return fmt.Sprintf(
			"host=%s port=%d user=%s password='%s' dbname=%s sslmode=%s sslrootcert=%s",
			c.Database.Host, c.Database.Port, c.Database.Username, c.Database.Password, c.Database.Name, c.Database.Sslmode, c.Database.DbCaCertFile,
		)
	}
	return fmt.Sprintf(
		"host=%s port=%d user=%s password='%s' dbname=%s sslmode=%s",
		c.Database.Host, c.Database.Port, c.Database.Username, c.Database.Password, c.Database.Name, c.Database.Sslmode,
	)
}

func LogSafeConnectionString(c *conf.Data) string {
	if c.Database.Sslmode != "disable" {
		return fmt.Sprintf(
			"host=%s port=%d user=%s password='<REDACTED>' dbname=%s sslmode=%s sslrootcert=<REDACTED>",
			c.Database.Host, c.Database.Port, c.Database.Username, c.Database.Name, c.Database.Sslmode,
		)
	}
	return fmt.Sprintf(
		"host=%s port=%d user=%s password='<REDACTED>' dbname=%s sslmode=%s",
		c.Database.Host, c.Database.Port, c.Database.Username, c.Database.Name, c.Database.Sslmode,
	)
}
