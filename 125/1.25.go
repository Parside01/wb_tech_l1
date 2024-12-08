package main

import (
	"time"
)

func Sleep(duration time.Duration) {
	<-time.After(duration)
}
