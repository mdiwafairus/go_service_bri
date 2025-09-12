package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GenerateRefCode(nik, mid, namaPupuk, namaKomoditas string, kgBeli, refNum, tanggalTransaksi int) string {
	data := fmt.Sprintf("%s|%s|%s|%s|%d|%d|%d",
		nik,
		mid,
		namaPupuk,
		namaKomoditas,
		kgBeli,
		refNum,
		tanggalTransaksi,
	)

	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
