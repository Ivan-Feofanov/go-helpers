package go_helpers

import (
	"io"
	"log"
	"os"
)

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func GetEnv(key string, def string) string {
	env := os.Getenv(key)
	if len(env) == 0 {
		return def
	}
	return env
}
