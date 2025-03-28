package main

func main() {
	patient := &Patient{name: "abc", registrationDone: true}

	cashier := &Cashier{}

	medical := &Medical{}
	medical.setNext(cashier)

	doctor := &Doctor{}
	doctor.setNext(medical)

	reception := &Reception{}
	reception.setNext(doctor)

	reception.execute(patient)
}
