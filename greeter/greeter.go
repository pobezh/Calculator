package greeter

import (
	"fmt"
	"strings"
	"unicode"
)

type Greeter struct {
	Name string
}

func New() *Greeter {
	return &Greeter{Name: "Guest"}
}

func (g *Greeter) Greet() string {
	return fmt.Sprintf("Hello, %s! Welcome to the multi-file Go project from CodeInterview.", g.Name)
}

func (g *Greeter) UpdateName(newName string) {
	if newName != "" && strings.TrimSpace(newName) != "" {
		g.Name = newName
	}
}

func (g *Greeter) IsNameValid(name string) bool {
	return name != "" && strings.IndexFunc(name, func(r rune) bool {
		return !unicode.IsLetter(r)
	}) == -1
}
