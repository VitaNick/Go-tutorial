package mylib

import (
	"fmt"
	"time"
)

type User struct {
	id        int
	name      string
	lastLogin *time.Time
}

func messageWriter(greeting, name string) string {
	message := fmt.Sprintf("%v, %v", greeting, name)

	return message
	message = ""
}
