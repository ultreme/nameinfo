package nameinfo

import "fmt"

func ExampleFrAutrenommasc() {
	fmt.Println(FrAutrenommasc{}.Name())
	fmt.Println(FrAutrenommasc{}.Generate("Manfred", "Touron"))
	// output:
	// nom alternatif masculin
	// Emile GIRARD
}
