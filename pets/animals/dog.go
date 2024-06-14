package animals



type Dog struct {
	Name string
}

func (d *Dog) Eat(amount uint8) (uint8, error) {
	if amount > 15 {
		return 0, newError("dog cannot eat more than 15", nil)
	}

	return amount, nil
}
