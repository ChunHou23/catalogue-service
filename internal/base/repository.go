package base

import (
	"log"
	"path/filepath"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbInstance *gorm.DB
var once sync.Once

func InitDBInstance() *gorm.DB {
	once.Do(func() {
		dsn := "user=chunhou password=password dbname=catalogue_service_db host=localhost port=5432 sslmode=disable search_path=catalogue_service"
		var err error
		dbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "catalogue_service.",
				SingularTable: false,
			},
		})
		if err != nil {
			log.Fatalf("failed to connect to the database: %v", err)
		}

		runMigrationScript()
	})
	log.Println("connected to database")

	return dbInstance
}

func runMigrationScript() {
	absPath, err := filepath.Abs("internal/db/migrations")
	if err != nil {
		log.Fatalf("failed to get absolute path: %v", err)
	}
	m, err := migrate.New(
		"file://"+absPath,
		"postgres://chunhou:password@localhost:5432/catalogue_service_db?sslmode=disable&search_path=catalogue_service",
	)
	if err != nil {
		log.Fatalf("failed to initialize migration: %v", err)
	}

	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("No migration needed")
		} else {
			log.Fatalf("failed to apply migrations: %v", err)
		}
	} else {
		log.Println("Migrations applied successfully")
	}
}
