package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
)

func InitDB() {
	migrateDb := flag.Bool("migratedb", false, "Initialize database's structure")
	seedDb := flag.Bool("seeddb", false, "Seeding database's data")
	flag.Parse()

	if *migrateDb {
		log.Printf("Initialize database's structure")
		mustMigrateDB()
	}

	if *seedDb {
		log.Printf("Seeding database's data")
		mustSeedDB()
	}
}

func mustMigrateDB() {
	config := GetConfig()
	mustExecMultilineSQLScript(filepath.Join(config.DB.Scripts, "structure.sql"))
}

func mustSeedDB() {
	config := GetConfig()
	mustExecMultilineSQLScript(filepath.Join(config.DB.Scripts, "seeder.sql"))
}

func mustExecMultilineSQLScript(path string) {
	config := GetConfig()
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?multiStatements=true", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Database))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx := db.MustBegin()
	buf, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	tx.MustExec(string(buf))
	tx.Commit()
}
