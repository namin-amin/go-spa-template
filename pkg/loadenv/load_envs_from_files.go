package loadenv

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadRequiredEnvFiles() {
	envs := []string{".env.dev", ".env"}
	var envsToLoad []string

	for _, v := range envs {
		if _, err := os.Stat(v); err == nil {
			envsToLoad = append(envsToLoad, v)
		} else if errors.Is(err, os.ErrNotExist) {
			continue
		} else {
			log.Fatalln(err)
		}
	}

	err := godotenv.Load(envsToLoad...)

	if err != nil {
		log.Fatalln(err)
	}
}
