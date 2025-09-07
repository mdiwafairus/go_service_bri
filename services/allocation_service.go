package services

import (
	"go-fiber-api/dto"
	"go-fiber-api/helpers"
	"go-fiber-api/repositories"
	"strconv"
	"strings"
)

type AllocationService struct {
	repo *repositories.AllocationRepository
}

func NewAllocationService(repo *repositories.AllocationRepository) *AllocationService {
	return &AllocationService{repo: repo}
}

func (s *AllocationService) QuotaServiceResponse(nik, mid string) (dto.KuotaResponse, error) {
	wallets, err := s.repo.CheckNikExists(nik)
	if err != nil {
		return dto.KuotaResponse{}, err
	}

	if len(wallets) == 0 {
		return dto.KuotaResponse{}, &NikNotFoundError{}
	}

	retailer, err := s.repo.GetRetailerByMid(mid)
	if err != nil || retailer.RetailerMid == "" {
		retailers, _ := s.repo.GetRetailersByNik(nik)

		var suggestKios []dto.SuggestKios
		for _, r := range retailers {
			suggestKios = append(suggestKios, dto.SuggestKios{
				Mid:  r.RetailerMid,
				Name: r.Name,
			})
		}

		return dto.KuotaResponse{}, &KiosNotMatchError{Suggest: suggestKios}
	}

	var totalUrea, totalZA, totalSP36, totalNPK, totalOrganic float64
	for _, w := range wallets {
		totalUrea += w.Urea
		totalZA += w.ZA
		totalSP36 += w.SP36
		totalNPK += w.NPK
		totalOrganic += w.Organic
	}

	kuotaPupukList := []dto.KuotaPupuk{
		{Pupuk: "urea", Kuota: totalUrea},
		{Pupuk: "za", Kuota: totalZA},
		{Pupuk: "sp36", Kuota: totalSP36},
		{Pupuk: "npk", Kuota: totalNPK},
		{Pupuk: "organic", Kuota: totalOrganic},
	}

	var kelompokTani []string
	for _, w := range wallets {
		kelompokTani = append(kelompokTani, w.FarmerGroupName)
	}

	response := dto.KuotaResponse{
		Mid:          wallets[0].RetailerMid,
		FarmerName:   wallets[0].FarmerName,
		Namakios:     wallets[0].RetailerName,
		KelompokTani: kelompokTani,
		KuotaPupuk:   kuotaPupukList,
	}

	return response, nil
}

func (s *AllocationService) InquiryServiceResponse(nik, komoditas, mid, NamaPupuk string, KgBeli int) (dto.InquiryResponse, error) {

	checkNik, err := s.repo.CheckNikExistsWallet(nik, komoditas)
	if err != nil || checkNik.FarmerNIK == "" {
		return dto.InquiryResponse{}, &NikNotFoundError{}
	}

	retailers, err := s.repo.GetRetailerByMidInquiry(mid)
	if err != nil || len(retailers) == 0 {
		return dto.InquiryResponse{}, &KiosNotMatchError{}
	}

	// retailer id
	var retailerIDs []int
	for _, r := range retailers {
		retailerIDs = append(retailerIDs, r.ID)
	}

	alokasiPetani, err := s.repo.CheckAlokasiPetani(nik, komoditas, retailerIDs)

	if err != nil {
		return dto.InquiryResponse{}, err
	}

	if alokasiPetani == nil {
		return dto.InquiryResponse{}, &AllocationNotFound{}
	}

	totalKuota := helpers.GetKuotaByPupuk(alokasiPetani, NamaPupuk)

	if totalKuota == 0 || totalKuota < KgBeli {
		return dto.InquiryResponse{}, &TidakMemilikiKuota{}
	}
	hargaPupuk, err := s.repo.GetHargaByNama(strings.ToUpper(NamaPupuk))
	if err != nil || hargaPupuk.Harga <= 0 {
		return dto.InquiryResponse{}, &PupukTidakValid{}
	}
	harga := hargaPupuk.Harga

	kuotaSisa := totalKuota - KgBeli
	totalBeli := harga * KgBeli

	resp := dto.InquiryResponse{
		NamaPetani:    alokasiPetani.FarmerName,
		NamaKios:      alokasiPetani.RetailerName,
		KelompokTani:  alokasiPetani.FarmerGroupName,
		NamaPupuk:     NamaPupuk,
		NamaKomoditas: komoditas,
		KgBeli:        strconv.Itoa(KgBeli),
		Harga:         strconv.Itoa(harga),
		KuotaSisa:     strconv.Itoa(kuotaSisa),
		KodeDesa:      alokasiPetani.SubDistrictCode,
		TotalBeli:     strconv.Itoa(totalBeli),
	}

	return resp, nil
}
