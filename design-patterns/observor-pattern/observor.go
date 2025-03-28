package main

import "fmt"

type Observor interface {
	getId() string
	update()
}

type Subject interface {
	register(Observor)
	deRegister(Observor)
	notifyAll()
}

type Item1 struct {
	name      string
	inStock   bool
	observors []Observor
}

func newItem(name string) *Item1 {
	return &Item1{
		name: name,
	}
}

func (i1 *Item1) updateStockAvailability() {
	fmt.Printf("Item %s is now available.\n", i1.name)
	i1.inStock = true
	i1.notifyAll()
}
func (i1 *Item1) register(o Observor) {
	for _, ro := range i1.observors {
		if ro.getId() == o.getId() {
			return
		}
	}
	i1.observors = append(i1.observors, o)
}
func (i1 *Item1) deRegister(o Observor) {
	roLen := len(i1.observors)
	for idx, ro := range i1.observors {
		if ro.getId() == o.getId() {
			i1.observors[roLen-1], i1.observors[idx] = i1.observors[idx], i1.observors[roLen-1]
			i1.observors = i1.observors[:roLen-1]
		}
	}
}
func (i1 *Item1) notifyAll() {
	for _, ro := range i1.observors {
		ro.update()
	}
}

type User struct {
	id string
}

func (u *User) getId() string {
	return u.id
}
func (u *User) update() {
	fmt.Printf("Sending out email to user: %s about stock availablity.\n", u.getId())
}
