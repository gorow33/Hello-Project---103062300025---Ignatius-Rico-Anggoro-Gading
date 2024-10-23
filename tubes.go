package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Buku mendefinisikan properti sebuah buku.
type Buku struct {
	Judul   string
	Penulis string
	Tahun   int
}

// perpustakaan menyimpan daftar buku.
var perpustakaan []Buku

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		tampilanMenu()
		choice := userInput(scanner)
		if choice == "5" {
			if len(perpustakaan) > 0 {
				simpanKeFile(perpustakaan[len(perpustakaan)-1])
			}
			fmt.Println("Terima kasih telah menggunakan aplikasi. Buku terakhir telah disimpan.")
			os.Exit(0)
		}
		pilihanMenu(choice, scanner)
	}
}

// tampilanMenu menampilkan menu utama.
func tampilanMenu() {
	fmt.Println("\n=== Aplikasi Perpustakaan ===")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Lihat Buku")
	fmt.Println("3. Update Buku")
	fmt.Println("4. Delete Buku")
	fmt.Println("5. Keluar")
	fmt.Print("Pilih opsi: ")
}

// userInput membaca input dari pengguna.
func userInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

// pilihanMenu menangani pilihan menu pengguna.
func pilihanMenu(choice string, scanner *bufio.Scanner) {
	switch choice {
	case "1":
		tambahBuku(scanner)
	case "2":
		lihatBuku()
	case "3":
		updateBuku(scanner)
	case "4":
		deleteBuku(scanner)
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
}

// tambahBuku menambahkan buku baru ke perpustakaan dan menyimpannya ke file.
func tambahBuku(scanner *bufio.Scanner) {
	fmt.Print("Masukkan Judul Buku: ")
	judul := userInput(scanner)

	fmt.Print("Masukkan Nama Penulis: ")
	penulis := userInput(scanner)

	tahun := inputTahun(scanner)

	bukuBaru := Buku{Judul: judul, Penulis: penulis, Tahun: tahun}
	perpustakaan = append(perpustakaan, bukuBaru)
	fmt.Println("Buku berhasil ditambahkan.")

	// Simpan buku baru ke file setelah ditambahkan.
	simpanKeFile(bukuBaru)
}

// inputTahun membaca dan memvalidasi tahun terbit.
func inputTahun(scanner *bufio.Scanner) int {
	for {
		fmt.Print("Masukkan Tahun Terbit: ")
		if tahun, err := strconv.Atoi(userInput(scanner)); err == nil {
			return tahun
		}
		fmt.Println("Input tidak valid. Masukkan angka.")
	}
}

// lihatBuku menampilkan semua buku di perpustakaan.
func lihatBuku() {
	if len(perpustakaan) == 0 {
		fmt.Println("Perpustakaan kosong.")
		return
	}
	fmt.Println("\nDaftar Buku:")
	for i, buku := range perpustakaan {
		fmt.Printf("%d. %s oleh %s (Terbit: %d)\n", i+1, buku.Judul, buku.Penulis, buku.Tahun)
	}
}

// updateBuku memperbarui informasi buku di perpustakaan.
func updateBuku(scanner *bufio.Scanner) {
	if len(perpustakaan) == 0 {
		fmt.Println("Perpustakaan kosong.")
		return
	}

	lihatBuku()
	fmt.Print("Pilih nomor buku yang ingin di-update: ")
	index, err := strconv.Atoi(userInput(scanner))
	if err != nil || index < 1 || index > len(perpustakaan) {
		fmt.Println("Nomor buku tidak valid.")
		return
	}

	index -= 1 // sesuaikan dengan indeks slice

	fmt.Print("Masukkan Judul Baru (kosongkan jika tidak diubah): ")
	judul := userInput(scanner)
	if judul != "" {
		perpustakaan[index].Judul = judul
	}

	fmt.Print("Masukkan Nama Penulis Baru (kosongkan jika tidak diubah): ")
	penulis := userInput(scanner)
	if penulis != "" {
		perpustakaan[index].Penulis = penulis
	}

	fmt.Print("Masukkan Tahun Terbit Baru (kosongkan jika tidak diubah): ")
	tahunStr := userInput(scanner)
	if tahunStr != "" {
		if tahun, err := strconv.Atoi(tahunStr); err == nil {
			perpustakaan[index].Tahun = tahun
		} else {
			fmt.Println("Input tidak valid. Tahun tidak diubah.")
		}
	}

	fmt.Println("Buku berhasil di-update.")
	// Simpan buku yang di-update ke file.
	simpanKeFile(perpustakaan[index])
}

// deleteBuku menghapus buku dari perpustakaan dan menghapus file terkait.
func deleteBuku(scanner *bufio.Scanner) {
	if len(perpustakaan) == 0 {
		fmt.Println("Perpustakaan kosong.")
		return
	}

	lihatBuku()
	fmt.Print("Pilih nomor buku yang ingin dihapus: ")
	index, err := strconv.Atoi(userInput(scanner))
	if err != nil || index < 1 || index > len(perpustakaan) {
		fmt.Println("Nomor buku tidak valid.")
		return
	}

	index -= 1 // sesuaikan dengan indeks slice

	buku := perpustakaan[index]
	namaFile := strings.ReplaceAll(buku.Judul, " ", "_") + ".txt"
	if err := os.Remove(namaFile); err != nil {
		fmt.Printf("Gagal menghapus file: %v\n", err)
	}

	perpustakaan = append(perpustakaan[:index], perpustakaan[index+1:]...)
	fmt.Println("Buku berhasil dihapus.")
}

// simpanKeFile menyimpan buku ke file txt.
func simpanKeFile(buku Buku) {
	// Membuat nama file dari judul buku dengan mengganti spasi dengan underscore dan menambahkan ekstensi .txt
	namaFile := strings.ReplaceAll(buku.Judul, " ", "_") + ".txt"

	file, err := os.Create(namaFile)
	if err != nil {
		fmt.Printf("Gagal membuat file: %v\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(fmt.Sprintf("%s oleh %s (Terbit: %d)\n", buku.Judul, buku.Penulis, buku.Tahun)); err != nil {
		fmt.Printf("Gagal menulis ke file: %v\n", err)
		return
	}

	if err := writer.Flush(); err != nil {
		fmt.Printf("Gagal menyimpan data: %v\n", err)
		return
	}
	fmt.Printf("Data buku '%s' disimpan ke file '%s'.\n", buku.Judul, namaFile)
}