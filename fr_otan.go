package nameinfo

import "fmt"

type FrOtan struct{}

var otanAlphabet = []string{
	"ALPHA",
	"BRAVO",
	"CHARLY",
	"DELTA",
	"ECHO",
	"FOXTROT",
	"GOLF",
	"HOTEL",
	"INDIA",
	"JULIETT",
	"KILO",
	"LIMA",
	"MIKE",
	"NOVEMBER",
	"OSCAR",
	"PAPA",
	"QUEBEC",
	"ROMEO",
	"SIERRA",
	"TANGO",
	"UNIFORM",
	"VICTOR",
	"WHISKEY",
	"XRAY",
	"YANKEE",
	"ZULU",
}

func (_ FrOtan) Name() string { return "nom otan" }
func (_ FrOtan) Generate(firstname, lastname string) string {
	return fmt.Sprintf("%s %s",
		otanAlphabet[firstLetterIdx(firstname)],
		otanAlphabet[lastLetterIdx(firstname)],
	)
}

func init() {
	Generators = append(Generators, FrOtan{})
}
