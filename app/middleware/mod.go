package middleware

import "app_land_x/app/controller"

type Middleware struct {
	ctrl *controller.Controller
}

func New(ctrl *controller.Controller) *Middleware {
	return &Middleware{
		ctrl,
	}
}
