package services

import (
	"errors"
	"go-fiber-api/constants"
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
		return dto.KuotaResponse{}, errors.New(constants.MsgNikNotFound)
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
		Mid:          mid,
		FarmerName:   wallets[0].FarmerName,
		Namakios:     wallets[0].RetailerName,
		KelompokTani: kelompokTani,
		KuotaPupuk:   kuotaPupukList,
	}

	return response, nil
}
