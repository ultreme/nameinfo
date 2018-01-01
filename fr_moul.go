package nameinfo

import (
	"crypto/md5"
	"fmt"
	"strings"
)

type FrMoul struct{}

func (_ FrMoul) Name() string { return "nom moulien" }
func (_ FrMoul) Generate(firstname, lastname string) string {
	partA := []string{"un(e) grand(e)", "un(e)", "un(e) petit(e)", "", "un paquet de", "au moins un(e)", "the ultimate", "", "", "un(e) maxi", "un(e) maxi best-of"}
	partB := []string{"personnage", "copain", "energumène", "vadim", "boite de bonbons", "flocon de soleil", "bonnet tricoté", "vin chaud", "crayon de bois", "poule naine brune", "mec avec un manteau", "nana gentille", "tableau de famille", "verre à pied cassé", "livre de recette", "recette de cuisine"}
	partC := []string{"du dimanche", "en sauce", "sans les mains", "déguisé en catwoman", "au chocolat", "d'amour", "qui me fait penser à un rêve", "jusqu'au bout", "saupoudré", "de cristal", "véritable", "faux-jeton", "avec 5 i", "à sa maman", "à main", "à voile", "à molette", ""} // prédicats cools
	partD := []string{"by night", "", "", "pour toujours", "forever", "corp", "et compagnie", "", "cool"}
	hash := md5.Sum([]byte(strings.ToLower(removeAccents(fmt.Sprintf("%s %s", firstname, lastname)))))
	parts := []string{}
	if part := partA[int(hash[0])%len(partA)]; part != "" {
		parts = append(parts, part)
	}
	if part := partB[int(hash[1])%len(partB)]; part != "" {
		parts = append(parts, part)
	}
	if part := partC[int(hash[2])%len(partC)]; part != "" {
		parts = append(parts, part)
	}
	if part := partD[int(hash[3])%len(partD)]; part != "" {
		parts = append(parts, part)
	}
	return strings.Join(parts, " ")
}

func init() {
	Generators = append(Generators, FrMoul{})
}
