package main

import (
	"fmt"
	"os"

	"github.com/camembertaulaitcrew/nameinfo"
)

func main() {
	person := nameinfo.FromFullName(os.Args[1])
	fmt.Printf("%s\n", os.Args[1])
	for _, generator := range nameinfo.Generators {
		fmt.Printf("%s: %s\n", generator.Name(), person.Generate(generator))
	}
}
