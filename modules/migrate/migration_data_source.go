package migrate

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func Read() (*Migrations, error) {
	migrations := New()

	// Open file
	file, err := os.Open(MIGRATION_FILEPATH)
	if err != nil {
		return migrations, err
	}

	defer file.Close()

	// Read file
	reader := csv.NewReader(file)

	// Read all records

	records, err := reader.ReadAll()
	if err != nil {
		return migrations, err
	}

	// Loop through lines & turn into object
	for _, record := range records[1:] {
		isMigrate, _ := strconv.ParseBool(record[2])
		id, _ := strconv.Atoi(record[0]) 

		migration := Migration{
			Id:        id,
			Name:      record[1],
			IsMigrate: isMigrate,
		}

		migrations.Add(migration)
	}
	
	return migrations, nil
}

func write(m []Migration) error {
	// Create migrations/migrations.csv file
	file, err := os.Create(MIGRATION_FILEPATH)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewWriter(file)
	defer reader.Flush()

	// Write header
	err = reader.Write([]string{
		"id",
		"filename",
		"isMigrate",
	})

	if err != nil {
		return err
	}

	// Write data
	for _, v := range m {
		err := reader.Write([]string{
			fmt.Sprintf("%05d", v.Id),
			v.Name,
			strconv.FormatBool(v.IsMigrate),
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func Init() error {
	// Create folder if not exists
	if _, err := os.Stat(MIGRATION_FOLDER); err != nil {
		err = os.Mkdir(MIGRATION_FOLDER, 0755)
		if err != nil {
			return err
		}
	}

	// Check if migrations/migrations.csv exists
	if _, err := os.Stat(MIGRATION_FILEPATH); err == nil {
		return errors.New("Migration file already exists")
	}

	// find .sql files and end with _up in migrations folder
	migrations := []Migration{}

	regex := regexp.MustCompile(`_up\.sql$`)
	err := filepath.Walk(MIGRATION_FOLDER, func(path string, info os.FileInfo, err error) error {
		fileName := filepath.Base(path)
		if regex.Match([]byte(fileName)) {
			migrationName := fileName[6 : len(fileName)-7]
			id, _ := strconv.Atoi(fileName[0:5])
			migration := Migration{
				Id:        id,
				Name:      migrationName,
				IsMigrate: false,
			}

			migrations = append(migrations, migration)

			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	if err := write(migrations); err != nil {
		return err
	}

	return nil
}

func createSQL(m Migration, table string) error {
	sqlUpName := fmt.Sprintf("%05d_%s_up.sql", m.Id, m.Name)
	
	// check file is not exists
	if _, err := os.Stat(filepath.Join(MIGRATION_FOLDER, sqlUpName)); err != nil {
		fileUp, err := os.Create(filepath.Join(MIGRATION_FOLDER, sqlUpName))
		if err != nil {
			return err
		}
		defer fileUp.Close()

		// Write SQL
		_, err = fileUp.WriteString(fmt.Sprintf("CREATE TABLE %s (\n\n)", table))

		if err != nil {
			return err
		}
	}

	sqlDownName := fmt.Sprintf("%05d_%s_down.sql", m.Id, m.Name)

	// check file is not exists
	if _, err := os.Stat(filepath.Join(MIGRATION_FOLDER, sqlDownName)); err != nil {
		fileDown, err := os.Create(filepath.Join(MIGRATION_FOLDER, sqlDownName))
		if err != nil {
			return err
		}
		defer fileDown.Close()

		// Write SQL
		_, err = fileDown.WriteString(fmt.Sprintf("DROP TABLE IF EXISTS %s", table))

		if err != nil {
			return err
		}
	}

	return nil
}

func IsMigrationFileIsExists() bool {
	// Check folder exist
	if _, err := os.Stat(MIGRATION_FOLDER); err != nil {
		return false
	}
	if _, err := os.Stat(MIGRATION_FILEPATH); err != nil {
		return false
	}

	return true
}