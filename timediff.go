package main

import (
	"time"
	"log"
)

func main() {
	log.SetFlags(0)

	start, err := time.Parse("2006-01-02 15:04:05", "2017-09-03 13:25:00")
	if err != nil {
		log.Fatal(err)
	}

	end, err := time.Parse("2006-01-02 15:04:05", "2017-09-04 20:14:07")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", end.Sub(start).Seconds())
}

