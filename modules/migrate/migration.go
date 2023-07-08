package migrate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/GraphZC/go-fiber-cmd/configs"
)

type Migration struct {
	// ID is the unique identifier for the migration.
	Id int
	// Name is the name of the migration.
	Name string
	// IsMigrate is the status of the migration.
	IsMigrate bool
}

type Migrations struct {
	migrations []Migration
}

var MIGRATION_FOLDER = filepath.Join(configs.PROJECT_DIR, "migrations")
var MIGRATION_FILEPATH = filepath.Join(MIGRATION_FOLDER, "migrations.csv")

func New() *Migrations {
	return &Migrations{
		migrations: []Migration{},
	}
}

func (m *Migrations) Add(migration Migration) {
	m.migrations = append(m.migrations, migration)
}

func (m *Migrations) AddAndCreateSQL(migration Migration, table string) {
	m.migrations = append(m.migrations, migration)
	createSQL(migration, table)
}

func (m *Migrations) Get() []Migration {
	return m.migrations
}

func (m *Migrations) GetLastId() int {
	if len(m.migrations) == 0 {
		return 0
	}
	return m.migrations[len(m.migrations)-1].Id
}

func (m *Migrations) Up() (int, error) {
	db := configs.Db
	n := 0
	// Loop through migrations
	for i := range m.migrations {
		// Check if migration is already applied
		if m.migrations[i].IsMigrate {
			continue
		}

		// Read *up.sql file
		upFilepath := filepath.Join(MIGRATION_FOLDER, fmt.Sprintf("%05d_%s_up.sql", m.migrations[i].Id, m.migrations[i].Name))
		sqlStatement, err := os.ReadFile(upFilepath)

		if err != nil {
			return 0, err
		}

		// Execute *up.sql file
		db.MustExec(string(sqlStatement))

		// Update migrations.csv
		m.migrations[i].IsMigrate = true

		// Print migrate success
		fmt.Printf("[Migrate] Migrate up %s success\n", m.migrations[i].Name)
		n++
	}

	return n, nil
}

func (m *Migrations) Down() (int, error) {
	db := configs.Db

	n := 0
	// Loop through migrations
	for i := range m.migrations {
		// Read *up.sql file
		upFilepath := filepath.Join(MIGRATION_FOLDER, fmt.Sprintf("%05d_%s_down.sql", m.migrations[i].Id, m.migrations[i].Name))
		sqlStatement, err := os.ReadFile(upFilepath)

		if err != nil {
			return 0, err
		}

		// Execute *up.sql file
		db.MustExec(string(sqlStatement))

		// Update migrations.csv
		m.migrations[i].IsMigrate = false

		// Print migrate success
		fmt.Printf("[Migrate] Migrate down %s success\n", m.migrations[i].Name)
		n++
	}

	return n, nil
}

func (m *Migrations) Done() error {
	return write(m.migrations)
}