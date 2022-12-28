package container

import (
	"go_base_project/application/usecase"
	v1 "go_base_project/interfaces/handler/v1"
)

func (c Container) GetAppHandler() v1.AppHandler {
	return &struct {
		v1.HogeHandler
	}{
		// v1/hoge
		HogeHandler: c.GetHogeHandler(
			c.GetHogeUsecase(
				c.GetHogeRepository(),
			),
		),
	}
}

func (c Container) GetHogeHandler(hu usecase.HogeUsecase) v1.HogeHandler {
	return v1.NewHogeHandler(hu)
}
