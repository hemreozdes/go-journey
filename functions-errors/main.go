package main

import (
	"errors"
	"fmt"
)

func vkiHesapla(kilo, boy float64) (float64, error) {
	if boy <= 0 {
		return 0, errors.New("boy sıfırdan büyük olmalı")
	}
	if boy > 3 {
		return 0, fmt.Errorf("geçersiz boy: %.2f", boy)
	}
	if kilo <= 0 {
		return 0, fmt.Errorf("geçersiz kilo: %.1f", kilo)
	}
	return kilo / (boy * boy), nil
}

func bolVeKalan(a, b int) (int, int) {
	return a / b, a % b
}

func main() {

	var kilo float64 = 82
	boy := 1.80
	vki := kilo / (boy * boy)
	var bolum int
	var kalan int

	bolum, kalan = bolVeKalan(17, 5) // 3, 2
	fmt.Printf("Bölüm: %d, Kalan: %d\n", bolum, kalan)

	vki, err := vkiHesapla(82, 0)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Printf("VKİ: %.2f\n", vki)
}
