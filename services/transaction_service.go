package services

import (
	"fmt"
	"go-fiber-api/dto"
	"go-fiber-api/helpers"
	"go-fiber-api/models"
	"go-fiber-api/repositories"
	"go-fiber-api/utils"
	"time"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func buildTransactionData(wallet models.PspWallet, request dto.TransactionResponse, refCode string, usedQuota int) map[string]interface{} {
	return map[string]interface{}{
		"wallet_id":  wallet.IdWallet,
		"nik":        request.Nik,
		"komoditas":  request.NamaKomoditas,
		"ref_code":   refCode,
		"used_quota": usedQuota,
	}
}

func (s *TransactionService) TransactionServiceResponse(nik, mid, NamaPupuk, NamaKomoditas string, KgBeli, TotalRupiah, RefNum, TanggalTransaksi int) (dto.TransactionResponse, error) {

	if err := helpers.ValidateNIK(nik); err != nil {
		return dto.TransactionResponse{}, &NikTidakValid{}
	}

	retailers, err := s.repo.GetRetailerByMidInquiry(mid)
	if err != nil || len(retailers) == 0 {
		return dto.TransactionResponse{}, &KiosNotMatchError{}
	}

	tx := s.repo.GetDB().Begin()
	if tx.Error != nil {
		return dto.TransactionResponse{}, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	checkWallet, err := s.repo.CheckAlokasiPetaniTransaction(nik, NamaKomoditas, []int{retailers[0].ID})

	if err != nil {
		return dto.TransactionResponse{}, err
	}

	if checkWallet == nil {
		return dto.TransactionResponse{}, &AllocationNotFound{}
	}

	checkTransaksi, err := s.repo.CheckDuplicateTransaction(RefNum)

	if err != nil {
		return dto.TransactionResponse{}, err
	}

	if checkTransaksi != nil {
		return dto.TransactionResponse{}, &DuplicateTransactionError{}
	}

	remainingKg := KgBeli

	var totalAvailableQuota int
	for _, wallet := range checkWallet {
		quotaMap := map[string]int{
			"UREA":        int(wallet.Urea),
			"NPK":         int(wallet.Npk),
			"NPK_FORMULA": int(wallet.NpkFormula),
			"SP36":        int(wallet.Sp36),
			"ZA":          int(wallet.Za),
			"ORGANIC":     int(wallet.Organic),
			"POC":         int(wallet.Poc),
		}
		totalAvailableQuota += quotaMap[NamaPupuk]
	}

	if totalAvailableQuota < remainingKg {
		return dto.TransactionResponse{}, &TidakMemilikiKuota{}
	}
	refCode := utils.GenerateRefCode(
		nik,
		mid,
		NamaPupuk,
		NamaKomoditas,
		KgBeli,
		RefNum,
		TanggalTransaksi,
	)

	fmt.Println("Reference Code:", refCode)

	listAllocationUsed := []map[string]interface{}{}
	listTransaction := []map[string]interface{}{}

	for _, wallet := range checkWallet {

		currentQuota := 30
		// currentQuota := helpers.GetKuotaByPupuk(wallet, NamaPupuk)

		usedQuota := min(remainingKg, currentQuota)

		listAllocationUsed = append(listAllocationUsed, map[string]interface{}{
			"update": map[string]interface{}{"id": wallet.IdWallet},
			"data":   map[string]interface{}{NamaPupuk: currentQuota - usedQuota},
		})

		listTransaction = append(listTransaction, buildTransactionData(wallet, dto.TransactionResponse{}, NamaKomoditas, usedQuota))

		remainingKg -= usedQuota

		if remainingKg <= 0 {
			break
		}
	}

	newTrx := models.KartanFarmerTransaction{
		SubdistrictCode:   "1234567890",
		DistrictCode:      "123456",
		CityCode:          "1234",
		ProvinceCode:      "12",
		NationalID:        nik,
		FarmerName:        "John Doe",
		UreaKg:            0.0,
		NpkKg:             0.0,
		Sp36Kg:            0.0,
		ZaKg:              0.0,
		NpkFormulaKg:      0.0,
		OrganicKg:         0.0,
		PocKg:             0.0,
		ApprovedBy:        nil,
		Notes:             nil,
		IsSent:            3,
		KioskName:         nil,
		RedeemDay:         0,
		RedeemMonth:       0,
		RedeemYear:        0,
		BankTimestamp:     time.Now(),
		UpdatedAt:         time.Now(),
		ApprovalTimestamp: nil,
		SourceType:        1,
		TransactionCode:   "TRX-001",
		ReferenceCode:     "REF-001",
		ProvinceName:      nil,
		DistrictName:      nil,
		SubdistrictName:   nil,
		VillageName:       nil,
		CreatedAt:         time.Now(),
		Commodity:         nil,
		FarmerGroupID:     nil,
		FarmerGroupName:   nil,
		ReferenceNumber:   nil,
		WalletID:          nil,
	}

	if err := s.repo.InsertTransaction(tx, &newTrx); err != nil {
		tx.Rollback()
		return dto.TransactionResponse{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return dto.TransactionResponse{}, err
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

func (s *TransactionService) ReversalTransactionService(nik, mid string, refnum int) (dto.ReversalTransactionResponse, error) {

	response := dto.ReversalTransactionResponse{
		Nik:    "1233",
		Mid:    "98777",
		RefNum: 123,
	}

	return response, nil
}
