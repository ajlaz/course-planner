package postgres

import (
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func LoadFromCSV(dir string, db *gorm.DB) {
	// load csv into db
	f, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines, err := ReadCSV(f)
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}

func ReadCSV(f *os.File) ([][]string, error) {
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}
