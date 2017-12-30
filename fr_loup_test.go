package nameinfo

import "fmt"

func ExampleFrLoup() {
	fmt.Println(FrLoup{}.Name())
	fmt.Println(FrLoup{}.Generate("Manfred", "Touron"))
	// output:
	// loup
	// feu/flamme lunaire
}
