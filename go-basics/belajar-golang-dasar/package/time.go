package main

import (
	"fmt"
	"time"
)

// # Package Time

func main() {
	fmt.Println("# Package Time: ")

	now := time.Now()

	fmt.Println("now: ", now)
	fmt.Println("year: ", now.Year())
	fmt.Println("month: ", now.Month())
	fmt.Println("day: ", now.Day())
	fmt.Println("hour: ", now.Hour())
	fmt.Println("minute: ", now.Minute())
	fmt.Println("second: ", now.Second())
	fmt.Println("nanosecond: ", now.Nanosecond())

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	
	fmt.Println("utc: ", utc)

	// or use this
	// timeFormat := time.RFC3339
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2020-10-01")
	fmt.Println("parse: ", parse)

}
 