package main

import (
	"fmt"
)

var (
	Saldo_Awal       = 3500000
	Saldo            = Saldo_Awal
	HistoriTransaksi []string
)

func LihatSaldo() {
	fmt.Printf("Saldo Anda Saat Ini Adalah: Rp%d\n", Saldo)
}

func TambahSaldo() {
	var jumlah int

	for {
	fmt.Print("Masukan Jumlah Yang Ingin Ditambahkan: Rp ")
	fmt.Scan(&jumlah)

	if jumlah <= 0 {
		fmt.Println("Jumlah harus positif. Silahkan coba lagi")
		continue // Meminta input lagi
		}
		break // Keluar dari loop jika input valid
	} 
	Saldo += jumlah
	HistoriTransaksi = append(HistoriTransaksi, fmt.Sprintf("Tambah saldo: Rp%d", jumlah))
	fmt.Println("Saldo berhasil ditambahkan.")
	LihatSaldo()
}

func TarikSaldo() {
	var jumlah int

	for {
		fmt.Print("Masukan Jumlah Yang Ingin Ditarik: Rp ")
		fmt.Scan(&jumlah)

		if jumlah <= 0 {
			fmt.Println("Jumlah harus positif. Silakan coba lagi.")
			continue // Meminta input lagi
		}
		if jumlah > Saldo {
			fmt.Println("Saldo tidak cukup. Silakan coba lagi.")
			continue // Meminta input lagi
		}
		break // Keluar dari loop jika input valid
	}

	Saldo -= jumlah
	HistoriTransaksi = append(HistoriTransaksi, fmt.Sprintf("Tarik saldo: Rp%d", jumlah))
	fmt.Println("Saldo berhasil ditarik.")
	LihatSaldo()
}

func LihatAkun(Username, Password string) {
	if Username == "Nurul" && Password == "2406361284" {
		fmt.Printf("Username: %s, Password: %s\n", Username, Password)
		} else { 
		fmt.Println("Akun Tidak Terdaftar")
		}
	}

func LihatHistoriTransaksi() {
	fmt.Println("=== Histori Transaksi ===")
	for _, transaksi := range HistoriTransaksi {
		fmt.Println(transaksi)
	}
}

func main() {
	var Username string
	var Password string

	for {
		//verifikasi akun
		fmt.Print("Masukan Username Anda: ")
		fmt.Scanln(&Username)

		fmt.Print("Masukan Password: ")
		fmt.Scanln(&Password)

		if Username == "Nurul" && Password == "2406361284" {
			fmt.Println("Verifikasi Anda berhasil. Selamat datang!")
			for {
				fmt.Println("\n=== Menu ATM ===")
				fmt.Println("1. Lihat Saldo")
				fmt.Println("2. Tambah Saldo")
				fmt.Println("3. Tarik Saldo")
				fmt.Println("4. Lihat Akun")
				fmt.Println("5. Lihat Histori Transaksi")
				fmt.Println("6. Keluar dari Program")

				var pilihan int
				fmt.Print("Silahkan pilih menu: ")
				fmt.Scan(&pilihan)

				switch pilihan {
				case 1:
					LihatSaldo()
				case 2:
					TambahSaldo()
				case 3:
					TarikSaldo()
				case 4:
					LihatAkun(Username, Password)
				case 5:
					LihatHistoriTransaksi()
				case 6:
					fmt.Println("Terima kasih! Silahkan datang kembali.")
					return
				default:
					fmt.Println("Pilihan tidak ada.")
			}
		}
		} else {
			fmt.Println("Username atau Password Anda salah.")
			continue //Minta input ulang
		} 
		break // Keluar dari loop jika input valid
	}
} 
