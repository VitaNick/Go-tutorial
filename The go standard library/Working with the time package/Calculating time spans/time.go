package main

import (
	"fmt"
	"time"
)

func main() {

	// Mon 2 Jan 15:04:05 MST 2006

	first := time.Now()
	fmt.Printf("It is currently %v\n", first.Format("15:04:05"))
	time.Sleep(time.Second)
	second := time.Now()
	fmt.Printf("It is now %v\n", second.Format("15:04:05"))

	today := time.Now()
	fmt.Printf("Time is currently %v\n", today.Format("Monday, Jan 2 2006"))
	startDate := time.Date(2018, 07, 04, 9, 00, 00, 0, time.UTC)
	elapsed := time.Since(startDate)

	fmt.Printf("%v has passed since %v\n", elapsed, startDate.Format("Monday, Jan 2 2006"))
	fmt.Printf("Hours: %.0f Minutes: %.0f Seconds: %.0f\n", elapsed.Hours(), elapsed.Minutes(), elapsed.Seconds())

	{
		future := today.AddDate(0, 6, 0)
		past := today.AddDate(0, -6, 0)

		fmt.Printf("In six months it will be %v\n", future.Format("Monday, January 2"))
		fmt.Printf("Six months ago it was %v\n", past.Format("Monday, January 2"))
	}

	{
		future := today.Add(6 * time.Hour)

		fmt.Printf("In six hours it will be %v\n", future.Format("15:04"))
	}

	bedtime := time.Date(2020, 12, 31, 23, 0, 0, 0, time.Local)
	fmt.Printf("There is %.0f hours until bed time\n", time.Until(bedtime).Hours())
}
