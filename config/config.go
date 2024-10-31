package config

import "go.uber.org/fx"
var Module = fx.Options(
	fx.Provide(NewConfigHolder),
	fx.Provide(NewDatabase),
)