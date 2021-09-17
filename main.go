package main

import (
	"fmt"
	"strings"
)

func contains(menu []string, str string) bool {
	for _, v := range menu {
		if v == str {
			return true
		}
	}

	return false
}

func main() {

	var pesanan = []string{}
	var prompt string
	var pesan string
	var menu = []string{"tahu", "tempe", "sambal", "lele", "ayam", "nasi"}

	fmt.Println("Toko Makanan Indonesia")
	fmt.Println("=================")
	for _, v := range menu {
		fmt.Printf("* %s\n", v)
	}
	fmt.Println("=================")

	for {
		fmt.Printf("Pilih menu: ")
		fmt.Scan(&pesan)

		if strings.TrimSpace(pesan) == "" {
			fmt.Printf("Tidak dapat memilih menu kosong\n")
		} else {
			cek := contains(menu, pesan)
			if !cek {
				fmt.Printf("Menu tidak ada\n")
			} else {
				pesanan = append(pesanan, pesan)
			}
		}


		fmt.Printf("Lanjut memesan menu lain? (y/t): ")
		fmt.Scan(&prompt)
		
		if strings.ToLower(prompt) == "y" {
			continue
		} else if strings.ToLower(prompt) == "t" {
			break
		} else {
			fmt.Println("Pilihan tidak ada")
			break
		}
	}

	if len(pesanan) == 0 {
		fmt.Println("Maaf anda belum memesan apapun")
	} else {
		fmt.Println("Pesanan anda :")
		for _, v := range pesanan {
			fmt.Printf("- %s\n", v)
		}
	
		fmt.Println("Terimakasih atas pesanan makanan anda!")
	}

	
}
