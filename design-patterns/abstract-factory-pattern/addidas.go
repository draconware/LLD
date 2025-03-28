package main

type AddidasFactory struct{}

func (af *AddidasFactory) makeShoe() IShoe {
	return &AddidasShoe{
		logo: "addidas",
		size: 14,
	}
}

func (af *AddidasFactory) makeShirt() IShirt {
	return &AddidasShirt{
		logo: "addidas",
		size: 42,
	}
}

type AddidasShoe struct {
	logo string
	size int
}

func (addidasshoe *AddidasShoe) getLogo() string {
	return addidasshoe.logo
}

func (addidasshoe *AddidasShoe) getSize() int {
	return addidasshoe.size
}

func (addidasshoe *AddidasShoe) setLogo(logo string) {
	addidasshoe.logo = logo
}

func (addidasshoe *AddidasShoe) setSize(size int) {
	addidasshoe.size = size
}

type AddidasShirt struct {
	logo string
	size int
}

func (addidasshirt *AddidasShirt) getLogo() string {
	return addidasshirt.logo
}

func (addidasshirt *AddidasShirt) getSize() int {
	return addidasshirt.size
}

func (addidasshirt *AddidasShirt) setLogo(logo string) {
	addidasshirt.logo = logo
}

func (addidasshirt *AddidasShirt) setSize(size int) {
	addidasshirt.size = size
}
