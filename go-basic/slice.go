package main

import "fmt"

func main() {
	months := [...]string{"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Ags", "Sep", "Sep", "Okt", "Nov", "Des"}

	slice1 := months[4:8]
	fmt.Println(slice1)

	slice2 := months[:5]
	fmt.Println(slice2)

	slice3 := months[10:]
	fmt.Println(slice3)
	
	slice4 := months[:]
	fmt.Println(slice4)

	// bentuk fullnya
	// var slice4 []string = months[:]
	// fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daysSlice1 := days[5:]
	daysSlice1[0] = "Sabtu Baru"
	daysSlice1[1] = "Minggu Baru"
	fmt.Println(days)
	
	daysSlice2 := append(daysSlice1, "Libur Baru")
	daysSlice2[0] = "Ups"
	fmt.Println(daysSlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Aidil"
	newSlice[1] = "Surya"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	
	newSlice2 := append(newSlice, "TES")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))
	
	newSlice2[0] = "ubah"
	fmt.Println(newSlice)
	fmt.Println(newSlice2)
	
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)
	fmt.Println(fromSlice)
	
	// perbedaan slice dan array
	iniArray := [...]int{1,2,3,4,5} //ada ... dalam []
	iniSlice := []int{1,2,3,4,5} // tidak ada ... dalam []
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}