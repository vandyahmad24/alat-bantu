package date

import "time"

// GetEnv : Simple helper function to generate future time
func GenerateFutureTimeSeconds(seconds int) time.Time {
	now := time.Now()
	future := now.Add(time.Duration(seconds) * time.Second)
	return future
}

// Change string YYYY-MM-DD to time.Time
func ChangeStringToTime(dateString string) (time.Time, error) {
	layout := "2006-01-02"

	date, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

// change time.Time to 2006-01-02
func ChangeTimeToString(timeReq time.Time) string {
	if timeReq.IsZero() {
		return "" // Mengembalikan string kosong jika waktu tidak valid (0001-01-01 00:00:00 +0000)
	}
	return timeReq.Format("2006-01-02")
}

// change time.Time to 2006-01-02 15:04:05
func ChangeTimeToStringWithSeconds(timeReq time.Time) string {
	if timeReq.IsZero() {
		return "" // Mengembalikan string kosong jika waktu tidak valid (0001-01-01 00:00:00 +0000)
	}
	return timeReq.Format("2006-01-02 15:04:05")
}
