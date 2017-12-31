package nameinfo

import "fmt"

type FrSuperheros struct{}

func (_ FrSuperheros) Name() string { return "nom de super héros" }
func (_ FrSuperheros) Generate(firstname, lastname string) string {
	return fmt.Sprintf("%s %s",
		[]string{"super", "méga", "docteur(e)", "capitaine", "l'incroyable", "le/la redoutable", "l'invincible", "fatal(e)", "agent(e)", "l'extraordinaire", "lieutenant", "commandant", "dark", "aqua", "rocket", "inspécteur", "l'invisible", "colonel", "le/la mystérieux(se)", "amiral", "ultra", "evil", "le/la fantastique", "professeur(e)", "turbo", "le/la spectaculaire"}[firstLetterIdx(firstname)],
		[]string{"géant(e)", "araignée", "vampire", "tornade", "ninja", "mutant", "chaton", "machine", "vengeance", "rockstar", "éclair", "licorne", "jedi", "faucon", "sirène", "cobra", "moutarde", "robot", "sorcier(ère)", "chose", "zombie", "deadpool", "fantôme", "poutine extra-fromage", "pouliche", "destructeur"}[firstLetterIdx(lastname)],
	)
}

func init() {
	Generators = append(Generators, FrSuperheros{})
}
