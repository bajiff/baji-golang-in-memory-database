package cache

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