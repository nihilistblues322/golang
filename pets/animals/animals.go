package animals

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Eater interface {
	Eat(amount uint8) (uint8, error)
}

type Walker interface {
	Walk() string
}

type EaterWalker interface {
	Eater
	Walker
	Named
}
type Named interface {
	GetName() string
}

type Animal struct {
	Name string
}

func (a *Animal) GetName() string {
	caser := cases.Title(language.English)
	return caser.String(a.Name)
}

func newError(msg string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", msg, err)
	}

	return fmt.Errorf("%s", msg)
}
