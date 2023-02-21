package main

import "log"

type Peasant struct {
	buffer []byte
}

func (p *Peasant) Name() string {
	return "peasant"
}

func (p *Peasant) Live() {
	p.Eat()
	p.Drink()
	p.Walk()
	p.Talk()
	p.Work()
}

func (p *Peasant) Eat() {
	log.Println(p.Name(), "eat")
	_ = make([]byte, 16*Mi)
}

func (p *Peasant) Drink() {
	log.Println(p.Name(), "drink")
}

func (p *Peasant) Walk() {
	log.Println(p.Name(), "walk")
}

func (p *Peasant) Talk() {
	log.Println(p.Name(), "talk")
}

func (p *Peasant) Work() {
	log.Println(p.Name(), "work")
	loop := 10000000000
	for i := 0; i < loop; i++ {
		// do nothing
	}
}
