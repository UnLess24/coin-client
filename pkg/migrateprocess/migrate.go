package migrateprocess

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/UnLess24/coin/client/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func MustProcess(args []string, cfg *config.Config) {
	fmt.Println("Migrating the database...")

	mType, mCount := mustMigrationTypeAndCount(args)

	url := URLConnectToDB(cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Database, cfg.DB.SslMode)
	db := mustNewDB(cfg.DB.Name, url)
	defer func() { _ = db.Close() }()

	driver := mustDBDriver(db)

	m := mustMigrate(cfg.DB.Name, cfg.MigrationsPath, driver)
	defer func() { _, _ = m.Close() }()

	var err error
	if mCount == 0 {
		switch mType {
		case "up":
			err = m.Up()
		case "down":
			err = m.Down()
		}
	} else {
		err = m.Steps(mCount)
	}
	if err != nil {
		fmt.Println(fmt.Errorf("migration error %w", err))
	}

	fmt.Println("Migration successful")
}

func URLConnectToDB(user, pass, host, port, database, sslMode string) string {
	return fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=%v",
		user,
		pass,
		host,
		port,
		database,
		sslMode,
	)
}

func mustNewDB(dbName, url string) *sql.DB {
	db, err := sql.Open(dbName, url)
	if err != nil {
		panic(fmt.Errorf("db connect error %w", err))
	}
	return db
}

func mustDBDriver(db *sql.DB) database.Driver {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(fmt.Errorf("driver instance create error %w", err))
	}
	return driver
}

func mustMigrate(name, path string, driver database.Driver) *migrate.Migrate {
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%v", path), name, driver)
	if err != nil {
		panic(fmt.Errorf("migration source path error %w", err))
	}
	return m
}

func mustMigrationTypeAndCount(args []string) (string, int) {
	mType, mCount := "up", 0
	if len(args) > 0 {
		mType = args[0]
	}
	if mType != "up" && mType != "down" {
		panic(fmt.Errorf("migration type error"))
	}

	if len(args) > 1 {
		n, err := strconv.Atoi(args[1])
		if err != nil {
			panic(fmt.Errorf("migration count get error %w", err))
		}

		mCount = n
		if mType == "down" {
			mCount = -n
		}
	}
	return mType, mCount
}
