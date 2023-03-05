package main

import (
	"log"
	"os"

	"github.com/hyperace9701/magicedensniper/internal/models"
	"github.com/hyperace9701/magicedensniper/internal/sniper"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	checkError(err)

	var events = make(chan *models.Token, 5)

	// create sniper instance
	s, err := sniper.New(&sniper.Options{
		Endpoint:   os.Getenv("NODE_ENDPOINT"),
		Events:    events,
		PrivateKey: os.Getenv("PRIVATE_KEY"),
	})
	checkError(err)

	// run sniper concurrently
	go func() {
		err = s.Start()
		checkError(err)
	}()

	for action := range events {
		action := action

		go func() {
			log.Printf("%#v\n", action)
		}()
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
