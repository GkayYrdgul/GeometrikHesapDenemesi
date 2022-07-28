package pack

import "fmt"

const PI = 3.14

type AlanHesabi interface {
	Alan() float64
}

type Kare struct {
	Kenar float64
}
type Dikdortgen struct {
	KisaKenar float64
	UzunKenar float64
}
type Yamuk struct {
	AltTaban  float64
	UstTaban  float64
	Yukseklik float64
}
type Daire struct {
	YariCap float64
}

func (kar Kare) Alan() float64 {
	return kar.Kenar * kar.Kenar
}

func (dik Dikdortgen) Alan() float64 {
	return dik.UzunKenar * dik.KisaKenar
}

func (yam Yamuk) Alan() float64 {
	return (yam.UstTaban + yam.UstTaban) * yam.Yukseklik / 2
}

func (da Daire) Alan() float64 {
	return PI * da.YariCap * da.YariCap
}

func AlanHesapla(object AlanHesabi) {
	fmt.Printf("\nSONUC: %f ", float32(object.Alan()))
}
