package models

import "database/sql"

type Record struct {
	RecordID        int             `json:"rec_id"`
	AppID           int             `json:"app_id"`
	Uuid            []byte          `json:"uuid"`
	Data            []byte          `json:"data"`
	RequestDate     sql.NullFloat64 `json:"request_date"`
	RequestLastDate sql.NullFloat64 `json:"request_last_date"`
	DeliveredDate   sql.NullFloat64 `json:"delivered_date"`
	Presented       bool            `json:"presented"`
	Style           int             `json:"style"`
	SnoozeFireDate  sql.NullFloat64 `json:"snooze_fire_date"`
}
