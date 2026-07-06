package audio

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
	err := godotenv.Load("audio/.env")
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
	now := time.Now().Format("15:04")

	file, ok := schedule[now]
	if !ok {
		return
	}

	fmt.Println("Playing:", file)

	cmd := exec.Command(
		"mpg123",
		"-q",
		//		"-a", audioDevice,
		"sounds/"+file,
	)

	_ = cmd.Start()
}
