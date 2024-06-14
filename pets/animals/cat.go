package animals

import (
	"fmt"
	
)

type Cat struct {
	Animal
	Age uint8
}

func (c *Cat) Eat(amount uint8) (uint8, error) {
	if amount > 5 {
		return 0, fmt.Errorf("cat cannot eat more than 5")
	}

	return amount, nil
}

func (c *Cat) Walk() string {
	return fmt.Sprintf("%s is walking", c.GetName())
}

