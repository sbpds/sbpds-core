// Copyright (c) 2022 Bryan Joshua Pedini
// License: GPL-3-or-later · see the LICENSE file for more information
package main

import (
	"fmt"
	"os"

	"github.com/miekg/dns"
)

var records = map[string]string{ //DEBUG · TODO: remove
	"example.com.": "192.168.1.1",
	"example.net.": "10.0.0.1",
}

func parse(m *dns.Msg) {
	for _, q := range m.Question {
		logger.Debug("Query type", q.Qtype, "for", q.Name)
		switch q.Qtype {
		case dns.TypeA:
			ipAddr := records[q.Name]
			if ipAddr != "" {
				rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, ipAddr))
				if err == nil {
					m.Answer = append(m.Answer, rr)
					m.Authoritative = true
				}
			}
		}
	}
}

func handle(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		parse(m)
	}
	w.WriteMsg(m)
}

func startServer(net string) {
	server := &dns.Server{
		Addr: SERVER_OPTIONS["BIND_ADDRESS"] + ":" + SERVER_OPTIONS["BIND_PORT"],
		Net:  net,
	}
	logger.Debug("Starting in", net, "at", SERVER_OPTIONS["BIND_ADDRESS"]+":"+SERVER_OPTIONS["BIND_PORT"])
	if err := server.ListenAndServe(); err != nil {
		logger.Fatal("Failed to start server:", err.Error())
		os.Exit(1)
	}
}
