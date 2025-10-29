package main

import "fmt"

func main() {
	type NoKTP string
	var KtpAidil NoKTP = "111111111"

	var contoh string = "22222222"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(KtpAidil)
	fmt.Println(contohKTP)
}