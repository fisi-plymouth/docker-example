package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println("Go app is running..")
		time.Sleep(time.Second * 5)
	}
}
