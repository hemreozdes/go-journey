package main

import (
	"fmt"
)

func ikiKatiDeger(n int)           { n = n * 2 }    // kopyayı değiştirir
func ikiKatiDegerGöster(n int) int { return n * 2 } // kopyayı değiştirir, sonucu döndürür
func ikiKatiPointer(n *int)        { *n = *n * 2 }  // orijinali değiştirir

func main() {
	x := 10
	p := &x         // & → "x'in adresini ver"   p'nin tipi: *int
	fmt.Println(p)  // 0xc000012345 gibi bir bellek adresi
	fmt.Println(*p) // * → "adresteki değeri ver"  10

	*p = 20        // adresteki değeri değiştir
	fmt.Println(x) // 20, çünkü p, x'in kendisini gösteriyor

	sayi := 5
	ikiKatiDeger(sayi)                    // sayi hâlâ 5
	fmt.Println(sayi)                     // 5
	ikiKatiDegerGöster(sayi)              // sayi hâlâ 5
	fmt.Println(sayi)                     // 5
	fmt.Println(ikiKatiDegerGöster(sayi)) // 10, ama sayi hâlâ 5
	ikiKatiPointer(&sayi)                 // sayi artık 10
	fmt.Println(sayi)                     // 10

	a := []int{1, 2, 3}
	b := a // sadece başlık kopyalanır, dizi paylaşılır
	b[0] = 99
	fmt.Println(a) // [99 2 3]  ← a da değişti!

	k := make([]int, 3, 10) // len=3, cap=10

	k[0] = 10
	k[1] = 20
	k[2] = 30

	// Değerleri sırayla ekle
	y := append(k, 4)
	y = append(y, 5)
	y = append(y, 6)
	fmt.Println("İlk değerler:", y)

	_ = append(k[:0], 100) // 0. indeks: 10 -> 100
	fmt.Println("1. ezmeden sonra:", y)

	_ = append(k[:1], 200) // 1. indeks: 20 -> 200
	fmt.Println("2. ezmeden sonra:", y)

	_ = append(k[:2], 300) // 2. indeks: 30 -> 300
	fmt.Println("3. ezmeden sonra:", y)

	// k ile y aynı alttaki diziyi kullanıyor.
	// Her append, ilgili konumdaki değeri değiştiriyor.
	y = append(y, 40) // 3. indeks: 4 -> 40
	fmt.Println("4. ezmeden sonra:", y)

	y = append(y, 50) // 4. indeks: 5 -> 50
	fmt.Println("5. ezmeden sonra:", y)

	y = append(y, 60) // 5. indeks: 6 -> 60
	fmt.Println("6. ezmeden sonra:", y)
}
