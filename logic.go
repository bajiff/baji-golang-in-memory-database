// 2. INTERFACE (KONTRAK): Manajer restoran mensyaratkan setiap penyimpanan harus punya fitur 'Set' dan 'Get'.
type StorageSOP[V any] interface {
	Set(key string, value V)
	Get(key string) (V, bool)
}

// 3. STRUCT + 4. GENERIC: Cetak biru wadah penyimpanan fisik yang mematuhi SOP di atas.
// [V any] berarti "Value" bisa berupa tipe data apa saja (string, int, struct lain).
type LaciAjaib[V any] struct {
	Data map[string]V // Menggunakan Map sebagai penyimpanan In-Memory In-Built Go
}