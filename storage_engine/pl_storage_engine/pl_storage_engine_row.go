package pl_storage_engine

//Index stored at 0, 7, 15, 23, 31...
//Pointer stored at 8, 16, 24, 32...
type PlStorageEngineRow struct {
	Index   uint64
	Pointer uint64
}

func IsValidRowIndexAddr(i uint64) bool {
	i = i + 1
	if i == 1 || i%8 == 0 {
		return true
	}
	return false
}

func IsValidRowPointerAddr(p uint64) bool {
	if p%8 == 0 {
		return true
	}
	return false
}
