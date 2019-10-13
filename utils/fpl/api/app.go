package api

import (
	"github.com/jonnoking/vidukavindaloo/utils/fpl/config"
)

//var Config *config.FPLConfig

type API struct {
	Config *config.FPLConfig
}

func New(config *config.FPLConfig) *API {
	return &API{
		Config: config,
	}
}
