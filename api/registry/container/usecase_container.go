package container

import (
	"go_base_project/application/usecase"
	"go_base_project/domain/repository"
)

func (c Container) GetHogeUsecase(hr repository.HogeRepository) usecase.HogeUsecase {
	return usecase.NewHogeUsecase(hr)
}
