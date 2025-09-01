package services

import (
	"go-fiber-api/dto"
	"go-fiber-api/repositories"
)

type AllocationService struct {
	repo *repositories.AllocationRepository
}

func NewAllocationService(repo *repositories.AllocationRepository) *AllocationService {
	return &AllocationService{repo: repo}
}

func (s *AllocationService) GetNikExistsResponse(nik, mid string) (dto.KuotaResponse, error) {
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

func (s *AllocationService) InquiryServiceResponse(nik string, mid int) (dto.InquiryResponse, error) {
	return dto.InquiryResponse{}, nil
}
