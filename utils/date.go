package utils

import "time"

func GetTime() string {
	return time.Now().Format(time.RFC1123Z)
}
