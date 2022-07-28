package main

import (
	"fmt"
	"os"
	"testproject/pack"
)

var sec int

func main() {
	fmt.Println("Lutfen Sekil Secin")
	fmt.Println("1-Kare 2-Dikdörtgen 3-Yamuk 4-Daire")
	fmt.Scanf("%d", &sec)

	switch sec {

	case 1:
		var object pack.Kare
		fmt.Println("Kenarı Girin")
		fmt.Scanf("%f", &object.Kenar)
		pack.AlanHesapla(object)
		break
	case 2:
		var object pack.Dikdortgen
		fmt.Println("Kenarlari Girin")
		fmt.Printf("Kisa Kenar :  ")
		fmt.Scanf("%f", &object.KisaKenar)
		fmt.Printf("Uzun Kenar :  ")
		fmt.Scanf("%f", &object.UzunKenar)
		pack.AlanHesapla(object)
		break

	case 3:
		var object pack.Yamuk
		fmt.Println("Verileri Girin")
		fmt.Printf("Ust Taban Uzunlugu:  ")
		fmt.Scanf("%f", &object.UstTaban)
		fmt.Printf("Alt Taban Uzunlugu:  ")
		fmt.Scanf("%f", &object.AltTaban)
		fmt.Printf("Yukseklik :   ")
		fmt.Scanf("%f", &object.Yukseklik)
		pack.AlanHesapla(object)
		break

	case 4:
		var object pack.Daire
		fmt.Println("Verileri Girin")
		fmt.Printf("Yaricap :  ")
		fmt.Scanf("%f", &object.YariCap)
		fmt.Println("PI = 3.14")
		pack.AlanHesapla(object)
		break

	default:
		println("HATALI GIRIS")
		os.Exit(1)
	}

}
