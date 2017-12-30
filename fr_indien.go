package nameinfo

import "fmt"

type FrIndien struct{}

func (_ FrIndien) Name() string { return "nom indien" }
func (_ FrIndien) Generate(firstname, lastname string) string {
	return fmt.Sprintf("%s %s",
		[]string{"chaton", "aigle", "bison", "blaireau", "cnard", "croco", "dindon", "cheval", "iguane", "renard", "condor", "saumon", "caribou", "grizzly", "aigle", "sanglier", "criquet", "chacal", "crapaud", "castor", "ragondin", "putois", "lama", "vautour", "serpent", "pigeon"}[firstLetterIdx(firstname)],
		[]string{"peteur", "bigleux", "cretin", "malin", "grognon", "cocu", "supreme", "apocalyptique", "boiteux", "mortel", "furtif", "charmeur", "enragé", "infernal", "fétide", "mortel", "merveilleux", "moisi", "festif", "troublant", "libidineux", "degueu", "cynique", "bossu", "moelleux", "intolerant"}[firstLetterIdx(lastname)],
	)
}

func init() {
	Generators = append(Generators, FrIndien{})
}
