package main

import (
	"fmt"
	"github.com/grandcat/zeroconf"
	"os"
	"os/signal"
)

var mdns *zeroconf.Server

func main() {
	fmt.Println("ok")
	PublishName("ream1", "test")

	// Wait for SIGNAL (CTRL-c), then close servers and exit.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
}

func PublishName(realm string, serviceName string) *zeroconf.Server {
	hostname, _ := os.Hostname()
	if mdns == nil {
		mdns, _ = zeroconf.Register(hostname, fmt.Sprintf("_%s._tcp", serviceName), "local.",
			5020, []string{fmt.Sprintf("realm=%s", realm)}, nil)
		return mdns
	}
	return nil
}
