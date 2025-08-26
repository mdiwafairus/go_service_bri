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
