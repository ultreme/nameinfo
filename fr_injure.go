package nameinfo

import "fmt"

type FrInjure struct{}

func (_ FrInjure) Name() string { return "injure" }
func (_ FrInjure) Generate(firstname, lastname string) string {
	partA := []string{
		"sale",
		"espèce de/d'",
		"gros(se)",
		"petit(e)",
		"résidu de/d'",
		"grand(e)",
		"ramassi(e) de/d'",
	}
	partB := []string{
		"déjection",
		"cacochyme",
		"gredin",
		"gueudin",
		"fiante",
		"pigeon",
		"nouille",
		"bachibouzouk",
		"sagroin",
		"baboin",
		"emberlificoteur",
		"malpoli",
		"sale",
		"mioche",
		"résidu",
		"ramassi",
		"gros",
		"petit",
		"naze",
		"bide",
		"chacal",
		"chiotte",
	}
	partC := []string{
		"folichon",
		"empauté",
		"sans poil",
		"sans organe",
		"sans les mains",
		"avec un micro-penis",
		"qui sait pas conduire",
		"aussi grand(e) que ma grand-mère",
		"drogué(e)",
		"qui pue du cul par la bouche",
		"rotophilé",
		"que j'aime pas",
		"hermaphrodite",
		"très ridé",
		"illetré",
		"qui n'a pas le rythme dans la peau",
		"qui va à la chasse et qui perd sa place",
	}
	return fmt.Sprintf("%s %s %s",
		partA[firstLetterIdx(firstname)%len(partA)],
		partB[firstLetterIdx(lastname)%len(partB)],
		partC[lastLetterIdx(lastname)%len(partC)],
	)
}

func init() {
	Generators = append(Generators, FrInjure{})
}
