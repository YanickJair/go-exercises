package pattern

import (
	"log"
	"os"
	"sync"
)

// HydraLogger - Embedd the log.Looger interface
type HydraLogger struct {
	*log.Logger
	filename string
}

var hlogger *HydraLogger
var once sync.Once

// GetInstance - call our struct only and only once
func GetInstance() *HydraLogger {
	once.Do(func() {
		hlogger = createLogger("hydralogger.log")
	})
	return hlogger
}

func createLogger(fname string) *HydraLogger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &HydraLogger{
		filename: fname,
		Logger:   log.New(file, "Hydra ", log.Lshortfile),
	}
}
