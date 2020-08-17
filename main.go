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