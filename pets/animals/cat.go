package animals

import (
	"fmt"
)

type Cat struct {
	Animal
	Age      uint8
	IsAsleep bool
}

func (c *Cat) Eat(amount uint8) (uint8, error) {
	if c.IsAsleep {
		return 0, &ActionError{Name: c.Name, Reason: "is sleeping"}
	}
	if amount > 5 {
		return 0, fmt.Errorf("cat cannot eat more than 5")
	}

	return amount, nil
}

func (c *Cat) Walk() string {
	return fmt.Sprintf("%s is walking", c.GetName())
}
