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
	bakiye   int64
	islemler []int64
}

func NewHesap(no, sahip string, acilisBakiyesi int64) *Hesap {
	return &Hesap{No: no, Sahip: sahip, bakiye: acilisBakiyesi}
}
func (h *Hesap) Bakiye() int64 {
	return h.bakiye
}

func (h *Hesap) ParaYatir(tutar int64) error {
	if tutar <= 0 {
		return errors.New("tutar 0 veya negatif olamaz")
	}
	h.bakiye += tutar
	h.islemler = append(h.islemler, tutar)
	return nil
}

func (h *Hesap) ParaCek(tutar int64) error {
	if tutar <= 0 {
		return errors.New("tutar 0 veya negatif olamaz")
	}
	if tutar > h.bakiye {
		return errors.New("yetersiz bakiye")
	}
	h.bakiye -= tutar
	h.islemler = append(h.islemler, -tutar)
	return nil
}

func (h *Hesap) IslemGecmisi() []int64 {
	return slices.Clone(h.islemler)
}

// fmt.Stringer interface'ini örtük olarak sağlar
func (h *Hesap) String() string {
	return fmt.Sprintf("[%s] %s - %s", h.No, h.Sahip, tlFormat(h.bakiye))
}

// ===================== REPOSITORY =====================
type HesapDeposu interface {
	Bul(no string) (*Hesap, error)
	Kaydet(h *Hesap) error
}
type BellekDepo struct {
	hesaplar map[string]*Hesap
}

// Derleme zamanı kontrolü: *BellekDepo, HesapDeposu'nu sağlamak zorunda
var _ HesapDeposu = (*BellekDepo)(nil)

func NewBellekDepo() *BellekDepo {
	return &BellekDepo{hesaplar: make(map[string]*Hesap)}
}
func (d *BellekDepo) Bul(no string) (*Hesap, error) {
	h, ok := d.hesaplar[no]
	if !ok {
		return nil, errors.New("hesap bulunamadı")
	}
	return h, nil
}
func (d *BellekDepo) Kaydet(h *Hesap) error {
	if h == nil {
		return errors.New("nil hesap kaydedilemez")
	}
	d.hesaplar[h.No] = h
	return nil
}

// ===================== YARDIMCILAR =====================
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
	return tl * 100, nil
}

func main() {
	var depo HesapDeposu = NewBellekDepo()

	baslangic := []*Hesap{
		NewHesap("0000", "Ali", 100000),
		NewHesap("1234", "Ayşe", 50000),
		NewHesap("5678", "Mehmet", 25000),
	}

	for _, h := range baslangic {
		if err := depo.Kaydet(h); err != nil {
			fmt.Println("Kayıt hatası:", err)
			return
		}
	}

	fmt.Println("Hesap numarası giriniz:")
	var hesapNo string
	fmt.Scanln(&hesapNo)
	hesap, err := depo.Bul(hesapNo)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println(hesap) // String() otomatik çağrılır

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

	// Depodan tekrar okuyarak güncel olduğunu doğrula
	guncel, _ := depo.Bul(hesapNo)
	fmt.Println("Depodaki son durum:", guncel)
}
