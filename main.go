package main

import (
	"fmt"
)

type Pilot struct {
	firstName string
	lastName  string
	callsign  string
	aircraft  string
}

type Squadron struct {
	name           string
	numberOfPilots int
	currentTopGun  string
}

type TopGunEntity interface {
	introduce() string
}

func introduceSomething(t TopGunEntity) {
	fmt.Println(t)
	fmt.Println(t.introduce())
}

func (p *Pilot) introduce() string {
	intro := fmt.Sprintf("Introducing %s %s, callsign %s, flying the %s", p.firstName, p.lastName, p.callsign, p.aircraft)
	return intro
}

func (s *Squadron) introduce() string {
	s.currentTopGun = "Maverick"
	intro := fmt.Sprintf("Introducing squadron %s, with %d pilots. Current top gun: %s", s.name, s.numberOfPilots, s.currentTopGun)
	return intro
}

func main() {
	p := Pilot{firstName: "Pete", lastName: "Mitchell", aircraft: "f14", callsign: "Maverick"}
	s := Squadron{name: "Top Gun", numberOfPilots: 10, currentTopGun: "Iceman"}
	entities := []TopGunEntity{&p, &s}
	for _, entity := range entities {
		introduceSomething(entity)
	}
}
