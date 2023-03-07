package main

import (
	"fmt"
	"strings"
)

func main() {
	shop := Shop{Products: []Product{
		{items: "Pensil", price: 2500},
		{items: "Pulpen", price: 3000},
		{items: "Buku", price: 4000},
	}}

	isRunning := true
	transaction := Transaction{}

	for isRunning {
		shop.PrintProducts()

		transaction.inputProducts(shop.Products)
		transaction.inputAmount()
		transaction.GoSumSubtotal(transaction.inputSubtotal())

		var MoreProduct string 

		fmt.Print("Mau Tambah Barang Lagi (y/n) ? ")
		fmt.Scanf("%s\n", &MoreProduct)

		isRunning = strings.ToLower(MoreProduct) == "y"

		if !isRunning {
			transaction.printTotal()
			fmt.Println()
			transaction.inputMoney()
			transaction.printChange()
			fmt.Println()

			var moreTransaction string
			fmt.Print("Mau Melakukan Transaksi Lagi (y/n) ? ")
			fmt.Scanf("%s\n", &moreTransaction)
			fmt.Println()

			if strings.ToLower(moreTransaction) == "y" {
				isRunning = strings.ToLower(moreTransaction) == "y"
			} else {
				fmt.Println("Terima Kasih Sampai Jumpa")
			}
		}
	} 
}