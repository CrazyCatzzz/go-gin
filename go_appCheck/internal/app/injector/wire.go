//+build wireinject

package injector

import (
	"appcheck/internal/app/router"

	"github.com/google/wire"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InjectorSet,
		InitWebEngine,
		router.RouterSet,
	)

	return new(Injector), nil, nil
}
