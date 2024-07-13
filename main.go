package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/OhVIton/ninsho-kun-backend/api"
)

var (
	tmpDir = os.Getenv("TMPDIR")
)

func main() {
	log.Println("Starting server...")

	dbPath := tmpDir + "../0/com.apple.notificationcenter/db2/db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	r := api.NewRouter(db)

	r.Run(":6350")
}
