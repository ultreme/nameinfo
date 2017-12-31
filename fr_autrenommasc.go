package nameinfo

import "fmt"

type FrAutrenommasc struct{}

func (_ FrAutrenommasc) Name() string { return "nom alternatif masculin" }
func (_ FrAutrenommasc) Generate(firstname, lastname string) string {
	return fmt.Sprintf("%s %s",
		[]string{
			"Zinedine",
			"Dominique",
			"Ugo",
			"Carlos",
			"Albert",
			"Bernard",
			"Rodolphe",
			"Xavier",
			"Juan",
			"Patrick",
			"Nicolas",
			"Octave",
			"Emile",
			"Hector",
			"Kevin",
			"Wolfgang",
			"Guillaume",
			"Yanick",
			"Victor",
			"Quentin",
			"Fabien",
			"Théodore",
			"Manuel",
			"Léonard",
			"Ivan",
			"Steven",
		}[firstLetterIdx(firstname)],
		[]string{
			"MARTIN",
			"THOMAS",
			"BERNARD",
			"PETIT",
			"ROBERT",
			"RICHARD",
			"DURAND",
			"DUBOIS",
			"MOREAU",
			"LAURENT",
			"SIMON",
			"MICHEL",
			"LEFEBVRE",
			"LEROY",
			"ROUX",
			"DAVID",
			"BERTRAND",
			"MOREL",
			"FOURNIER",
			"GIRARD",
			"BONNET",
			"DUPONT",
			"LAMBERT",
			"FONTAINE",
			"ROUSSEAU",
			"VINCENT",
		}[firstLetterIdx(lastname)],
	)
}

func init() {
	Generators = append(Generators, FrAutrenommasc{})
}
