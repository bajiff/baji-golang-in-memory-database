package cache

import "sync" // Package bawaan Golang untuk urusan Concurrency

type Cache[V any] interface {
	Set(key string, value V)
	Get(key string) (V, bool)
	Delete(key string)
}

type Item[V any] struct {
	Value V
	Expiration int64
}

// 1. Modifikasi Struct
type MemoryCache[V any] struct {
    // Menambahkan RWMutex ke dalam cetak biru laci kita[cite: 1]
    mu    sync.RWMutex 
    items map[string]Item[V]
}

func New[V any]() *MemoryCache[V] {
    return &MemoryCache[V]{
        items: make(map[string]Item[V]),
        // mu tidak perlu di-make, otomatis siap digunakan karena berupa struct bawaan.
    }
}

// 2. Modifikasi Method SET (Write - Mengubah Data)
func (c *MemoryCache[V]) Set(key string, value V) {
    // TUGAS ANDA: Panggil fungsi Lock() pada properti 'mu' milik 'c' 
    // untuk mengunci akses tulis secara eksklusif[cite: 1].
		c.mu.Lock()
		defer c.mu.Unlock()
    // TUGAS ANDA: Gunakan kata kunci 'defer' diikuti fungsi Unlock() pada properti 'mu' 
    // agar kunci OTOMATIS dikembalikan saat fungsi selesai[cite: 1].
    c.items[key] = Item[V]{
			Value: value,
			Expiration: 0,
		}
    
    // Masukkan data (Kode Anda sebelumnya sudah benar)
}

// 3. Modifikasi Method GET (Read - Melihat Data)
func (c *MemoryCache[V]) Get(key string) (V, bool) {
    // TUGAS ANDA: Karena ini hanya membaca, panggil RLock() (Read Lock) pada 'mu'.
    // Ini mengizinkan ribuan pembacaan secara bersamaan[cite: 1].
		c.mu.RLock()
		defer c.mu.RUnlock()
    
    // TUGAS ANDA: Gunakan 'defer' diikuti fungsi RUnlock() (Read Unlock) pada 'mu'.
    
    item, ditemukan := c.items[key]
		if !ditemukan {
			var zero V
			return  zero, false
		}
    return item.Value, true
}