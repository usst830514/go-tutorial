package main

import "fmt"

type iSportsFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	switch brand {
	case "adidas":
		return new(adidas), nil
	case "nike":
		return new(nike), nil
	default:
		return nil, fmt.Errorf("wrong brand type passed")
	}
}
