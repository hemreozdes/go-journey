package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func sayiOku(mesaj string) (float64, error) {
	fmt.Print(mesaj)

	// 1) Girdiyi string olarak oku
	var girdi string
	if _, err := fmt.Scan(&girdi); err != nil {
		return 0, errors.New("girdi okunamadı")
	}

	// 2) Virgülleri noktaya çevir: "1,80" -> "1.80"
	girdi = strings.ReplaceAll(girdi, ",", ".")

	// 3) String'i float64'e dönüştür
	sayi, err := strconv.ParseFloat(girdi, 64)
	if err != nil {
		return 0, fmt.Errorf("%q geçerli bir sayı değil", girdi)
	}

	// 4) Başarılı: sayı ve nil hata
	return sayi, nil
}

func vkiHesapla(kilo, boy float64) (float64, error) {
	if boy <= 0.5 || boy > 3 {
		return 0, errors.New("geçersiz boy: 0.5 ile 3 arasında olmalıdır")
	}
	if kilo <= 0 {
		return 0, fmt.Errorf("geçersiz kilo: %.1f", kilo)
	}
	return kilo / (boy * boy), nil
}

func kategoriBelirle(vki float64) string {
	switch {
	case vki < 18.5:
		return "Zayıf"
	case vki < 25:
		return "Normal"
	case vki < 30:
		return "Fazla kilolu"
	default:
		return "Obez"
	}
}

func main() {

	kilo, err := sayiOku("Kilonuzu (kg) girin: ")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	boy, err := sayiOku("Boyunuzu (metre, ör. 1.80) girin: ")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	vki, err := vkiHesapla(kilo, boy) // önce doğrula + hesapla
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Printf("VKİ: %.2f | Kategori: %s\n", vki, kategoriBelirle(vki))
}
