package main

import (
	"context"
	"github.com/grandcat/zeroconf"
	"log"
	"time"
)

func main() {
	for {
		// Discover all services on the network (e.g. _workstation._tcp)
		resolver, err := zeroconf.NewResolver(nil)
		if err != nil {
			log.Fatalln("Failed to initialize resolver:", err.Error())
		}

		entries := make(chan *zeroconf.ServiceEntry)
		go func(results <-chan *zeroconf.ServiceEntry) {
			for entry := range results {
				log.Println(entry)
			}
			log.Println("No more entries.")
		}(entries)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()
		err = resolver.Browse(ctx, "a_workstation._tcp", "local.", entries)
		if err != nil {
			log.Fatalln("Failed to browse:", err.Error())
		}

		<-ctx.Done()
	}
}
