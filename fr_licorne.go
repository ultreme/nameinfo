package nameinfo

import "fmt"

type FrLicorne struct{}

func (_ FrLicorne) Name() string { return "nom de licorne" }
func (_ FrLicorne) Generate(firstname, lastname string) string {
	return fmt.Sprintf("%s %s",
		[]string{"arc-en-ciel", "tourbillon", "joyau", "désir", "légende", "harmonie", "pelage", "soleil", "lune", "corne", "lumière", "divinité", "poney", "flocon", "brise", "esprit", "océan", "créature", "coeur", "étoile", "sabot", "sirène", "miracle", "sagesse", "illusion", "fleur"}[firstLetterIdx(firstname)],
		[]string{"intrépide", "ailé(e)", "féérique", "de paillettes", "magique", "étincelant(e)", "agile", "resplendissant(e)", "divin(e)", "brillant(e)", "charmant(e)", "cosmique", "furtif(ve)", "malin(e)", "sauvage", "ardent(e)", "supreme", "infernal(e)", "apocalyptique", "enchanté(e)", "charmeur(se)", "enflammé(e)", "enragé(e)", "amoureux(se)", "malicieux(se)", "fabuleux(se)"}[firstLetterIdx(lastname)],
	)
}

func init() {
	Generators = append(Generators, FrLicorne{})
}
