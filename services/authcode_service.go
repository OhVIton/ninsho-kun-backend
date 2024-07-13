package services

import (
	"bytes"
	"log"
	"regexp"
	"time"

	"github.com/OhVIton/ninsho-kun-backend/repositories"
	"howett.net/plist"
)

func (s *AppService) FetchAuthCodeService(after time.Time) (string, error) {
	const mailApp = "com.microsoft.Outlook"
	re := regexp.MustCompile(`【YNU情報基盤センター】認証コード :([0-9]*)`)

	recArray, err := repositories.SelectRecords(s.db)
	if err != nil {
		return "", err
	}

	log.Println("Decoding record.data")
	var authCode string
	for _, rec := range recArray {

		buf := bytes.NewReader(rec.Data)
		decoder := plist.NewDecoder(buf)

		unixtime := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC).Unix() + int64(rec.DeliveredDate.Float64)
		if unixtime <= after.Unix() {
			continue
		}

		var data interface{}
		err := decoder.Decode(&data)
		if err != nil {
			return "", err
		}

		dataMap := data.(map[string]interface{})
		if dataMap["app"] != mailApp {
			continue
		}

		reqMap := dataMap["req"].(map[string]interface{})

		body := reqMap["body"].(string)
		matches := re.FindStringSubmatch(body)
		if len(matches) == 2 {
			authCode = matches[1]
			break
		}
	}

	log.Println("Returning auth code")
	if authCode == "" {
		return "", nil
	}
	return authCode, nil
}
