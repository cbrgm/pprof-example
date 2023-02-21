package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"
)

type NPC interface {
	Name() string
	Live()
	Eat()
	Drink()
	Walk()
	Talk()
	Work()
}

var Entities = []NPC{
	&Peasant{},
	&Bartender{},
	&Knight{},
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	for {
		for _, e := range Entities {
			e.Live()
		}
		time.Sleep(time.Second)
	}
}

// helpers
const (
	Ki = 1024
	Mi = Ki * Ki
	Gi = Ki * Mi
	Ti = Ki * Gi
	Pi = Ki * Ti
)
