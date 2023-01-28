package impl

import (
	"go.uber.org/fx"
)

var Module = fx.Module("app",
	fx.Provide(NewUserAppService),
)
