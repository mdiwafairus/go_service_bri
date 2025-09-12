package services

import "go-fiber-api/dto"

type NikNotFoundError struct{}

func (e *NikNotFoundError) Error() string {
	return "nik_not_found"
}

type KiosNotMatchError struct {
	Suggest []dto.SuggestKios
}

func (e *KiosNotMatchError) Error() string {
	return "kios_not_match"
}

type AllocationNotFound struct{}

func (e *AllocationNotFound) Error() string {
	return "allocation_not_found"
}

type TidakMemilikiKuota struct{}

func (e *TidakMemilikiKuota) Error() string {
	return "tidak_memiliki_kuota"
}

type NikTidakValid struct{}

func (e *NikTidakValid) Error() string {
	return "nik_tidak_valid"
}

type PupukTidakValid struct{}

func (e *PupukTidakValid) Error() string {
	return "pupuk_tidak_valid"
}

type DuplicateTransactionError struct{}

func (e *DuplicateTransactionError) Error() string {
	return "transaksi_duplikat"
}
