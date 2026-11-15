package main

import "fmt"

func main() {
	// 1) Uzun tanım: tip açıkça yazılır
	var ad string = "Mete"

	// 2) Tip çıkarımı (Java'daki var gibi)
	var yas = 25

	// 3) Kısa tanım := (sadece fonksiyon içinde kullanılabilir)
	boy := 1.80

	// 4) Sabit
	const pi = 3.14159

	// 5) Zero value: değer vermeden tanım
	var sayi int
	var metin string
	var aktif bool
	var oran float64

	fmt.Println(ad, yas, boy, pi)
	fmt.Printf("sayi=%d metin=%q aktif=%t oran=%f\n", sayi, metin, aktif, oran)

	x := 10  // int
	y := 2.5 // float64
	// z := x * y           // DERLEME HATASI
	z := float64(x) * y // dönüşümü açıkça yazmalısın
	fmt.Println(z)      //Go, kullanılmayan yerel değişkenlerde derlemeyi durdurur.
}

//Java’da değer atanmamış bir String alanı null olur ve NullPointerException riski taşır.
//Go’da ise her tipin varsayılan bir sıfır değeri vardır:

//Java int’i sessizce double’a yükseltir, Go bunu yapmaz
