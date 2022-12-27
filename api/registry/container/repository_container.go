package container

import "go_base_project/domain/repository"

// GetHogeRepository .
func (c Container) GetHogeRepository() repository.HogeRepository {
	return repository.NewHogeRepository()
}
