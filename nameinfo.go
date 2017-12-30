package nameinfo

import "strings"

var Generators = []Generator{}

type Generator interface {
	Name() string
	Generate(firstname, lastname string) string
}

type Person struct {
	firstname string
	lastname  string
}

func (p *Person) Generate(generator Generator) string {
	return generator.Generate(p.firstname, p.lastname)
}

func FromFullName(fullname string) Person {
	return Person{
		firstname: strings.Split(fullname, " ")[0],
		lastname:  strings.Split(fullname, " ")[1],
	}
}
