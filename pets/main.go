package main

import (
	"fmt"
	"os"
	"pets/animals"
)

func main() {
	myCat := animals.Cat{
		Animal: animals.Animal{Name: "garfield"},
		Age:    5}
	myDog := animals.Dog{
		Animal: animals.Animal{Name: "fido"},
		Age:    3,
		Weight: 30}

	var feedToCat uint8 = 5
	var feedToDog uint8 = 15

	catFeed, err := feed(&myCat, feedToCat)

	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred: %s\n", err)
	} else {
		fmt.Printf("I just fed %s %d times", myCat.Name, catFeed)
	}

	fmt.Print("\n\t ====== \n")

	dogFeed, err := feed(&myDog, feedToDog)

	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred: %s\n", err)
	} else {
		fmt.Printf("I just fed %s %d times", myDog.Name, dogFeed)
	}

	fmt.Print("\n-----------------------------\n")
	displayInfo(&myCat)
	displayInfo(&myDog)
	displayInfo(42)
}

func feed(animal animals.EaterWalker, amount uint8) (uint8, error) {
	switch animal.(type) {
	case *animals.Cat:
		fmt.Println("I am a cat")
	case *animals.Dog:
		fmt.Println("I am a dog")
	default:
		fmt.Println("I am an animal")
	}
	fmt.Print("First lets walk, ")
	fmt.Println(animal.Walk())
	fmt.Println("Now, lets feed our", animal.GetName())
	return animal.Eat(amount)
}

func displayInfo(i interface{}) {
	switch v := i.(type) {
	case animals.EaterWalker:
		fmt.Println("I am an eater and walker")
	case animals.Eater:
		fmt.Println("I am an eater")
	case animals.Walker:
		fmt.Println("I am a walker")
	case animals.Cat:
		fmt.Println("I am a cat and my age is", v.Age)
	case animals.Dog:
		fmt.Println("I am a dog and my weight is", v.Weight)
	case int :
		fmt.Println("I am an int")
	}
}
