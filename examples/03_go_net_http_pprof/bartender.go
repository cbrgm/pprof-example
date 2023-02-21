package main

import (
	"log"
	"time"
)

type Bartender struct{}

func (p *Bartender) Name() string {
	return "bartender"
}

func (p *Bartender) Live() {
	p.Eat()
	p.Drink()
	p.Walk()
	p.Talk()
	p.Work()
}

func (p *Bartender) Eat() {
	log.Println(p.Name(), "eat")
}

func (p *Bartender) Drink() {
	log.Println(p.Name(), "drink")
}

func (p *Bartender) Walk() {
	log.Println(p.Name(), "walk")
}

func (p *Bartender) Talk() {
	log.Println(p.Name(), "talk")
	<-time.After(time.Second)
}

func (p *Bartender) Work() {
	log.Println(p.Name(), "work")
}
