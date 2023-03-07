package main

import "fmt"

type Transaction struct {
	Product     Product
	Amount      int
	Money       int
	Change      int
	Total       int
	Subtotal    int
	SumSubtotal []int
}

func (transaction *Transaction) inputProducts(Products []Product) {
	var selectedProduct int
	fmt.Print("Masukkan Pilihan Barang \t: ")
	fmt.Scanf("%d\n", &selectedProduct)
	transaction.Product = Products[selectedProduct-1]
}

func (transaction *Transaction) inputAmount() {
	fmt.Print("Masukkan Jumlah Barang \t\t: ")
	fmt.Scanf("%d\n", &transaction.Amount)
}

func (transaction *Transaction) inputSubtotal() int {
	transaction.Subtotal = transaction.Product.price * transaction.Amount
	fmt.Println("Subtotalnya Adalah \t Rp. ", transaction.Subtotal)
	fmt.Println()
	return transaction.Subtotal
}

func (transaction *Transaction) GoSumSubtotal(savesub int) {
	transaction.SumSubtotal = append(transaction.SumSubtotal, savesub)
}

func (transaction *Transaction) printTotal()  {
	transaction.Total = 0
	for _, sum := range transaction.SumSubtotal {
		transaction.Total += sum
	}

	fmt.Println("Totalnya Adalah \t Rp. ", transaction.Total)

}

func (transaction *Transaction) inputMoney() {
	transaction.Money = 0

	for transaction.Money < transaction.Total {
		fmt.Print("Masukkan Uang Yang Akan Dibayar \t Rp. ")
		fmt.Scanf("%d\n", &transaction.Money)

		if transaction.Money < transaction.Total {
			moreMoney := transaction.Total - transaction.Money
			fmt.Println("Maaf Uang Anda Kurang Rp.", moreMoney)
			fmt.Println()
		}
	}
	for i := range transaction.SumSubtotal {
		transaction.SumSubtotal[i] = 0
	}
}

func (transaction *Transaction) printChange() {
	transaction.Change = transaction.Money - transaction.Total
	if transaction.Money > transaction.Total {
		fmt.Println("Kembalian Anda Sebesar \t Rp. ", transaction.Change)
	} else {
		fmt.Println("Uang Anda Pas")
	}
}
