package config

import (
	"os"
)

var (
	// HOSTNAME current hostname
	HOSTNAME, _ = os.Hostname()
	// GMAPS API Key
	GMAPS_API = os.Getenv("GMAPS_API")
)
