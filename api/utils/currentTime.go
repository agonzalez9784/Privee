package utils

import "time"

func GetCurrentTimeStamp() string {
	// Get the current time
	currentTime := time.Now()

	// Format the time as a string in the 'YYYY-MM-DD HH:MM:SS' format
	return currentTime.Format("2006-01-02 15:04:05")
}
