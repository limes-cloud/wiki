package logic

import "h1/config"

type Logic struct {
	conf *config.Config
}

func NewLogic(conf *config.Config) *Logic {
	return &Logic{
		conf: conf,
	}
}
