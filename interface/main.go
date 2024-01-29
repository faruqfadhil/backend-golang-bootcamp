package main

import (
	"fmt"
	"math"
)

type hitung interface{}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

type kendaraan interface {
	gas() string
	rem() string
}

type kendaraanManual interface {
	kendaraan
	injakKopling() string
}

type mobil struct{}

func (m *mobil) gas() string {
	fmt.Println("mobil gas")

	return ""
}

func (m *mobil) rem() string {
	fmt.Println("mobil rem")
	return ""
}

func (m *mobil) injakKopling() string {
	fmt.Println("mobil kopling")
	return ""
}

type motor struct{}

func (m *motor) gas() string {
	fmt.Println("motor gas")
	// bensin masuk ke ruang bakar
	// dynamo bergerak
	// penggerak roda depan maju

	return ""
}

func (m *motor) rem() string {
	fmt.Println("motor rem")
	return ""
}

type pesawat struct{}

func (m *pesawat) gas() string {
	fmt.Println("pesawat gas")
	// avtur masuk ke ruang bakar
	// mesin 1 nyala
	// penggerak roda depan maju
	return ""
}

func (m *pesawat) rem() string {
	fmt.Println("pesawat rem")
	return ""
}

func main() {
	// var bangunDatar hitung
	// psg := persegi{
	// 	sisi: 10.0,
	// }
	// bangunDatar = psg
	// fmt.Println("===== persegi")
	// fmt.Println("luas      :", bangunDatar.luas())
	// fmt.Println("keliling  :", bangunDatar.keliling())

	// bangunDatar = lingkaran{14.0}
	// fmt.Println("===== lingkaran")
	// fmt.Println("luas      :", bangunDatar.luas())
	// fmt.Println("keliling  :", bangunDatar.keliling())
	// fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())

	// today := "selasa"
	// var usageKendaraan kendaraan
	// if today == "senin" {
	// 	// pengen pake mobil.
	// 	usageKendaraan = &mobil{}
	// } else {
	// 	// pengen pake motor.
	// 	usageKendaraan = &motor{}
	// }

	var names any

	a := float64(10)
	// names = []string{"sadjhbsd", "sadd"}
	names = &a

	s := *names.(*float64) * 100
	fmt.Println(s)

	// for _, a := range names.(int) {
	// 	fmt.Println(a)
	// }

	// // kendarai
	// usageKendaraan.gas()
	// usageKendaraan.rem()
}
