package main

import "time"

var c = time.NewTicker(time.Minute * 1).C
var bucketSize = 60

// Ready is used to limit the max process size within one minute,
// the max size is 60 per minute.
func Ready() bool {
	select {
	case <-c:
		bucketSize = 60
	default:
	}
	if bucketSize <= 0 {
		return false
	}
	bucketSize--
	return true
}
