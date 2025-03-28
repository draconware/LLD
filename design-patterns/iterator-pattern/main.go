package main

import "fmt"

func main() {
	user1 := &User{
		name: "mayank",
		age:  26,
	}

	user2 := &User{
		name: "palak",
		age:  24,
	}

	user3 := &User{
		name: "mom",
		age:  56,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2, user3},
	}

	userIterator := userCollection.createIterator()

	for userIterator.hasNext() {
		user := userIterator.getNext()
		fmt.Printf("name: %s\tage: %d\n", user.name, user.age)
	}
}
