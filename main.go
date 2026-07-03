package main

import (
	"log"
	"time"

	"github.com/marunyann416/gojiho/audio"
)

func main() {
	err := audio.Init()
	if err != nil {
		log.Fatal(err)
	}

	for {
		audio.CheckAndPlay()

		// 次の分まで待つ
		now := time.Now()
		next := now.Truncate(time.Minute).Add(time.Minute)

		time.Sleep(time.Until(next))
	}
}
