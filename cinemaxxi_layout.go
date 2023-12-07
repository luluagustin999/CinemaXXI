package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type SeatStatus int

const (
	Free SeatStatus = iota
	Sold
)

type StudioBioskop struct {
	JumlahKursi      int
	LableKursi       string
	KursiStatus      map[string]SeatStatus
	LaporanPenjualan []Penjualan
}

type Penjualan struct {
	NomorKursi     string
	WaktuPenjualan time.Time
}

func NewStudioBioskop() *StudioBioskop {
	return &StudioBioskop{
		KursiStatus: make(map[string]SeatStatus),
	}
}

func (s *StudioBioskop) KonfigurasiDenah() {
	fmt.Println("=============== Selamat Datang (Cinema XXI), Silahkan masukkan konfigurasi denah studio ===============")
	fmt.Print("$ Masukkan Label Kursi: ")
	fmt.Scan(&s.LableKursi)
	fmt.Print("$ Masukkan Jumlah Kursi: ")
	fmt.Scan(&s.JumlahKursi)
	fmt.Println("")
	fmt.Printf("Denah Studio berhasil dikonfigurasi dengan label kursi %s dan jumlah kursi %d.\n", s.LableKursi, s.JumlahKursi)
}

func (s *StudioBioskop) KonfigurasiMenu() {
	fmt.Println("")
	fmt.Println("=================== Aplikasi Cinema XXI Layout (kursi tersedia A-5) ===================")
	fmt.Println("")
	fmt.Println("=================== Pilih salah satu menu di bawah ini ===================")
	fmt.Println("A) Pemesanan Kursi —> book_seat {seat_code}")
	fmt.Println("B) Batalkan Kursi —> cancel_seat {seat_code}")
	fmt.Println("C) Laporan Denah —> seats_status")
	fmt.Println("D) Laporan Transaksi —> transaction_status")
	fmt.Println("")
	fmt.Println("Masukkan 'Exit' untuk keluar.")
	fmt.Println("")
}

func (s *StudioBioskop) TampilkanDenahStatus() {
	fmt.Println("")
	fmt.Println("=================== Denah Status ===================")
	for i := 1; i <= s.JumlahKursi; i++ {
		labelKursi := fmt.Sprintf("A%d", i)
		status := s.KursiStatus[labelKursi]
		fmt.Printf("%s - %s\n", labelKursi, getStatusString(status))
	}
}

func getStatusString(status SeatStatus) string {
	if status == Free {
		return "Free"
	}
	return "Sold"
}

func (s *StudioBioskop) BeliTiket(nomorKursi string) {
	if status, exists := s.KursiStatus[nomorKursi]; exists {
		if status == Free {
			s.KursiStatus[nomorKursi] = Sold
			waktuPenjualan := time.Now()
			penjualan := Penjualan{NomorKursi: nomorKursi, WaktuPenjualan: waktuPenjualan}
			s.LaporanPenjualan = append(s.LaporanPenjualan, penjualan)
			fmt.Println("")
			fmt.Printf("Tiket untuk kursi %s berhasil terjual pada %s.\n", nomorKursi, waktuPenjualan.Format("2-Jan-2006 15:04:05"))
		} else {
			fmt.Println("")
			fmt.Println("Maaf, kursi sudah terjual.")
		}
	} else {
		fmt.Println("Nomor kursi tidak valid.")
	}
}

func (s *StudioBioskop) BatalkanPembelian(nomorKursi string) {
	if status, exists := s.KursiStatus[nomorKursi]; exists {
		if status == Sold {
			s.KursiStatus[nomorKursi] = Free
			fmt.Println("")
			fmt.Printf("Pembelian tiket untuk kursi %s berhasil dibatalkan.\n", nomorKursi)
		} else {
			fmt.Println("")
			fmt.Println("Maaf, kursi tidak terjual.")
		}
	} else {
		fmt.Println("Nomor kursi tidak valid.")
	}
}

func (s *StudioBioskop) TampilkanLaporanPenjualan() {
	fmt.Println("")
	fmt.Println("=================== Denah Status ===================")
	totalFree := 0
	totalSold := 0

	for i := 1; i <= s.JumlahKursi; i++ {
		labelKursi := fmt.Sprintf("A%d", i)
		status := s.KursiStatus[labelKursi]
		fmt.Printf("%s - %s\n", labelKursi, getStatusString(status))

		if status == Free {
			totalFree++
		} else if status == Sold {
			totalSold++
		}
	}

	fmt.Println("")
	fmt.Printf("=================== Total %d Free, %d Sold, format (seat_code, datetime) ===================\n", totalFree, totalSold)

	for _, penjualan := range s.LaporanPenjualan {
		fmt.Printf("%s, %s\n", penjualan.NomorKursi, penjualan.WaktuPenjualan.Format("2-Jan-2006 15:04:05"))
	}
}

func main() {
	studioBioskop := NewStudioBioskop()
	studioBioskop.KonfigurasiDenah()
	studioBioskop.KonfigurasiMenu()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("$ Masukkan: ")

	for {

		scanner.Scan()
		command := scanner.Text()

		switch {
		case strings.HasPrefix(command, "book_seat"):
			parts := strings.Fields(command)
			if len(parts) == 2 {
				studioBioskop.BeliTiket(parts[1])
			} else {
				fmt.Println("Format perintah tidak valid.")
			}

		case strings.HasPrefix(command, "cancel_seat"):
			parts := strings.Fields(command)
			if len(parts) == 2 {
				studioBioskop.BatalkanPembelian(parts[1])
			} else {
				fmt.Println("Format perintah tidak valid.")
			}

		case command == "seats_status":
			studioBioskop.TampilkanDenahStatus()

		case command == "transaction_status":
			studioBioskop.TampilkanLaporanPenjualan()

		case command == "Exit" || command == "exit":
			os.Exit(0)

		}
	}
}
