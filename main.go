package main

import (
	"github.com/isshoni-soft/roxxy"
)

var logger = roxxy.NewLogger("[media-manager]: ")

func main() {
	defer logger.Shutdown()

	logger.Log("Hello, World")
}
