package main

import (
	"time"
	"log"
)

func main() {
	log.SetFlags(0)

	start, err := time.Parse("2006-01-02 15:04:05", "2017-09-09 12:14:00")
	if err != nil {
		log.Fatal(err)
	}

	end, err := time.Parse("2006-01-02 15:04:05", "2017-09-09 12:29:02")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", end.Sub(start).Seconds())
}

