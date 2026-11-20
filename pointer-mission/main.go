package main

import (
	"fmt"
)

type Sayac struct{ deger int }

func (s Sayac) ArttirKopya()    { s.deger++ }
func (s *Sayac) ArttirPointer() { s.deger++ }

func degistir(s []int)  { s[0] = 100 }
func ekleKopya(s []int) { s = append(s, 4) }
func eklePointer(s []int) []int {
	return append(s, 4)
}

func main() {
	// 1
	x := 1
	p := &x
	*p = 5
	fmt.Println("1:", x)

	// 2
	s := Sayac{}
	s.ArttirKopya()
	fmt.Println("2:", s.deger) // 0, çünkü kopya değişti
	// ve fonksiyon bittiğinde kopya yok oldu.
	s.ArttirPointer()
	s.ArttirPointer()
	fmt.Println("2:", s.deger)

	// 3
	dilim := []int{1, 2, 3}
	degistir(dilim)
	fmt.Println("3:", dilim)

	// 4
	ekleKopya(dilim)
	fmt.Println("4:", dilim)   // 4 eklenmedi, çünkü kopya değişti
	dilim = eklePointer(dilim) // 4 eklendi, çünkü slice başlığı değişti
	fmt.Println("4:", dilim)

	// 5
	var h *Sayac
	fmt.Println("5:", h == nil) // true, çünkü h henüz bir Sayac'ı göstermiyor
	h = &Sayac{}
	fmt.Println("5:", h == nil) // false, çünkü h artık bir Sayac'ı gösteriyor
}
