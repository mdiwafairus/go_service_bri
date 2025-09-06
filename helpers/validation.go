package helpers

import (
	"errors"
	"go-fiber-api/dto"
	"regexp"
	"strings"
	"unicode"

	"github.com/rs/zerolog/log"
)

var JenisPupukMap = map[string]bool{
	"UREA":        true,
	"NPK":         true,
	"NPK_FORMULA": true,
	"SP36":        true,
	"ZA":          true,
	"ORGANIC":     true,
	"POC":         true,
}

var KomoditasMap = map[string]bool{
	"PADI":         true,
	"CABAI":        true,
	"BAWANG_MERAH": true,
	"BAWANG_PUTIH": true,
	"KAKAO":        true,
	"TEBU_RAKYAT":  true,
	"JAGUNG":       true,
	"KOPI":         true,
	"KEDELAI":      true,
}

func NormalizeKey(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)

	re := regexp.MustCompile(`[^A-Z0-9]+`)
	s = re.ReplaceAllString(s, "_")

	s = strings.Trim(s, "_")

	return s
}

// Validator pupuk
func IsValidPupuk(pupuk string) bool {
	key := NormalizeKey(pupuk)
	if !JenisPupukMap[key] {
		log.Warn().Str("raw", pupuk).Str("normalized", key).Msg("Invalid Jenis Pupuk")
		return false
	}
	return true
}

// Validator komoditas
func IsValidKomoditas(komoditas string) bool {
	key := NormalizeKey(komoditas)
	if KomoditasMap[key] {
		log.Warn().Str("raw", komoditas).Str("normalized", key).Msg("Invalid Jenis Komoditas")
		return true
	}

	return false
}

func GetKuotaByPupuk(wallet *dto.NikExistsResponse, pupuk string) int {
	total := 0

	switch pupuk {
	case "UREA":
		total += int(wallet.Urea)
	case "ZA":
		total += int(wallet.ZA)
	case "SP36":
		total += int(wallet.SP36)
	case "NPK":
		total += int(wallet.NPK)
	case "ORGANIC":
		total += int(wallet.Organic)
	case "NPK_FORMULA":
		total += int(wallet.NpkFormula)
	case "POC":
		total += int(wallet.Poc)
	}

	return total
}

func ValidateNIK(nik string) error {
	if len(nik) != 16 {
		return errors.New("NIK harus 16 digit")
	}
	for _, r := range nik {
		if !unicode.IsDigit(r) {
			return errors.New("NIK hanya boleh angka")
		}
	}
	return nil
}
