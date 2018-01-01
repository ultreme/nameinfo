package nameinfo

import "fmt"

type FrZinzin struct{}

func (_ FrZinzin) Name() string { return "nom de zinzin" }
func (_ FrZinzin) Generate(firstname, lastname string) string {
	return fmt.Sprintf("%s %s",
		[]string{
			"inceste",
			"chips",
			"ténia",
			"abruti",
			"cagnotte",
			"blouson",
			"néo-nazi",
			"grec",
			"tasse",
			"porte",
			"morpion",
			"tapis",
			"chaussure",
			"chapeau chinois",
			"accordéon",
			"pick-pocket",
			"un gentrifié",
			"psychopate",
			"filou",
			"moule à gauffres",
			"crétin des alpes",
			"super canard",
			"cannibale",
			"roller",
			"traiteur pour chien",
			"gars cool",
		}[firstLetterIdx(firstname)],
		[]string{
			"familial",
			"superbe",
			"bleu",
			"incinéré",
			"irracible",
			"malhonette",
			"cool",
			"zinzin",
			"circulaire",
			"fermée",
			"sale",
			"jaune",
			"indépendante",
			"partouweur de droite",
			"fanatiqwue",
			"primordial",
			"extremiste",
			"salutaire",
			"perenisé",
			"cuit",
			"mariné",
			"enjolivé",
			"aggressif",
			"furieux",
			"somalien",
			"trisomique",
		}[firstLetterIdx(lastname)],
	)
}

func init() {
	Generators = append(Generators, FrZinzin{})
}
