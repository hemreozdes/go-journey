package main

import "fmt"

func main() {
	var gunler [3]string // ["", "", ""]  zero value
	sayilar := [3]int{10, 20, 30}
	fmt.Println(gunler)
	fmt.Println(sayilar)

	islemler := []int64{}             // boş slice // java arraylist gibi
	islemler = append(islemler, 5000) // ekle
	islemler = append(islemler, -2000)

	fmt.Println(len(islemler)) // 2
	fmt.Println(islemler[0])   // 5000
	fmt.Println(islemler[1:])  // [-2000]  alt dilim (Python'daki gibi)

	for i, tutar := range islemler {
		fmt.Printf("%d. işlem: %d\n", i+1, tutar)
	}
	// İndekse ihtiyacın yoksa
	// for _, tutar := range islemler { ... }

	bakiyeler := map[string]int64{ //java'daki HashMap gibi
		"1234": 150000, // kuruş cinsinden = 1500.00 TL
		"5678": 42000,
	}

	bakiyeler["9999"] = 0     // ekle veya güncelle
	delete(bakiyeler, "5678") // sil

	//bakiye := bakiyeler["0000"]   // olmayan anahtar → 0 döner, hata yok!
	//Java’da olmayan bir anahtar null döner. Go’da ise zero value döner

	bakiye, ok := bakiyeler["0000"]
	if !ok {
		fmt.Println("Hesap bulunamadı! -> Bakiye:", bakiye)
	}

	bakiye, ok = bakiyeler["1234"]
	if !ok {
		fmt.Println("Hesap bulunamadı! -> Bakiye:", bakiye)
	}
	fmt.Println("Bakiye:", bakiye)
}
