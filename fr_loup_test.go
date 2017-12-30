package nameinfo

import "fmt"

func ExampleFrLoup() {
	fmt.Println(FrLoup{}.Name())
	fmt.Println(FrLoup{}.Generate("Manfred", "Touron"))
	// output:
	// nom de loup
	// feu/flamme lunaire
}
