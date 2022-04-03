// Copyright (c) 2022 Bryan Joshua Pedini
// License: GPL-3-or-later · see the LICENSE file for more information
package main

import (
	"os"

	gobasiclogger "git.bjphoster.com/b.pedini/go-basic-logger"
)

var (
	SERVER_OPTIONS = map[string]string{
		"BIND_ADDRESS":  "0.0.0.0",
		"BIND_PORT":     "53",
		"BIND_NET":      "both",
		"LOGLEVEL":      "INFO",
		"PROVIDER_FILE": "/config.yml",
	}
	logger gobasiclogger.Logger
)

func main() {
	for option := range SERVER_OPTIONS {
		if value, set := os.LookupEnv(option); set {
			SERVER_OPTIONS[option] = value
		}
	}
	logger = *new(gobasiclogger.Logger)
	logLevel := SERVER_OPTIONS["LOGLEVEL"]
	logger.Initialize(&logLevel)
}
