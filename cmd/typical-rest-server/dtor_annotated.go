package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-rest-server/internal/infra"
)

func init() {
	typapp.AppendDtor(
		&typapp.Destructor{Fn: infra.Disconnect},
	)
}
