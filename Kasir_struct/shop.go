package main

import "fmt"

type Shop struct {
	Products []Product
}

func (shop *Shop)PrintProducts() {
	fmt.Println("Selamat Datang Di Toko Saya")
	fmt.Println("=============================")
	for i, Product := range shop.Products {
		fmt.Println(i+1, ".", Product.items, "\t Rp.", Product.harga)
	}
}