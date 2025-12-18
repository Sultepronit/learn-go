package main

import (
	"fmt"
	"time"
)

func main() {
	pl := fmt.Println

	now := time.Now()
	pl(now) // 2025-12-16 14:36:36.60995144 +0000 UTC m=+0.000040641

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	// then := time.Date(2009, 11, 17, 20, 34, 58, 651387237) not enough arguments in call to time.Date
	pl(then) // 2009-11-17 20:34:58.651387237 +0000 UTC
	pl(then.Year()) // 2009
	pl(then.Month()) // November
	pl(then.Day()) // 17
	pl(then.Hour()) // 20
	pl(then.Minute()) // 34
	pl(then.Second()) // 58
	pl(then.Nanosecond()) // 651387237
	pl(then.Location()) // UTC
	pl(then.Weekday()) // Tuesday

	pl(then.Before(now)) // true
	pl(then.After(now)) // false
	pl(then.Equal(now)) // false

	diff := now.Sub(then)
	pl(diff) // 140993h53m45.705392616s
	pl(then.Sub(now)) // -140993h56m7.348664965s

	pl(diff.Hours()) // 140993.9650966631
	pl(diff.Minutes()) // 8.459638276113147e+06
	pl(diff.Nanoseconds()) // 507578325616869904

	pl(then.Add(-diff)) // 1993-10-18 02:35:07.79180884 +0000 UTC
}