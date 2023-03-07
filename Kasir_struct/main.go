package main

import (
	"fmt"
	"strings"
)

func main() {
	shop := Shop{Products: []Product{
		{items: "Pensil", harga: 2500},
		{items: "Pulpen", harga: 3000},
		{items: "Buku", harga: 4000},
	}}

	isRunning := true
	for isRunning {
		shop.PrintProducts()

		transaction := Transaction{}
		transaction.inputProduct(shop.Products)
		transaction.inputAmount()
		transaction.printTotal()
		fmt.Println()
		transaction.inputMoney()
		transaction.printChange()
		fmt.Println()
		var again string
		fmt.Print("Mau Melakukan Transaksi Lagi (y/n) ? \t: ")
		fmt.Scanf("%s\n", &again)
		isRunning = strings.ToUpper(again) == "Y"
	}

	fmt.Println("Terima Kasih Dan Sampai Jumpa Lagi !!")
}