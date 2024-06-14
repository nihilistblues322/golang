package main

import (
	"fmt"
	"os"
	"pets/animals"
)

func main() {
	myCat := animals.Cat{Name: "Garfield"}
	myDog := animals.Dog{Name: "Fido"}

	var feedToCat uint8 = 5
	var feedToDog uint8 = 15

	catFeed, err := feed(&myCat, feedToCat)

	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred: %s\n", err)
	} else {
		fmt.Printf("I just fed %s %d times", myCat.Name, catFeed)
	}

	fmt.Print("\n\n\t ====== \n\n\n")

	dogFeed, err := feed(&myDog, feedToDog)

	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred: %s\n", err)
	} else {
		fmt.Printf("I just fed %s %d times", myDog.Name, dogFeed)
	}
}

func feed(animal animals.Eater, amount uint8) (uint8, error) {
	return animal.Eat(amount)
}
