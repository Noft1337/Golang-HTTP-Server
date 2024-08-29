package main

import (
	"fmt"
	"xhttp/src/logging"
	"xhttp/src/utils"
)

var log *logging.Logger

func init() {
	log = logging.NewLogger(0, logging.LOG_NAME)
}

func main() {
	log.Debug("How r yoused today ?!?")
	utils.ValidateIP("127.0.0.1")

	fmt.Printf("Hellow World\n")
}
