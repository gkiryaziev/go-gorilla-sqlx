package util

import "time"

// "2015-05-06" to RFC3339
func DateToRFC3339(date string) (time.Time, error) {
	result, err := time.Parse(time.RFC3339, date+"T00:00:00.00Z")
	if err != nil {
		return time.Time{}, err
	}
	return result, nil
}
