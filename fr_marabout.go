package nameinfo

import "fmt"

type FrMarabout struct{}

func (_ FrMarabout) Name() string { return "nom de marabout" }
func (_ FrMarabout) Generate(firstname, lastname string) string {
	partA := []string{
		"Professeur",
		"Mamadou",
		"Monsieur",
		"Guy",
		"Ethanaël",
		"Maitre",
	}
	partB := []string{
		"Bassa",
		"Saibou",
		"Bibi",
		"N'guéyé",
		"Mohamed",
		"Voyance",
		"Bruno",
		"Faroukh",
		"Karamady",
		"Soumah",
	}
	partC := []string{
		"Medium Voyant Authentique",
		"Grand Voyant Medium",
		"Bras droit du tout puissant",
		"Réparation de votre PC par téléphone",
		"Résout tous les problèmes sexuels!",
		"Exhorciste",
		"Homme de parole",
		"Professionnel depuis 36 ans",
		"20 ans d'expérience, paiement apres résultat, 100% de réussite",
	}
	return fmt.Sprintf("%s %s - %s",
		partA[firstLetterIdx(firstname)%len(partA)],
		partB[firstLetterIdx(lastname)%len(partB)],
		partC[lastLetterIdx(lastname)%len(partC)],
	)
}

func init() {
	Generators = append(Generators, FrMarabout{})
}
