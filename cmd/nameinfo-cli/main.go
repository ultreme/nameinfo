package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/camembertaulaitcrew/nameinfo"
)

var generatorName = flag.String("generator", "", "generator")

func main() {
	flag.Parse()
	person := nameinfo.FromFullName(strings.Join(flag.Args(), " "))
	if *generatorName != "" {
		for _, generator := range nameinfo.Generators {
			if generator.Name() == *generatorName {
				fmt.Printf("ton %s est %q\n", generator.Name(), person.Generate(generator))
			}
		}
	} else {
		fmt.Printf("%s\n", strings.Join(flag.Args(), " "))
		for _, generator := range nameinfo.Generators {
			fmt.Printf("ton %s est %q\n", generator.Name(), person.Generate(generator))
		}
	}
}
