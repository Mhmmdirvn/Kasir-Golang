package main

import (
	"fmt"
)

type Transaction struct {
	Product Product
	Amount  int
	Money   int
	Change  int
	Total   int
}

func (transaction *Transaction) inputProduct(Products []Product) {
	var selectedProduct int
	fmt.Print("Masukkan Pilihan Barang \t: ")
	fmt.Scanf("%d\n", &selectedProduct)
	transaction.Product = Products[selectedProduct-1]
}

func (transaction *Transaction) inputAmount() {
	fmt.Print("Masukkan Jumlah Barang \t: ")
	fmt.Scanf("%d\n", &transaction.Amount)
}

func (transaction *Transaction) printTotal() {
	transaction.Total = transaction.Product.harga * transaction.Amount
	fmt.Println("Totalnya Adalah \t Rp. ", transaction.Total)
}

func (transaction *Transaction) inputMoney() {
	transaction.Money = 0
	for transaction.Money < transaction.Total {
		fmt.Print("Masukkan Uang Yang Akan Di Bayar ")
		fmt.Scanf("%d\n", &transaction.Money)

		if transaction.Money < transaction.Total {
			moreMoney := transaction.Total - transaction.Money
			fmt.Println("Maaf Uang Anda Kurang \t Rp.", moreMoney)
			fmt.Println()
		}
	}
}

func (transaction *Transaction) printChange() {
	transaction.Change = transaction.Money - transaction.Total
	if transaction.Money > transaction.Total {
		fmt.Println("Kembalian Anda Sebesar \t Rp.", transaction.Change)
	}else {
		fmt.Println("Uang Anda Pas")
	}
}

