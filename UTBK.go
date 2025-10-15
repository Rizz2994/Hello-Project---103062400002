package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX = 100

// Tipe bentukan untuk menyimpan data mahasiswa
// Berisi nama, jurusan, nilai tes, dan status diterima
type Mahasiswa struct {
	Nama      string
	Jurusan   string
	NilaiTest int
	Diterima  bool
}

// Array utama untuk menyimpan data mahasiswa
var dataMahasiswa [MAX]Mahasiswa
var jumlahMahasiswa int

// sequentialSearch mencari mahasiswa berdasarkan nama secara berurutan dari awal hingga akhir array
func sequentialSearch(nama string) int {
	for i := 0; i < jumlahMahasiswa; i++ {
		if strings.EqualFold(dataMahasiswa[i].Nama, nama) {
			return i
		}
	}
	return -1
}

// binarySearch mencari data mahasiswa berdasarkan nama menggunakan algoritma pencarian biner.
// Data mahasiswa harus sudah terurut berdasarkan nama (ascending) sebelum fungsi ini dipanggil.
func binarySearch(nama string) int {
	low := 0
	high := jumlahMahasiswa - 1
	for low <= high {
		mid := (low + high) / 2
		if strings.EqualFold(dataMahasiswa[mid].Nama, nama) {
			return mid
		} else if strings.ToLower(nama) < strings.ToLower(dataMahasiswa[mid].Nama) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// selectionSortNilai mengurutkan data mahasiswa berdasarkan nilai tes
// Gunakan ascending (asc=true) atau descending (asc=false)
func selectionSortNilai(asc bool) {
	for i := 0; i < jumlahMahasiswa-1; i++ {
		idx := i
		for j := i + 1; j < jumlahMahasiswa; j++ {
			if (asc && dataMahasiswa[j].NilaiTest < dataMahasiswa[idx].NilaiTest) ||
				(!asc && dataMahasiswa[j].NilaiTest > dataMahasiswa[idx].NilaiTest) {
				idx = j
			}
		}
		dataMahasiswa[i], dataMahasiswa[idx] = dataMahasiswa[idx], dataMahasiswa[i]
	}
}

// insertionSortNama mengurutkan data mahasiswa berdasarkan nama
// Gunakan ascending (asc=true) atau descending (asc=false)
func insertionSortNama(asc bool) {
	for i := 1; i < jumlahMahasiswa; i++ {
		temp := dataMahasiswa[i]
		j := i - 1
		for (j >= 0) && ((asc && strings.ToLower(dataMahasiswa[j].Nama) > strings.ToLower(temp.Nama)) ||
			(!asc && strings.ToLower(dataMahasiswa[j].Nama) < strings.ToLower(temp.Nama))) {
			dataMahasiswa[j+1] = dataMahasiswa[j]
			j--
		}
		dataMahasiswa[j+1] = temp
	}
}

// tambahMahasiswa menambahkan data baru ke array mahasiswa jika masih ada kapasitas
// Nilai juga menentukan apakah mahasiswa diterima (nilai >= 70)
func tambahMahasiswa(nama, jurusan string, nilai int) {
	if jumlahMahasiswa < MAX {
		dataMahasiswa[jumlahMahasiswa] = Mahasiswa{Nama: nama, Jurusan: jurusan, NilaiTest: nilai, Diterima: nilai >= 600}
		jumlahMahasiswa++
	} else {
		fmt.Println("Kapasitas penuh!")
	}
}

// editMahasiswa mengubah nilai mahasiswa berdasarkan nama menggunakan pencarian sequential
// Status diterima diperbarui sesuai nilai baru
func editMahasiswa(nama string, nilaiBaru int) {
	idx := sequentialSearch(nama)
	if idx != -1 {
		dataMahasiswa[idx].NilaiTest = nilaiBaru
		dataMahasiswa[idx].Diterima = nilaiBaru >= 600
		fmt.Println("Data berhasil diubah.")
	} else {
		fmt.Println("Mahasiswa tidak ditemukan")
	}
}

// hapusMahasiswa menghapus data mahasiswa berdasarkan nama menggunakan pencarian sequential
// Data setelah mahasiswa yang dihapus akan digeser ke kiri
func hapusMahasiswa(nama string) {
	idx := sequentialSearch(nama)
	if idx != -1 {
		for i := idx; i < jumlahMahasiswa-1; i++ {
			dataMahasiswa[i] = dataMahasiswa[i+1]
		}
		jumlahMahasiswa--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Mahasiswa tidak ditemukan")
	}
}

// tampilkanData menampilkan semua data mahasiswa yang ada di array
func tampilkanData() {
	fmt.Println("Data Mahasiswa:")
	for i := 0; i < jumlahMahasiswa; i++ {
		fmt.Printf("%d. %s | Jurusan: %s | Nilai: %d | Status: %s\n",
			i+1,
			dataMahasiswa[i].Nama,
			dataMahasiswa[i].Jurusan,
			dataMahasiswa[i].NilaiTest,
			func(b bool) string {
				if b {
					return "DITERIMA"
				}
				return "DITOLAK"
			}(dataMahasiswa[i].Diterima))
	}
}

// Fungsi utama yang menampilkan menu dan memproses input pengguna untuk operasi pada data mahasiswa
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Tampilkan Data Mahasiswa")
		fmt.Println("3. Urutkan berdasarkan Nilai (ASC/DESC)")
		fmt.Println("4. Urutkan berdasarkan Nama (ASC/DESC)")
		fmt.Println("5. Edit Mahasiswa")
		fmt.Println("6. Hapus Mahasiswa")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu: ")
		pilihStr, _ := reader.ReadString('\n')
		pilihStr = strings.TrimSpace(pilihStr)
		pilih, _ := strconv.Atoi(pilihStr)

		if pilih == 1 {
			fmt.Print("Nama: ")
			nama, _ := reader.ReadString('\n')
			nama = strings.TrimSpace(nama)
			fmt.Print("Jurusan: ")
			jurusan, _ := reader.ReadString('\n')
			jurusan = strings.TrimSpace(jurusan)
			fmt.Print("Nilai Tes: ")
			nilaiStr, _ := reader.ReadString('\n')
			nilaiStr = strings.TrimSpace(nilaiStr)
			nilai, _ := strconv.Atoi(nilaiStr)
			tambahMahasiswa(nama, jurusan, nilai)
		} else if pilih == 2 {
			tampilkanData()
		} else if pilih == 3 {
			fmt.Print("Urutkan ASC (true) atau DESC (false): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			selectionSortNilai(input == "true")
		} else if pilih == 4 {
			fmt.Print("Urutkan ASC (true) atau DESC (false): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			insertionSortNama(input == "true")
		} else if pilih == 5 {
			fmt.Print("Nama Mahasiswa yang akan diedit: ")
			nama, _ := reader.ReadString('\n')
			nama = strings.TrimSpace(nama)
			fmt.Print("Nilai Baru: ")
			nilaiStr, _ := reader.ReadString('\n')
			nilaiStr = strings.TrimSpace(nilaiStr)
			nilai, _ := strconv.Atoi(nilaiStr)
			editMahasiswa(nama, nilai)
		} else if pilih == 6 {
			fmt.Print("Nama Mahasiswa yang akan dihapus: ")
			nama, _ := reader.ReadString('\n')
			nama = strings.TrimSpace(nama)
			hapusMahasiswa(nama)
		} else if pilih == 7 {
			fmt.Println("Terima kasih!")
			return
		} else {
			fmt.Println("Menu tidak valid")
		}
	}
}
