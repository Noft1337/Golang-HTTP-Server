package utils

import (
	"xhttp/src/logging"
)

var log *logging.Logger

func init() {
	log = logging.NewLogger(0, logging.LOG_NAME)
}
