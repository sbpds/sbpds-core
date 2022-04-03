// Copyright (c) 2022 Bryan Joshua Pedini
// License: GPL-3-or-later · see the LICENSE file for more information
package main

import (
	"os"
	"os/signal"
	"syscall"

	gobasiclogger "git.bjphoster.com/b.pedini/go-basic-logger"
	"github.com/miekg/dns"
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
	dns.HandleFunc(".", handle)
	switch SERVER_OPTIONS["BIND_NET"] {
	case "tcp":
		go startServer("tcp")
	case "udp":
		go startServer("udp")
	case "both":
		go startServer("tcp")
		go startServer("udp")
	default:
		logger.Fatal("Unable to start the server, network \"" + SERVER_OPTIONS["BIND_NET"] + "\" specified not valid")
		os.Exit(1)
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit
	logger.Info("Quitting because signal", s)
}
