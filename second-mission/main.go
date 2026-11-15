package main

import "fmt"

func main() {
	var kilo float64 = 82
	boy := 1.80
	vki := kilo / (boy * boy)

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

	for santimetre := 150; santimetre <= 200; santimetre += 10 {
		boy := float64(santimetre) / 100
		vki := kilo / (boy * boy)

		kategori := ""
		switch {
		case vki < 18.5:
			kategori = "Zayıf"
		case vki < 25:
			kategori = "Normal"
		case vki < 30:
			kategori = "Fazla kilolu"
		default:
			kategori = "Obez"
		}

		fmt.Printf("Boy: %.2f m | VKİ: %.2f | Kategori: %s\n", boy, vki, kategori)
	}
}
