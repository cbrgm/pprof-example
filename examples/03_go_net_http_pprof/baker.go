package main

import (
	"log"
	"sync"
	"time"
)

type Baker struct{}

func (p *Baker) Name() string {
	return "baker"
}

func (p *Baker) Live() {
	p.Eat()
	p.Drink()
	p.Walk()
	p.Talk()
	p.Work()
}

func (p *Baker) Eat() {
	log.Println(p.Name(), "eat")
	m := &sync.Mutex{}
	m.Lock()
	go func() {
		time.Sleep(time.Second)
		m.Unlock()
	}()
	m.Lock()
}

func (p *Baker) Drink() {
	log.Println(p.Name(), "drink")
}

func (p *Baker) Walk() {
	log.Println(p.Name(), "walk")
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(30 * time.Second)
		}()
	}
}

func (p *Baker) Talk() {
	log.Println(p.Name(), "talk")
}

func (p *Baker) Work() {
	log.Println(p.Name(), "work")
}
