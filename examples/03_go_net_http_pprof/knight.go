package main

import (
	"log"
)

type Knight struct {
	buffer [][Mi]byte
}

func (p *Knight) Name() string {
	return "knight"
}

func (p *Knight) Live() {
	p.Eat()
	p.Drink()
	p.Walk()
	p.Talk()
	p.Work()
}

func (p *Knight) Eat() {
	log.Println(p.Name(), "eat")
}

func (p *Knight) Drink() {
	log.Println(p.Name(), "drink")
}

func (p *Knight) Walk() {
	log.Println(p.Name(), "walk")
	max := Gi
	for len(p.buffer)*Mi < max {
		p.buffer = append(p.buffer, [Mi]byte{})
	}
}

func (p *Knight) Talk() {
	log.Println(p.Name(), "talk")
}

func (p *Knight) Work() {
	log.Println(p.Name(), "work")
}
