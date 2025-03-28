package main

type Collection interface {
	createIterator() Iterator
}

type UserCollection struct {
	users []*User
}

func (uc *UserCollection) createIterator() Iterator {
	return &UserIterator{
		index: 0,
		users: uc.users,
	}
}

type Iterator interface {
	hasNext() bool
	getNext() *User
}

type UserIterator struct {
	index int
	users []*User
}

func (ui *UserIterator) hasNext() bool {
	return ui.index < len(ui.users)
}

func (ui *UserIterator) getNext() *User {
	user := ui.users[ui.index]
	ui.index++
	return user
}

type User struct {
	name string
	age  int
}
