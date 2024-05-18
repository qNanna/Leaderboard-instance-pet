package config

import (
	"fmt"

	"github.com/gofor-little/env"
)

func Init() {
	if err := env.Load(".env"); err != nil {
		panic(fmt.Errorf("failed to load environment variables: %v", err))
	}
}
