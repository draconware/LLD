package main

import "fmt"

type Department interface {
	execute(*Patient)
	setNext(Department)
}

type Patient struct {
	name             string
	registrationDone bool
	checkUpDone      bool
	medicalDone      bool
	paymentDone      bool
}

type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(d Department) {
	r.next = d
}

type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.checkUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.checkUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(de Department) {
	d.next = de
}

type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicalDone {
		fmt.Println("Medical examination already done")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical examination of patient started")
	p.medicalDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(d Department) {
	m.next = d
}

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment already done")
	} else {
		fmt.Println("Cashier getting payment for patient")
	}
}

func (c *Cashier) setNext(d Department) {
	c.next = d
}
