package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	Year := t.Year()
	Month := t.Month()
	Day := t.Day()

	fmt.Printf("Today is %d/%d/%d\n", Day, Month, Year)
	fmt.Println(t.Format(time.Kitchen))
	fmt.Println(t.Format("Mon 2 Jan 15:04:05 MST 2006"))
	fmt.Println(t.Format("Monday, 2 January in the year 2006"))

}
