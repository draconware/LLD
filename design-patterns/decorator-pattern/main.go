package main

import "fmt"

func main() {
	vegPizza := &VegPizza{}
	fmt.Println("Price for Veg Pizza: ", vegPizza.getPrice())

	fmt.Println("Adding tomato topping...")
	vegPizzaWithTomato := &TomatoTopping{
		pizza: vegPizza,
	}
	fmt.Println("Price for Veg Pizza with tomato topping: ", vegPizzaWithTomato.getPrice())

	fmt.Println("Adding Cheese topping...")
	vegPizzaWithCheese := &CheeseTopping{
		pizza: vegPizza,
	}
	fmt.Println("Price for Veg Pizza with cheese topping: ", vegPizzaWithCheese.getPrice())

	fmt.Println("Adding both tomato and cheeese topping...")
	vegPizzaWithCheeseAndTomato := &CheeseTopping{
		pizza: vegPizzaWithTomato,
	}
	fmt.Println("Price for Veg Pizza with cheese and tomato topping: ", vegPizzaWithCheeseAndTomato.getPrice())
}
