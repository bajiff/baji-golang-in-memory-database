package cache

import (
	"testing"
	"time"
)

// Nama fungsi testing WAJIB diawali dengan kata 'Test' dan menerima parameter t *testing.T
func TestSetAndGet(t *testing.T) {
	// 1. Kita buat laci baru khusus untuk pengetesan.
	// Karena ini pengetesan yang cepat, petugas kebersihan tidak perlu dinyalakan dulu.
	testCache := New[string](0)

	// 2. Kita tentukan kunci dan nilai yang mau diuji.
	kunciTest := "menu_rahasia"
	nilaiTest := "Nasi Goreng Spesial"

	// TUGAS ANDA 1: Panggil fungsi Set milik 'testCache'
	// Masukkan 'kunciTest' dan 'nilaiTest' ke dalamnya dengan batas waktu (TTL) 1 Menit.
	ttl := 1 * time.Minute
	testCache.Set(kunciTest, nilaiTest, ttl)

	// TUGAS ANDA 2: Panggil fungsi Get milik 'testCache' untuk mengambil 'kunciTest'.
	// Simpan kembaliannya di variabel 'hasil' dan 'ditemukan'.
	hasil, ditemukan := testCache.Get(kunciTest); 

	// 3. Inspektur Mutu Melakukan Evaluasi (Verifikasi)
	if !ditemukan {
		// t.Errorf akan menggagalkan pengetesan jika kondisi ini terjadi
		t.Errorf("GAGAL: Data seharusnya ada, tapi dilaporkan tidak ditemukan.")
	} else if hasil != nilaiTest {
		t.Errorf("GAGAL: Data yang diambil tidak cocok. Harapan: %s, Realita: %s", nilaiTest, hasil)
	} 
}
