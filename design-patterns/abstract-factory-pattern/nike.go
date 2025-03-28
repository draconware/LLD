package main

type NikeFactory struct{}

func (af *NikeFactory) makeShoe() IShoe {
	return &NikeShoe{
		logo: "nike",
		size: 14,
	}
}

func (af *NikeFactory) makeShirt() IShirt {
	return &NikeShirt{
		logo: "nike",
		size: 42,
	}
}

type NikeShoe struct {
	logo string
	size int
}

func (addidasshoe *NikeShoe) getLogo() string {
	return addidasshoe.logo
}

func (addidasshoe *NikeShoe) getSize() int {
	return addidasshoe.size
}

func (addidasshoe *NikeShoe) setLogo(logo string) {
	addidasshoe.logo = logo
}

func (addidasshoe *NikeShoe) setSize(size int) {
	addidasshoe.size = size
}

type NikeShirt struct {
	logo string
	size int
}

func (addidasshirt *NikeShirt) getLogo() string {
	return addidasshirt.logo
}

func (addidasshirt *NikeShirt) getSize() int {
	return addidasshirt.size
}

func (addidasshirt *NikeShirt) setLogo(logo string) {
	addidasshirt.logo = logo
}

func (addidasshirt *NikeShirt) setSize(size int) {
	addidasshirt.size = size
}
