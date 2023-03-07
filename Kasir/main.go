package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	fmt.Println("Selamat Datang Di Toko Saya")
	systemKasir()
}

func systemKasir() {
	fmt.Println("====================================")

	var items = map[string]int{"Pensil": 2500, "Pulpen": 3000, "Buku": 4000}

	fmt.Println("Daftar Barang ")

	for key, val := range items {
		fmt.Println(key, " \t:", val)
	}

	fmt.Println("====================================")

	var jumlahBarang, total, bayar int
	var namabarang string

	fmt.Print("Masukkan Nama Barang : ")
	fmt.Scanf("%s\n", &namabarang)
	fmt.Print("Masukkan Jumlah Barang  : ")
	fmt.Scanf("%d\n", &jumlahBarang)

	total = items[namabarang] * jumlahBarang
	fmt.Printf("Totalnya Adalah 	: Rp. %d\n", total)

	fmt.Println()

	fmt.Print("Masukkan Uang Yang Dibayar : ")
	fmt.Scanf("%d\n", &bayar)
	

	var kembalian int 	
	var uangKurang int
	
	for bayar < total {
		uangKurang = total - bayar
		fmt.Println("Maaf Uang Anda Kurang Rp.", uangKurang)
		fmt.Println()
		fmt.Print("Masukkan Uang Yang Dibayar : ")
		fmt.Scanf("%d\n", &bayar)
	}

	if bayar == total  {
		fmt.Println("Terimakasih uang anda pas ")
	} else {
		kembalian = bayar - total
		fmt.Println("Jumlah Kembalian Anda Adalah Rp.", kembalian)
	}

	ulangi()
}

func ulangi() {
	fmt.Println()

	var ulangi string

	fmt.Print("Mau Melakukan Transaksi Lagi (y/n) ? ")
	fmt.Scanf("%s\n", &ulangi)

	fmt.Println()

	if strings.ToUpper(ulangi) == "Y" {
		main()
	} else {
		fmt.Println("Terima Kasih Dan Sampai Jumpa!!")
	}
}
