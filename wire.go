//+build wireinject

package main

import (
	"di-in-golang/pkg"
	"github.com/google/wire"
)

func SetupApplication() (*Models.Hochbeet, error) {
	wire.Build(Models.GeneriereHochbeet, Models.GeneriereVentil)
	return nil, nil
}
