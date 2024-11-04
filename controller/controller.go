package controller

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewLoginController),
	fx.Provide(NewComplaintController),
)