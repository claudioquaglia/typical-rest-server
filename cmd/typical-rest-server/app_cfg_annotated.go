package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-rest-server/internal/infra"
)

func init() {
	typapp.AppendCtor(
		&typapp.Constructor{
			Name: "",
			Fn: func() (*infra.AppCfg, error) {
				var cfg infra.AppCfg
				if err := envconfig.Process("APP", &cfg); err != nil {
					return nil, err
				}
				return &cfg, nil
			},
		},
		&typapp.Constructor{
			Name: "",
			Fn: func() (*infra.RedisCfg, error) {
				var cfg infra.RedisCfg
				if err := envconfig.Process("REDIS", &cfg); err != nil {
					return nil, err
				}
				return &cfg, nil
			},
		},
		&typapp.Constructor{
			Name: "",
			Fn: func() (*infra.PostgresCfg, error) {
				var cfg infra.PostgresCfg
				if err := envconfig.Process("PG", &cfg); err != nil {
					return nil, err
				}
				return &cfg, nil
			},
		},
	)
}
