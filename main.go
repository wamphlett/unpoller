package main

//nolint:gci
import (
	"log"
	"os"
	"time"

	"github.com/unpoller/poller"
	// Load input plugins!
	_ "github.com/unpoller/inputunifi"
	// Load output plugins!
	_ "github.com/wamphlett/influxunifi"
)

// Keep it simple.
func main() {
	// Set time zone based on TZ env variable.
	setTimeZone(os.Getenv("TZ"))

	if err := poller.New().Start(); err != nil {
		log.Fatalln("[ERROR]", err)
	}
}

func setTimeZone(timezone string) {
	if timezone == "" {
		return
	}

	var err error

	if time.Local, err = time.LoadLocation(timezone); err != nil {
		log.Printf("[ERROR] Loading TZ Location '%s': %v\n", timezone, err)
	}
}
