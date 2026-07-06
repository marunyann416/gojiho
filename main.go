package main

import (
	"log"
	"time"

	//"github.com/marunyann416/gojiho/audio"
	"github.com/marunyann416/gojiho/audio2"
)

func main() {
	err := audio2.Init()
	if err != nil {
		log.Fatal(err)
	}

	for {
		audio2.CheckAndPlay()

		// 次の分まで待つ
		now := time.Now()
		next := now.Truncate(time.Minute).Add(time.Minute)

		time.Sleep(time.Until(next))
	}
}
