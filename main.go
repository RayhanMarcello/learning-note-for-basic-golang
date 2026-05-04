package main

import (
	"fmt"
)

type Data struct {
	angka1 int
	angka2 int
}

// struct method

type Customer struct {
	Name string
	Age  int
}

func (customer Customer) saysay() {
	fmt.Println("namaku ", customer.Name)
}

func operasi(data *Data) {
	hasil := data.angka1 + data.angka2
	fmt.Println(hasil)
}

// pointer di func
type Address struct {
	Country string
}

func ChangeAddtoIdn(address *Address) {
	address.Country = "indonesia"
}

// func parameter

func sayHai(nama string, umur int) {
	fmt.Println(nama, umur)
}

// func return value

func getHai(nama string) string {
	return "hai " + nama
}

// func multiple values

func getFullName() (string, string) {
	return "cello", "keren"
}

// func named return values

func getCompleteName() (pertama, kedua string) {
	pertama = "kontolodon"
	kedua = "hai"
	return pertama, kedua
}

// defer
func logging() {
	fmt.Println("loging")
}
func runApp() {
	defer logging()
	fmt.Println("app running")
}

func main() {

	// struct literals
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

	// func parameter
	sayHai("kontolodon", 2121)

	// func return value
	res := getHai("rayhan")
	fmt.Println(res)

	// func multiple return value
	firstname, lastname := getFullName()
	fmt.Println(firstname, lastname)

	// jika buuth cuman 1 value
	firstnames, _ := getFullName()
	fmt.Println(firstnames)

	// named return values
	pertama, kedua := getCompleteName()
	fmt.Println(pertama, kedua)

	// defer
	runApp()

	// struct methode
	orang := Customer{Name: "anjayani"}
	orang.saysay()

	// pointer di func
	addre := Address{
		Country: "malay",
	}
	ChangeAddtoIdn(&addre)
	fmt.Println(addre)
}
