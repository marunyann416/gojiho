package audio2

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/joho/godotenv"
)

var audioDevice string
var schedule map[string]string

func Init() error {
	err := godotenv.Load("audio2/.env")
	if err != nil {
		return err
	}

	audioDevice = os.Getenv("AUDIO_DEVICE")

	schedule = make(map[string]string)

	for _, env := range os.Environ() {
		if len(env) < 9 {
			continue
		}

		if env[:9] == "SCHEDULE_" {
			eq := -1
			for i, c := range env {
				if c == '=' {
					eq = i
					break
				}
			}

			if eq == -1 {
				continue
			}

			key := env[:eq]
			value := env[eq+1:]

			t := key[9:]
			t = t[:2] + ":" + t[3:]

			schedule[t] = value
		}
	}

	return nil
}

func CheckAndPlay() {
	fmt.Println("[INFO] Loading .env")

	err := godotenv.Overload("audio2/.env")
	if err != nil {
		fmt.Println("[ERROR] Failed to load .env:", err)
		return
	}

	now := time.Now()

	fmt.Printf(
		"[INFO] Current time: %s (%s)\n",
		now.Format("2006-01-02 15:04:05"),
		now.Weekday(),
	)

	var key string

	// 平日判定
	if now.Weekday() != time.Saturday &&
		now.Weekday() != time.Sunday {

		key = fmt.Sprintf(
			"SCHEDULE_WEEKDAY_%02d_%02d",
			now.Hour(),
			now.Minute(),
		)

		fmt.Println("[INFO] Weekday detected.")
		fmt.Println("[INFO] Looking for:", key)

		if file := os.Getenv(key); file != "" {
			fmt.Println("[INFO] Schedule found:", file)
			fmt.Println("[INFO] Starting playback.")

			cmd := exec.Command(
				"mpg123",
				"-q",
				"-a", audioDevice,
				"sounds/"+file,
			)

			err := cmd.Start()
			if err != nil {
				fmt.Println("[ERROR] Failed to start mpg123:", err)
				return
			}

			fmt.Println("[INFO] Playback started successfully.")
			return
		}

		fmt.Println("[INFO] No weekday schedule found.")
	} else {
		fmt.Println("[INFO] Weekend detected. Skipping weekday schedule.")
	}

	// 毎日共通スケジュール確認
	key = fmt.Sprintf(
		"SCHEDULE_%02d_%02d",
		now.Hour(),
		now.Minute(),
	)

	fmt.Println("[INFO] Looking for:", key)

	file := os.Getenv(key)
	if file == "" {
		fmt.Println("[INFO] No schedule found for this time.")
		return
	}

	fmt.Println("[INFO] Schedule found:", file)
	fmt.Println("[INFO] Starting playback.")

	cmd := exec.Command(
		"mpg123",
		"-q",
		"-a", audioDevice,
		"sounds/"+file,
	)

	err = cmd.Start()
	if err != nil {
		fmt.Println("[ERROR] Failed to start mpg123:", err)
		return
	}

	fmt.Println("[INFO] Playback started successfully.")
}
