package services

import "time"

type AuthCodeServicer interface {
	FetchAuthCodeService(after time.Time) (string, error)
}
