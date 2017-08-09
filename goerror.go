package main

import (
	"fmt"
	"time"

	"github.com/uiaworker/is105/goifc/uiaatoi"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
	i, err := uiaatoi.Uiaatoi("0s")
	if err != nil {
		fmt.Println(err)
	}
	// 0 returneres selv om det skrives en feilmelding
	//
	fmt.Println(i)
}
