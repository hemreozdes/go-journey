package main

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
)

type Hesap struct {
	No       string
	Sahip    string
	bakiye   int64 // küçük harf = paket dışından erişilemez
	islemler []int64
}

// Go’da constructor yoktur. Onun yerine New ile başlayan bir fonksiyon yazmak gelenektir:
func NewHesap(no, sahip string, acilisBakiyesi int64) *Hesap {
	return &Hesap{No: no, Sahip: sahip, bakiye: acilisBakiyesi}
}

func (h *Hesap) Bakiye() int64 {
	return h.bakiye
}

func tlFormat(kurus int64) string {
	isaret := ""
	if kurus < 0 {
		isaret = "-"
		kurus = -kurus
	}
	return fmt.Sprintf("%s%d.%02d TL", isaret, kurus/100, kurus%100)
}

func tutarOku(mesaj string) (int64, error) {
	fmt.Println(mesaj)
	var girdi string
	if _, err := fmt.Scanln(&girdi); err != nil {
		return 0, errors.New("girdi okunamadı")
	}
	tl, err := strconv.ParseInt(girdi, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%q geçerli bir tam sayı değil", girdi)
	}
	return tl * 100, nil // TL -> kuruş
}

/*
❌ Value receiver: h bir KOPYADIR, değişiklik kaybolur
func (h Hesap) ParaYatirCopy(tutar int64) {
	h.bakiye += tutar
}
*/

// ✅ Pointer receiver: h orijinali gösterir, değişiklik kalıcıdır
func (h *Hesap) ParaYatir(tutar int64) error {
	if tutar <= 0 {
		return fmt.Errorf("tutar 0 veya negatif olamaz")
	}
	h.bakiye += tutar
	h.islemler = append(h.islemler, tutar)
	return nil
}

func (h *Hesap) ParaCek(tutar int64) error {
	if tutar <= 0 {
		return fmt.Errorf("tutar 0 veya negatif olamaz")
	} else if tutar > h.bakiye {
		return fmt.Errorf("yetersiz bakiye")
	}
	h.bakiye -= tutar
	h.islemler = append(h.islemler, -tutar)
	return nil
}

func (h *Hesap) IslemGecmisi() []int64 {
	return slices.Clone(h.islemler)
}

func islemFormat(kurus int64) string {
	if kurus > 0 {
		return "+" + tlFormat(kurus)
	}
	return tlFormat(kurus)
}

func main() {
	hesaplar := map[string]*Hesap{
		"0000": NewHesap("0000", "Ali", 100000),
		"1234": NewHesap("1234", "Ayşe", 50000),
		"5678": NewHesap("5678", "Mehmet", 25000),
	}

	fmt.Println("Hesap numarası giriniz:")
	var hesapNo string
	fmt.Scanln(&hesapNo)

	hesap, ok := hesaplar[hesapNo]
	if !ok {
		fmt.Println("Hesap bulunamadı!")
		return
	}
	fmt.Println("Bakiye:", tlFormat(hesap.Bakiye()))

	yatirma, err := tutarOku("Eklenecek tutarı giriniz:")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	if err := hesap.ParaYatir(yatirma); err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println(tlFormat(yatirma), "yatırıldı.")

	cekme, err := tutarOku("Çekilecek tutarı giriniz:")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	if err := hesap.ParaCek(cekme); err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println(tlFormat(cekme), "çekildi.")

	for i, islem := range hesap.IslemGecmisi() {
		fmt.Printf("%d. işlem: %s\n", i+1, islemFormat(islem))
	}
	fmt.Println("Güncel bakiye:", tlFormat(hesaplar[hesapNo].Bakiye()))
}
