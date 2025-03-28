package main

func main() {
	shirtItem := newItem("Nike Shirt")

	subscriber1 := &User{
		id: "mayank@gmail.com",
	}
	subscriber2 := &User{
		id: "palak@gmail.com",
	}

	shirtItem.register(subscriber1)
	shirtItem.register(subscriber2)

	shirtItem.updateStockAvailability()
}
