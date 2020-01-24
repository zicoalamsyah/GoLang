package main

import (
	z "fmt"

	"github.com/zicoalamsyah/GoLang/Helloww/helper"
)

func main() {
	z.Println("Hello World")
	z.Println(helper.Tambah(5, 7))
	z.Println(helper.Kurang(7, 5))
	z.Println(helper.Kali(7, 5))
}
