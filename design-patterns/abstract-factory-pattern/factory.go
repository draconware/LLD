package main

type IFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

type IShoe interface {
	getLogo() string
	setLogo(string)
	getSize() int
	setSize(int)
}

type IShirt interface {
	getLogo() string
	setLogo(string)
	getSize() int
	setSize(int)
}

func GetBrandFactory(factoryName string) IFactory {
	switch factoryName {
	case "addidas":
		return &AddidasFactory{}
	case "nike":
		return &NikeFactory{}
	default:
		return nil
	}
}
