package animals

import (
	"fmt"
)

type Dog struct {
	Animal
	Age uint8
	Weight uint8
}

func (d *Dog) Eat(amount uint8) (uint8, error) {
	if amount > 15 {
		return 0, newError("dog cannot eat more than 15", nil)
	}

	return amount, nil
}

func (d *Dog) Walk() string {
	return fmt.Sprintf("%s is walking", d.GetName())
}
