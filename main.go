package main

import (
	"fmt"
)

type Data struct {
	angka1 int
	angka2 int
}

func operasi(data *Data) {
	hasil := data.angka1 + data.angka2
	fmt.Println(hasil)
}

func main() {
	datas := Data{
		angka1: 1,
		angka2: 2,
	}
	operasi(&datas)

	// array
	var names [3]string
	names[0] = "wow"
	name := [3]string{"kontolodon", "mSD"}

	fmt.Println("array", name, names)

	// slice
	nama := []string{"rayhan", "marcello"}
	namaTambahan := append(nama, "wkw")
	fmt.Println(namaTambahan)

	// make slice
	person := map[string]string{
		"nama": "rayhan",
		"nim":  "11231086",
	}
	fmt.Println(person["nama"])

	// for loops
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}
}
