package main

import "fmt"

func main() {
	var kilo float64 = 82
	boy := 1.80
	vki := kilo / (boy * boy)

	//Parantez yok, süslü parantez zorunlu
	if vki < 18.5 {
		fmt.Println("Zayıf")
	} else if vki < 25 {
		fmt.Println("Normal")
	} else {
		fmt.Println("Fazla kilolu")
	}

	if vki := kilo / (boy * boy); vki >= 30 {
		fmt.Println("Obez:", vki)
	}
	//if içinde kısa tanım yapabilirsin

	//Go’da while ve do-while yoktur. Hepsi for ile yazılır
	// Klasik
	for i := 0; i < 5; i++ {
	}

	// while karşılığı
	sayac := 0
	for sayac < 5 {
		sayac++
	}

	// Sonsuz döngü (bankamatik menüsünde kullanacağız)
	/*
		for {
			if kosul { break }
		}
	*/

	switch {
	case vki < 18.5:
		fmt.Println("Zayıf")
	case vki < 25:
		fmt.Println("Normal")
	case vki < 30:
		fmt.Println("Fazla kilolu")
	default:
		fmt.Println("Obez")
	}
	//break gerekmez
}
