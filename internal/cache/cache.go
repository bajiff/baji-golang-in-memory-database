package cache

import (

	// Package bawaan Golang untuk urusan Concurrency
	"sync"
	"time"
)

type Cache[V any] interface {
	Set(key string, value V, ttl time.Duration)
	Get(key string) (V, bool)
	Delete(key string)
}

type Item[V any] struct {
	Value      V
	Expiration int64
}

// 1. Modifikasi Struct
type MemoryCache[V any] struct {
	// Menambahkan RWMutex ke dalam cetak biru laci kita[cite: 1]
	mu    sync.RWMutex
	items map[string]Item[V]
}

func (c *MemoryCache[V]) DeleteExpired() {
	// TUGAS ANDA 1: Panggil fungsi Lock() dan defer Unlock() pada 'c.mu'.
	// Ingat, membuang data dari map adalah proses 'Write', jadi wajib pakai Lock biasa, BUKAN RLock!

	c.mu.Lock()
	defer c.mu.Unlock()
	sekarang := time.Now().UnixNano()

	// TUGAS ANDA 2: Lakukan perulangan (for range) ke dalam properti 'c.items'.
	// Dalam Go, formatnya: for key, item := range c.items { ... }
	// Di dalam perulangan tersebut, cek apakah 'item.Expiration' lebih besar dari 0
	// DAN apakah 'sekarang' sudah melewati (>) 'item.Expiration'.
	// Jika YA, gunakan fungsi delete(c.items, key) untuk menghapusnya.
	for key, item := range c.items {
		if item.Expiration > 0 && sekarang > item.Expiration {
			delete(c.items, key)
		}
	}
}

// 5. Method startJanitor (Alarm Petugas Kebersihan)
// Method ini akan berjalan selamanya di latar belakang.
func (c *MemoryCache[V]) startJanitor(interval time.Duration) {
	// Membuat jam alarm yang akan berdetak setiap 'interval' waktu.
	ticker := time.NewTicker(interval)

	// Infinite loop (bekerja selamanya)
	for {
		// Menunggu sampai alarm berdetak (Channel / Concurrency)
		<-ticker.C

		// Jika berdetak, bangun dan bersihkan laci!
		c.DeleteExpired()
	}
}

// 6. Modifikasi Pabrik (Fungsi New)
// Kita harus mempekerjakan petugas kebersihan ini tepat saat laci 'MemoryCache' pertama kali dibuat.
func New[V any](cleanupInterval time.Duration) *MemoryCache[V] { // <-- Tambahkan parameter interval
	cacheBaru := &MemoryCache[V]{
		items: make(map[string]Item[V]),
	}

	// TUGAS ANDA 3: Panggil method 'startJanitor' milik 'cacheBaru' sebagai Goroutine.
	// Petunjuk: Gunakan kata kunci 'go' diikuti pemanggilan fungsinya,
	// dan masukkan 'cleanupInterval' ke dalam fungsinya.
	go cacheBaru.startJanitor(cleanupInterval)

	return cacheBaru
}

func (c *MemoryCache[V]) Set(key string, value V, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// TUGAS ANDA: Hitung kapan data ini akan basi.
	// Caranya: Ambil waktu saat ini (time.Now()), lalu tambah dengan durasi (Add(ttl)),
	waktuBasi := time.Now().Add(ttl).UnixNano()
	// lalu ubah menjadi angka integer (UnixNano()).
	// Simpan hasilnya ke dalam variabel bernama 'waktuBasi'.
	// *Petunjuk sintaks: time.Now().Add(...).UnixNano()

	c.items[key] = Item[V]{
		Value:      value,
		Expiration: waktuBasi,
	}

}

// 3. Modifikasi Method GET (Read - Melihat Data)
func (c *MemoryCache[V]) Get(key string) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ditemukan := c.items[key]
	if !ditemukan {
		var zero V
		return zero, false
	}
	// TUGAS ANDA: Lakukan Pengecekan Edge Case (Kasus Ekstrem)
	// Jika item.Expiration lebih besar dari 0 (artinya data ini punya batas waktu)
	// DAN (&&) waktu saat ini dalam format UnixNano() sudah melewati (>) item.Expiration,
	// maka data tersebut sudah BASI!
	// Jika basi, kembalikan 'zero' dan 'false' seolah-olah datanya tidak ditemukan.
	// *Petunjuk sintaks: time.Now().UnixNano()
	if item.Expiration > 0 && time.Now().UnixNano() > item.Expiration {
		var zero V
		return zero, false
	}
	return item.Value, true

}
