package main

import "fmt"

func main() {
	yas := 25
	var kilo float64 = 82
	boy := 1.80

	vki := kilo / (boy * boy)
	fmt.Printf("VKI: %.2f\n", vki)

	fmt.Println(float64(yas) * vki)

	fmt.Printf("vki: %T, yas: %T, kilo: %T, boy: %T\n", vki, yas, kilo, boy)
}
