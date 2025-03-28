package main

import "fmt"

func main() {
	nikeFactory := GetBrandFactory("nike")
	addidasFactory := GetBrandFactory("addidas")

	nikeShoes := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	addidasShoes := addidasFactory.makeShoe()
	addidasShirt := addidasFactory.makeShirt()

	fmt.Println("Printing nike factory products....")
	printShoes(nikeShoes)
	printShirt(nikeShirt)

	fmt.Println("\nPrinting addidas factory products....")
	printShoes(addidasShoes)
	printShirt(addidasShirt)
}

func printShoes(shoe IShoe) {
	fmt.Println("Shoe size: ", shoe.getSize())
	fmt.Println("Shoe logo: ", shoe.getLogo())
}

func printShirt(shirt IShirt) {
	fmt.Println("Shirt size: ", shirt.getSize())
	fmt.Println("Shirt logo: ", shirt.getLogo())
}
