package main

import (
	"fmt"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
	layoutEU  = "2 January, 2006"
)

func main() {
	// Mon 2 Jan 15:04:05 MST 2006

	startDate := time.Date(2018, 07, 04, 9, 00, 00, 00, time.UTC)

	fmt.Println(startDate.Format(layoutISO))
	fmt.Println(startDate.Format(layoutUS))
	fmt.Println(startDate.Format(layoutEU))

}
