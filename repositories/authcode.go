package repositories

import (
	"database/sql"
	"log"

	"github.com/OhVIton/ninsho-kun-backend/models"
)

func SelectRecords(db *sql.DB) ([]models.Record, error) {
	const limit = 5
	const sqlStr = `
		SELECT * FROM record
		ORDER BY rec_id DESC
		limit ?
	`

	log.Println("SelectRecords")
	rows, err := db.Query(sqlStr, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	log.Println("Seeking records")
	recArray := []models.Record{}
	for rows.Next() {
		rec := models.Record{}
		err := rows.Scan(
			&rec.RecordID,
			&rec.AppID,
			&rec.Uuid,
			&rec.Data,
			&rec.RequestDate,
			&rec.RequestLastDate,
			&rec.DeliveredDate,
			&rec.Presented,
			&rec.Style,
			&rec.SnoozeFireDate,
		)
		if err != nil {
			return nil, err
		}
		recArray = append(recArray, rec)
	}

	log.Println("Returning records")

	return recArray, nil
}
