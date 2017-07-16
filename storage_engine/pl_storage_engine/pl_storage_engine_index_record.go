package pl_storage_engine

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

//Index stored at 0, 16, 32, 48, 64...
//Pointer stored at 8, 24, 40, 56, 72...
type PlStorageEngineRow struct {
	Index   uint64
	Pointer uint64
}

func IsValidRowIndexAddr(i uint64) bool {
	if i == 0 || i%16 == 0 {
		return true
	}
	return false
}

func IsValidRowPointerAddr(p uint64) bool {
	p = p - 8
	if IsValidRowIndexAddr(p) {
		return true
	}
	return false
}

func writeIndex(f *os.File, index uint64, row PlStorageEngineRow) (n int, err error) {
	if !IsValidRowIndexAddr(index) {
		return 0, fmt.Errorf("Invalid index address: %v", index)
	}
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.LittleEndian, &row)
	if err != nil {
		return 0, err
	}
	n, err = f.WriteAt(buf.Bytes(), int64(index))
	return n, err
}
