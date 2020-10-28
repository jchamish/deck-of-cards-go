package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int

}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}


func main() {

	jim := person{
		firstName: "Jim",
		lastName: "Anderson",
		contact: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.updateName("new_name")
	jim.print()

	// Map is a key value pairs
	colors := map[string]string {
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	//var colors map[string]string
	//colors := make(map[string]string)
	//colors["red"] = "$ff0000"

	printMap(colors)

	//cards := newDeck()
	//cards.print()
	//
	//// deal func
	//hand, remainingCards := deal(cards, 5)
	//// print
	//hand.print()
	//remainingCards.print()
	//
	//// save cards to file
	//cards.saveToFile("my_cards.csv")
	//
	//// shuffle
	//remainingCards.shuffle()

	// struct

}