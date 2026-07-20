package cache

import "sync"

// 1. Buat Interface bernama 'Cache' dengan tipe Generic [V any].
// Interface ini wajib memiliki 3 metode dasar:
// - Set(key string, value V)
// - Get(key string) (V, bool)
// - Delete(key string)
type Cache[V any] interface {
	Set(key string, value V)
	Get(key string) (V, bool)
	Delete(key string)
}

// 2. Buat Struct bernama 'MemoryCache' dengan tipe Generic [V any].
// Di dalam struct ini, buat satu properti bernama 'items' yang berupa map.
// Map tersebut harus memiliki Key bertipe string, dan Value bertipe V.
type MemoryCache[V any] struct {
	mu sync.RWMutex
		items map[string]V
}

// 3. Buat fungsi 'New' sebagai Constructor (Pabrik).
// Fungsi ini bertugas mencetak objek MemoryCache baru yang siap pakai.
// Ingat, map di Golang harus diinisialisasi dengan 'make' sebelum digunakan, 
// jika tidak akan menyebabkan Panic (Fatal Error).
func New[V any]() *MemoryCache[V] {
    return &MemoryCache[V]{
				items: make(map[string]V),
    }
}
// 1. Method SET (Memasukkan barang ke laci)
// (c *MemoryCache[V]) adalah Receiver. 'c' adalah inisial dari Cache.
func (c *MemoryCache[V]) Set(key string, value V) {
	// TUGAS ANDA: Masukkan parameter 'value' ke dalam properti map 'items' milik 'c' 
	// menggunakan 'key' sebagai indeksnya.
	c.items[key] = value
}

// 2. Method GET (Mengambil barang dari laci)
func (c *MemoryCache[V]) Get(key string) (V, bool) {
	// TUGAS ANDA: Ambil nilai dari map berdasarkan 'key'.
	// Ingat, map di Go mengembalikan 2 nilai secara berurutan: data dan status boolean.
	// Kembalikan (return) kedua nilai tersebut agar aman dari Edge Cases.
	nilai, ditemukan := c.items[key]
	return nilai, ditemukan
	
}

// 3. Method DELETE (Membuang barang dari laci)
func (c *MemoryCache[V]) Delete(key string) {
	// TUGAS ANDA: Gunakan fungsi bawaan Go yaitu delete(map, key) 
	// untuk menghapus data dari memori.
	delete(c.items, key)
}