package nameinfo

import "fmt"

func ExampleFrIndien() {
	fmt.Println(FrIndien{}.Name())
	fmt.Println(FrIndien{}.Generate("Manfred", "Touron"))
	// output:
	// indien
	// caribou troublant
}
