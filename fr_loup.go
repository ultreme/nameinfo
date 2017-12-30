package nameinfo

import "fmt"

type FrLoup struct{}

func (_ FrLoup) Name() string { return "nom de loup" }
func (_ FrLoup) Generate(firstname, lastname string) string {
	return fmt.Sprintf("%s %s",
		[]string{"crépuscule", "nuit", "patte", "éclair", "roc", "chef", "beau/belle", "pluie", "croc", "ombre", "orage", "griffe", "feu/flamme", "guerrier/guerrière", "lune", "tonnerre", "chasseur/chasseresse", "calme", "louveteau", "loupiote", "étoile", "pelage", "songe", "voyage", "sérénité", "soleil", "ténèbre"}[firstLetterIdx(firstname)],
		[]string{"rêveur/rêveuse", "grognon/grognonne", "flamboyant/flamboyante", "parfait/parfaite", "gracieux/gracieuse", "libre", "vengeur/vengeresse", "majestueux/majestueuse", "blanc/blanche", "solitaire", "pataud/pataude", "hurleur/hurleuse", "sage", "rouge", "magique", "sombre", "ebène", "royal/royale", "fier/fière", "lunaire", "loyal/loyale", "peureux/peureuse", "gris/grise", "nocturne", "affranchi/affranchie", "céleste"}[firstLetterIdx(lastname)],
	)
}

func init() {
	Generators = append(Generators, FrLoup{})
}
