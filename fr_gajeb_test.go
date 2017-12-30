package nameinfo

import "fmt"

func ExampleFrGajeb() {
	fmt.Println(FrGajeb{}.Name())
	fmt.Println(FrGajeb{}.Generate("Manfred", "Touron"))
	// output:
	// identité de gajeb
	// un bracelet rose
}
