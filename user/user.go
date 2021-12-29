package user

import "time"

func GetToday() string {
	t := time.Now()
	return t.Format("2006-01-02")
}