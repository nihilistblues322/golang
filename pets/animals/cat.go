package animals

import "fmt"

type Cat struct {
	Name string
}

func (c *Cat) Eat(amount uint8) (uint8, error) {
	if amount > 5 {
		return 0, fmt.Errorf("cat cannot eat more than 5")
	}

	return amount, nil
}
