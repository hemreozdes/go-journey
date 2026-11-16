package main

import "fmt"

func tlFormat(kurus int64) string {
	isaret := ""
	if kurus < 0 {
		isaret = "-"
		kurus = -kurus
	}
	return fmt.Sprintf("%s%d.%02d TL", isaret, kurus/100, kurus%100)
}

func islemFormat(kurus int64) string {
	if kurus > 0 {
		return "+" + tlFormat(kurus)
	}
	return tlFormat(kurus)
}

func main() {
	hesaplar := map[string]int64{
		"0000": 100000,
		"1234": 50000,
		"5678": 25000,
	}

	fmt.Println("Hesap numarası giriniz:")
	var hesapNo string
	fmt.Scanln(&hesapNo)

	bakiye, ok := hesaplar[hesapNo]
	if !ok {
		fmt.Println("Hesap bulunamadı!")
		return
	}
	fmt.Println("Hesap bulundu! ->", tlFormat(bakiye))

	islemler := []int64{}

	// 50.00 TL yatırma
	yatirma := int64(5000)
	bakiye += (yatirma)
	islemler = append(islemler, yatirma)
	fmt.Println(tlFormat(yatirma), "yatırıldı.")

	// 20.00 TL çekme
	cekme := int64(2000)
	if (cekme) > bakiye {
		fmt.Println("Yetersiz bakiye.")
	} else {
		bakiye -= cekme
		islemler = append(islemler, -cekme)
		fmt.Println(tlFormat(cekme), "çekildi.")
	}

	hesaplar[hesapNo] = bakiye

	for i, tutar := range islemler {
		fmt.Printf("%d. işlem: %s\n", i+1, islemFormat(tutar))
	}

	fmt.Println("Güncel bakiye:", tlFormat(bakiye))
	fmt.Println("Map'teki bakiye:", tlFormat(hesaplar[hesapNo]))

	// tlFormat sınır testi
	fmt.Println("Test -50:", tlFormat(-50))
	fmt.Println("Test 5:", tlFormat(5))
}
