package cinemaxxi_layout

import (
	"fmt"
	"time"
)

type StudioBioskop struct {
	JumlahKursi      int
	LableKursi       string
	Menu             string
	KursiTerjual     map[string]bool
	LaporanPenjualan []Penjualan
}

type Penjualan struct {
	NomorKursi     string
	WaktuPenjualan time.Time
}

func NewStudioBioskop() *StudioBioskop {
	return &StudioBioskop{
		KursiTerjual: make(map[string]bool),
	}
}

func (s *StudioBioskop) KonfigurasiDenah() {
	fmt.Println("=============== Selamat Datang (Cinema XXI), Silahkan masukkan konfigurasi denah studio ===============")
	fmt.Print("$ Masukkan Label Kursi: ")
	fmt.Scan(&s.LableKursi)
	// labelKursi := "A" // Gantilah sesuai kebutuhan
	fmt.Print("$ Masukkan Jumlah Kursi: ")
	fmt.Scan(&s.JumlahKursi)
	fmt.Printf("Denah Studio berhasil dikonfigurasi dengan label kursi %s dan jumlah kursi %d.\n", s.LableKursi, s.JumlahKursi)
}
func (s *StudioBioskop) KonfigurasiMenu() {
	fmt.Println("=================== Aplikasi Cinema XXI Layout (kursi tersedia A-5) ===================")
	fmt.Println("=================== Pilih salah satu menu di bawah ini ===================")
	fmt.Println("A) Pemesanan Kursi —> book_seat {seat_code}")
	fmt.Println("B) Batalkan Kursi —> cancel_seat {seat_code}")
	fmt.Println("C) Laporan Denah —> seats_status")
	fmt.Println("D) Laporan Transaksi —> transaction_status")
	fmt.Print("$ Masukkan: ")
	fmt.Scan(&s.Menu)
}

func (s *StudioBioskop) TampilkanStatusDenah() {
	denah := s.buatDenah()
	for i := 0; i < len(denah); i += 5 {
		fmt.Println(denah[i : i+5])
	}
}

func (s *StudioBioskop) buatDenah() string {
	denah := ""
	for i := 1; i <= s.JumlahKursi; i++ {
		labelKursi := fmt.Sprintf("A%d", i)
		if s.KursiTerjual[labelKursi] {
			denah += "X " // Kursi sudah terjual
		} else {
			denah += "O " // Kursi tersedia
		}
	}
	return denah
}

func (s *StudioBioskop) BeliTiket(nomorKursi string) {
	if s.KursiTerjual[nomorKursi] {
		fmt.Println("Maaf, kursi sudah terjual.")
	} else {
		s.KursiTerjual[nomorKursi] = true
		waktuPenjualan := time.Now()
		penjualan := Penjualan{NomorKursi: nomorKursi, WaktuPenjualan: waktuPenjualan}
		s.LaporanPenjualan = append(s.LaporanPenjualan, penjualan)
		fmt.Printf("Tiket untuk kursi %s berhasil terjual pada %s.\n", nomorKursi, waktuPenjualan.Format(time.RFC3339))
	}
}

func (s *StudioBioskop) BatalkanPembelian(nomorKursi string) {
	if s.KursiTerjual[nomorKursi] {
		delete(s.KursiTerjual, nomorKursi)
		fmt.Printf("Pembelian tiket untuk kursi %s berhasil dibatalkan.\n", nomorKursi)
	} else {
		fmt.Println("Maaf, kursi tidak terjual.")
	}
}

func (s *StudioBioskop) TampilkanLaporanPenjualan() {
	fmt.Println("Laporan Penjualan:")
	for _, penjualan := range s.LaporanPenjualan {
		fmt.Printf("Kursi %s terjual pada %s\n", penjualan.NomorKursi, penjualan.WaktuPenjualan.Format(time.RFC3339))
	}
}

func main() {
	studioBioskop := NewStudioBioskop()
	studioBioskop.KonfigurasiDenah()
	studioBioskop.KonfigurasiMenu()

	// scanner := bufio.NewScanner(os.Stdin)

	// for {
	// 	fmt.Print("\nMasukkan perintah: ")
	// 	scanner.Scan()
	// 	command := scanner.Text()

	// 	switch {
	// 	case strings.HasPrefix(command, "book_seat"):
	// 		parts := strings.Fields(command)
	// 		if len(parts) == 2 {
	// 			studioBioskop.BeliTiket(parts[1])
	// 			studioBioskop.TampilkanStatusDenah()
	// 		} else {
	// 			fmt.Println("Format perintah tidak valid.")
	// 		}

	// 	case strings.HasPrefix(command, "cancel_seat"):
	// 		parts := strings.Fields(command)
	// 		if len(parts) == 2 {
	// 			studioBioskop.BatalkanPembelian(parts[1])
	// 			studioBioskop.TampilkanStatusDenah()
	// 		} else {
	// 			fmt.Println("Format perintah tidak valid.")
	// 		}

	// 	case command == "seats_status":
	// 		studioBioskop.TampilkanStatusDenah()

	// 	case command == "transaction_status":
	// 		studioBioskop.TampilkanLaporanPenjualan()

	// 	default:
	// 		fmt.Println("Perintah tidak valid. Coba lagi.")
	// 	}

	// Tampilan status denah awal
	fmt.Println("\nDenah Awal:")
	studioBioskop.TampilkanStatusDenah()

	// Order tiket
	studioBioskop.BeliTiket("A01")
	studioBioskop.BeliTiket("A03")
	studioBioskop.BeliTiket("A05")

	// Tampilan status denah setelah pembelian
	fmt.Println("\nDenah Setelah Pembelian:")
	studioBioskop.TampilkanStatusDenah()

	// Cancel pembelian
	studioBioskop.BatalkanPembelian("A03")

	// Tampilan status denah setelah pembatalan
	fmt.Println("\nDenah Setelah Pembatalan:")
	studioBioskop.TampilkanStatusDenah()

	// Tampilan laporan penjualan
	studioBioskop.TampilkanLaporanPenjualan()
}
