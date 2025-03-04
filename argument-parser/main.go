package main

import (
	"argument-parser/parser"
	"fmt"
)

func main() {

	parser := parser.NewArgumentParser(
		"Argument Parser Aracı",
		"Bu araç, komut satırı argümanlarını işler.",
		"Kullanım: go run main.go --name <isim> --age <yaş> --debug",
	)

	parser.AddArgument("name", "Kullanıcı ismi (zorunlu)", true, "")
	parser.AddArgument("age", "Kullanıcı yaşı (isteğe bağlı)", false, "18")
	parser.AddArgument("debug", "Debug modu (isteğe bağlı)", false, "false")

	err := parser.Parse()
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("İsim:", parser.Get("name"))
	fmt.Println("Yaş:", parser.Get("age"))
	fmt.Println("Debug:", parser.Get("debug"))
}
