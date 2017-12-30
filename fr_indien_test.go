package nameinfo

import "fmt"

func ExampleFrIndien() {
	fmt.Println(FrIndien{}.Name())
	fmt.Println(FrIndien{}.Generate("Manfred", "Touron"))
	// output:
	// nom indien
	// caribou troublant
}
