package main

import (
	"fmt"

	"github.com/mastik5h/LLD/splitwise/client"
)

func main() {
	fmt.Println("Welcome to Spliwise..")
	client.Initialize()
	user1 := client.CreateUserClient("Mayank", "may@gmail.com", "897845873985")
	user2 := client.CreateUserClient("Ayush", "aay@gmail.com", "3434783934594")
	user3 := client.CreateUserClient("Sudhanshi", "sud@gmail.com", "786983457938")

	group1 := client.CreateGroupClient("Bangalore house", []string{user1, user2, user3})

	client.CreateGroupExpenseClient(user1, group1, "Maid", "", "", 9000.0, map[string]float64{
		user1: 9000.0,
		user2: 0.0,
		user3: 0.0,
	}, map[string]float64{
		user1: 3000.0,
		user2: 3000.0,
		user3: 3000.0,
	})

	client.CreateGroupExpenseClient(user2, group1, "Friday Trip", "", "", 4500.0, map[string]float64{
		user1: 0.0,
		user2: 0.0,
		user3: 4500.0,
	}, map[string]float64{
		user1: 1500.0,
		user2: 1500.0,
		user3: 1500.0,
	})

	client.CreateGroupExpenseClient(user2, group1, "Score", "", "", 5000.0, map[string]float64{
		user1: 0.0,
		user2: 5000.0,
		user3: 0.0,
	}, map[string]float64{
		user1: 1666.67,
		user2: 1666.67,
		user3: 1666.67,
	})

	client.CreateGroupExpenseClient(user2, group1, "Score2", "", "", 10000.0, map[string]float64{
		user1: 0.0,
		user2: 5000.0,
		user3: 5000.0,
	}, map[string]float64{
		user1: 3333.33,
		user2: 3333.33,
		user3: 3333.33,
	})
	se := client.SettleGroupExpenseClient(user3, group1)
	fmt.Println(se)
}
