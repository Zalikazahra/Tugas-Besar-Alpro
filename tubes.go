package main
import "fmt"
const MAX = 100
type Barang struct {
	Kode       string
	Nama       string
	StokKg     int 
	HargaPerKg int 
}
var data [MAX]Barang
var n int
func tambah() {
	if n >= MAX {
		fmt.Println("Gudang penuh!")
		return
	}
	fmt.Print("Kode Barang      : ")
	fmt.Scan(&data[n].Kode)

	fmt.Print("Nama Produk      : ")
	fmt.Scan(&data[n].Nama)

	fmt.Print("Stok Awal (kg)   : ")
	fmt.Scan(&data[n].StokKg)

	fmt.Print("Harga (per 1 kg) : Rp ")
	fmt.Scan(&data[n].HargaPerKg)

	n++
	fmt.Println("Data barang berhasil ditambah.")
}

func tampil() {
	if n == 0 {
		fmt.Println("Data kosong.")
		return
	}

	fmt.Println("\n==============================================")
	fmt.Printf("%-6s   %-14s   %-7s   %-10s\n", "Kode", "Nama Produk", "Stok", "Harga/kg")
	fmt.Println("==============================================")

	for i := 0; i < n; i++ {
		fmt.Printf("%-6s   %-14s   %-3d kg   Rp %-10d\n", data[i].Kode, data[i].Nama, data[i].StokKg, data[i].HargaPerKg)
	}
	fmt.Println("==============================================")
}

func ubah() {
	var kode string
	fmt.Print("Masukkan kode barang yang akan diubah: ")
	fmt.Scan(&kode)

	for i := 0; i < n; i++ {
		if data[i].Kode == kode {
			fmt.Print("Nama Produk baru      : ")
			fmt.Scan(&data[i].Nama)

			fmt.Print("Stok baru (kg)        : ")
			fmt.Scan(&data[i].StokKg)

			fmt.Print("Harga baru (per 1 kg) : Rp ")
			fmt.Scan(&data[i].HargaPerKg)

			fmt.Println("Data berhasil diubah!")
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func hapus() {
	var kode string
	fmt.Print("Masukkan kode barang yang akan dihapus: ")
	fmt.Scan(&kode)

	for i := 0; i < n; i++ {
		if data[i].Kode == kode {
			for j := i; j < n-1; j++ {
				data[j] = data[j+1]
			}
			n--
			fmt.Println("Data berhasil dihapus.")
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func masuk() {
	var kode string
	var jumlah int

	fmt.Print("Kode Barang     : ")
	fmt.Scan(&kode)

	fmt.Print("Jumlah Masuk (kg): ")
	fmt.Scan(&jumlah)

	for i := 0; i < n; i++ {
		if data[i].Kode == kode {
			data[i].StokKg += jumlah
			fmt.Println("Stok berhasil ditambah. Stok sekarang:", data[i].StokKg, "kg")
			return
		}
	}
	fmt.Println("Barang tidak ditemukan.")
}

func keluar() {
	var kode string
	var jumlah int

	fmt.Print("Kode Barang      : ")
	fmt.Scan(&kode)

	fmt.Print("Jumlah Keluar (kg): ")
	fmt.Scan(&jumlah)

	for i := 0; i < n; i++ {
		if data[i].Kode == kode {
			if jumlah <= data[i].StokKg {
				data[i].StokKg -= jumlah
				fmt.Println("Stok berhasil dikurangi. Stok sekarang:", data[i].StokKg, "kg")
				
				totalHarga := jumlah * data[i].HargaPerKg
				fmt.Printf("Total Harga yang harus dibayar: Rp %d\n", totalHarga)
			} else {
				fmt.Println("Transaksi Gagal! Stok di gudang tidak mencukupi.")
			}
			return
		}
	}
	fmt.Println("Barang tidak ditemukan.")
}

func sequentialSearch() {
	var key string
	fmt.Print("Cari (Masukkan Kode atau Nama Produk): ")
	fmt.Scan(&key)

	found := false
	for i := 0; i < n; i++ {
		if data[i].Kode == key || data[i].Nama == key {
			if !found {
				fmt.Println("\nData Ditemukan:")
				fmt.Println("==============================================")
				fmt.Printf("%-6s   %-14s   %-7s   %-10s\n", "Kode", "Nama Produk", "Stok", "Harga/kg")
				fmt.Println("==============================================")
				found = true
			}
			fmt.Printf("%-6s   %-14s   %-3d kg   Rp %-10d\n", data[i].Kode, data[i].Nama, data[i].StokKg, data[i].HargaPerKg)
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan.")
	} else {
		fmt.Println("==============================================")
	}
}

func sortKode() {
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if data[j].Kode < data[min].Kode {
				min = j
			}
		}
		temp := data[i]
		data[i] = data[min]
		data[min] = temp
	}
}

func binarySearch() {
	if n == 0 {
		fmt.Println("Data kosong.")
		return
	}

	sortKode()

	var key string
	fmt.Print("Cari Kode Barang: ")
	fmt.Scan(&key)

	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2

		if data[mid].Kode == key {
			fmt.Println("\nData Ditemukan:")
			fmt.Println("==============================================")
			fmt.Printf("%-6s   %-14s   %-7s   %-10s\n", "Kode", "Nama Produk", "Stok", "Harga/kg")
			fmt.Println("==============================================")
			fmt.Printf("%-6s   %-14s   %-3d kg   Rp %-10d\n", data[mid].Kode, data[mid].Nama, data[mid].StokKg, data[mid].HargaPerKg)
			fmt.Println("==============================================")
			return
		}

		if key < data[mid].Kode {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	fmt.Println("Data tidak ditemukan.")
}

func selectionSort() {
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if data[j].StokKg < data[min].StokKg {
				min = j
			}
		}
		temp := data[i]
		data[i] = data[min]
		data[min] = temp
	}
}

func insertionSort() {
	var pass int
	var temp Barang

	for pass = 1; pass < n; pass++ {
		temp = data[pass]
		i := pass
		for i > 0 && temp.StokKg > data[i-1].StokKg {
			data[i] = data[i-1]
			i--
		}
		data[i] = temp
	}
}

func statistik() {
	if n == 0 {
		fmt.Println("Data masih kosong. Tidak dapat menampilkan statistik.")
		return
	}

	totalAset := 0
	min := 0

	for i := 0; i < n; i++ {
		totalAset += data[i].StokKg * data[i].HargaPerKg
		if data[i].StokKg < data[min].StokKg {
			min = i
		}
	}

	fmt.Println("\n========== STATISTIK GUDANG ==========")
	fmt.Printf("Total Nilai Aset Gudang : Rp %d\n", totalAset)
	fmt.Println("Barang dengan Stok Paling Menipis:")
	fmt.Printf("- Kode        : %s\n- Nama Produk : %s\n- Stok        : %d kg\n", data[min].Kode, data[min].Nama, data[min].StokKg)
	fmt.Println("======================================")
}

func main() {
	var pilih int

	for {
		fmt.Println("\n===== SIF GUDANG IN (EDISI SEMBAKO KG) =====")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Ubah Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Transaksi: Barang Masuk (Restock)")
		fmt.Println("5. Transaksi: Barang Keluar (Penjualan)")
		fmt.Println("6. Tampil Semua Data")
		fmt.Println("7. Pencarian: Sequential Search (Kode/Nama)")
		fmt.Println("8. Pencarian: Binary Search (Kode)")
		fmt.Println("9. Urutkan Stok: Terkecil - Terbesar (Selection Sort)")
		fmt.Println("10. Urutkan Stok: Terbesar - Terkecil (Insertion Sort)")
		fmt.Println("11. Statistik Gudang")
		fmt.Println("0. Keluar Aplikasi")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambah()
		case 2:
			ubah()
		case 3:
			hapus()
		case 4:
			masuk()
		case 5:
			keluar()
		case 6:
			tampil()
		case 7:
			sequentialSearch()
		case 8:
			binarySearch()
		case 9:
			selectionSort()
			fmt.Println("\n[Berhasil diurutkan dari STOK TERKECIL ke TERBESAR]")
			tampil()
		case 10:
			insertionSort()
			fmt.Println("\n[Berhasil diurutkan dari STOK TERBESAR ke TERKECIL]")
			tampil()
		case 11:
			statistik()
		case 0:
			fmt.Println("Terima kasih telah menggunakan GudangIn!")
			return
		default:
			fmt.Println("Pilihan menu tidak valid!")
		}
	}
}
