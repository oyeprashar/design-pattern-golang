package main

import "log"

// defining a shoe interface and implementing it for NikeBrand

type IShoes interface {
	setLogo(logo string)
	setPrice(price float64)
	getLogo() string
	getPrice() float64
}

// NikeShoes implements IShoes
type NikeShoes struct {
	logo  string
	price float64
}

func (nikeShoes *NikeShoes) setLogo(logo string) {
	nikeShoes.logo = logo

}

func (nikeShoes *NikeShoes) setPrice(price float64) {
	nikeShoes.price = price

}

func (nikeShoes *NikeShoes) getLogo() string {
	return nikeShoes.logo

}

func (nikeShoes *NikeShoes) getPrice() float64 {
	return nikeShoes.price

}

type IShoesBrand interface {
	getShoes() IShoes
}

// NikeBrand implements IShoesBrand
type NikeBrand struct {
	location string
}

func (nike *NikeBrand) getShoes() IShoes {
	return &NikeShoes{
		logo:  "nike_logo",
		price: 75000,
	}
}

func (nike *NikeBrand) setLocation(location string) {
	nike.location = location
}

func (nike *NikeBrand) getLocation() string {
	return nike.location
}

func ShoesFactory(brand string) IShoesBrand {

	switch brand {
	case "nike":
		return &NikeBrand{}

	default:
		panic("invalid brand")
	}

}

func main() {

	shoesBrand := ShoesFactory("nike")
	log.Println(shoesBrand.getShoes().getPrice())

}
