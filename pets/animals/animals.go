package animals

import "fmt"

type Eater interface {
	Eat(amount uint8) (uint8, error)
}

func newError(msg string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", msg, err)
	}

	return fmt.Errorf("%s", msg)
}
