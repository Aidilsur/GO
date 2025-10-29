package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var nilaiAbsensi = 90

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = nilaiAbsensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)
}