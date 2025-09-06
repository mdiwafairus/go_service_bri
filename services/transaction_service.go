package services

import (
	"go-fiber-api/dto"
	"go-fiber-api/helpers"
	"go-fiber-api/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) TransactionServiceResponse(nik, mid, NamaPupuk, NamaKomoditas string, KgBeli, TotalRupiah, RefNum, TanggalTransaksi int) (dto.TransactionResponse, error) {

	if err := helpers.ValidateNIK(nik); err != nil {
		return dto.TransactionResponse{}, &NikTidakValid{}
	}

	response := dto.TransactionResponse{
		Mid:              mid,
		Nik:              nik,
		NamaPupuk:        NamaPupuk,
		NamaKomoditas:    NamaKomoditas,
		KgBeli:           KgBeli,
		TotalRupiah:      TotalRupiah,
		RefNum:           RefNum,
		TrxId:            1,
		ClientId:         1,
		TanggalTransaksi: 20230101,
	}
	return response, nil
}
